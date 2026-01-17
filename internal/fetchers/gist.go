package fetchers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/soulteary/github-readme-stats/internal/common"
)

const gistQuery = `
query gistInfo($gistName: String!) {
    viewer {
        gist(name: $gistName) {
            description
            owner {
                login
            }
            stargazerCount
            forks {
                totalCount
            }
            files {
                name
                language {
                    name
                }
                size
            }
        }
    }
}
`

// GistFileInfo represents file information from GraphQL response
type GistFileInfo struct {
	Name     string `json:"name"`
	Language *struct {
		Name string `json:"name"`
	} `json:"language"`
	Size int `json:"size"`
}

// GistGraphQLResponse represents the GraphQL response for gist
type GistGraphQLResponse struct {
	Data struct {
		Viewer struct {
			Gist *struct {
				Description string `json:"description"`
				Owner       struct {
					Login string `json:"login"`
				} `json:"owner"`
				StargazerCount int `json:"stargazerCount"`
				Forks          struct {
					TotalCount int `json:"totalCount"`
				} `json:"forks"`
				Files []GistFileInfo `json:"files"`
			} `json:"gist"`
		} `json:"viewer"`
	} `json:"data"`
	Errors []struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"errors"`
}

// CalculatePrimaryLanguage calculates the primary language of a gist by files size
func CalculatePrimaryLanguage(files []GistFileInfo) string {
	languages := make(map[string]int)

	for _, file := range files {
		if file.Language != nil {
			languages[file.Language.Name] += file.Size
		}
	}

	if len(languages) == 0 {
		return "Unspecified"
	}

	var primaryLanguage string
	maxSize := 0
	for lang, size := range languages {
		if size > maxSize {
			maxSize = size
			primaryLanguage = lang
		}
	}

	return primaryLanguage
}

// FetchGist fetches GitHub gist information by given ID
func FetchGist(id string) (*GistData, error) {
	if id == "" {
		return nil, common.NewMissingParamError([]string{"id"}, "/api/gist?id=GIST_ID")
	}

	variables := map[string]interface{}{
		"gistName": id,
	}

	req := common.GraphQLRequest{
		Query:     gistQuery,
		Variables: variables,
	}

	resp, err := common.Retryer(func(v interface{}, token string) (*http.Response, error) {
		return common.Request(req, token)
	}, variables, 0)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	var graphqlResp GistGraphQLResponse
	if err := json.Unmarshal(body, &graphqlResp); err != nil {
		return nil, err
	}

	if len(graphqlResp.Errors) > 0 {
		return nil, common.NewCustomError(
			graphqlResp.Errors[0].Message,
			common.ErrorTypeGraphQLError,
		)
	}

	if graphqlResp.Data.Viewer.Gist == nil {
		return nil, common.NewCustomError("Gist not found", common.ErrorTypeUserNotFound)
	}

	data := graphqlResp.Data.Viewer.Gist

	// Convert files
	var gistFiles []GistFile
	for _, file := range data.Files {
		langName := ""
		if file.Language != nil {
			langName = file.Language.Name
		}
		gistFiles = append(gistFiles, GistFile{
			Name:     file.Name,
			Language: langName,
			Size:     file.Size,
		})
	}

	primaryLang := CalculatePrimaryLanguage(data.Files)

	gistData := &GistData{
		Name:        primaryLang,
		Description: data.Description,
		Files:       gistFiles,
		Owner: GistOwner{
			Login: data.Owner.Login,
		},
		Stargazers: StargazersData{
			TotalCount: data.StargazerCount,
		},
		Forks: ForksData{
			TotalCount: data.Forks.TotalCount,
		},
		IsPublic: true,
	}

	return gistData, nil
}
