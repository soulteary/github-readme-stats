package fetchers

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
	"sort"
	"strings"

	"github.com/soulteary/github-readme-stats/internal/common"
)

const topLanguagesQuery = `
      query userInfo($login: String!) {
        user(login: $login) {
          repositories(ownerAffiliations: OWNER, isFork: false, first: 100) {
            nodes {
              name
              languages(first: 10, orderBy: {field: SIZE, direction: DESC}) {
                edges {
                  size
                  node {
                    color
                    name
                  }
                }
              }
            }
          }
        }
      }
    `

// TopLanguagesGraphQLResponse represents the GraphQL response for top languages
type TopLanguagesGraphQLResponse struct {
	Data struct {
		User struct {
			Repositories struct {
				Nodes []struct {
					Name      string `json:"name"`
					Languages struct {
						Edges []struct {
							Size int64 `json:"size"`
							Node struct {
								Color string `json:"color"`
								Name  string `json:"name"`
							} `json:"node"`
						} `json:"edges"`
					} `json:"languages"`
				} `json:"nodes"`
			} `json:"repositories"`
		} `json:"user"`
	} `json:"data"`
	Errors []struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"errors"`
}

// LanguageStatsWithCount represents language statistics with count
type LanguageStatsWithCount struct {
	Name      string
	Color     string
	Size      int64
	Count     int
	RankIndex float64
}

// FetchTopLanguages fetches top languages for a given username
func FetchTopLanguages(
	username string,
	excludeRepo []string,
	sizeWeight float64,
	countWeight float64,
) (map[string]*LanguageStats, error) {
	if username == "" {
		return nil, common.NewMissingParamError([]string{"username"}, "")
	}

	if sizeWeight == 0 {
		sizeWeight = 1
	}

	variables := map[string]interface{}{
		"login": username,
	}

	req := common.GraphQLRequest{
		Query:     topLanguagesQuery,
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

	var graphqlResp TopLanguagesGraphQLResponse
	if err := json.Unmarshal(body, &graphqlResp); err != nil {
		return nil, err
	}

	if len(graphqlResp.Errors) > 0 {
		if graphqlResp.Errors[0].Type == "NOT_FOUND" {
			return nil, common.NewCustomError(
				graphqlResp.Errors[0].Message,
				common.ErrorTypeUserNotFound,
			)
		}
		if graphqlResp.Errors[0].Message != "" {
			msg := strings.Split(graphqlResp.Errors[0].Message, "\n")[0]
			if len(msg) > 90 {
				msg = msg[:90]
			}
			return nil, common.NewCustomError(msg, graphqlResp.Errors[0].Type)
		}
		return nil, common.NewCustomError(
			"Something went wrong while trying to retrieve the language data using the GraphQL API.",
			common.ErrorTypeGraphQLError,
		)
	}

	repoNodes := graphqlResp.Data.User.Repositories.Nodes
	allExcludedRepos := append(excludeRepo, common.GetExcludeRepositories()...)
	repoToHide := make(map[string]bool)
	for _, repo := range allExcludedRepos {
		repoToHide[repo] = true
	}

	// Filter out excluded repositories
	var filteredRepos []struct {
		Name      string `json:"name"`
		Languages struct {
			Edges []struct {
				Size int64 `json:"size"`
				Node struct {
					Color string `json:"color"`
					Name  string `json:"name"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"languages"`
	}

	for _, node := range repoNodes {
		if !repoToHide[node.Name] {
			filteredRepos = append(filteredRepos, node)
		}
	}

	// Flatten language edges
	type langEdge struct {
		Size int64
		Node struct {
			Color string
			Name  string
		}
	}

	var allLangEdges []langEdge
	for _, repo := range filteredRepos {
		if len(repo.Languages.Edges) > 0 {
			for _, edge := range repo.Languages.Edges {
				allLangEdges = append(allLangEdges, langEdge{
					Size: edge.Size,
					Node: struct {
						Color string
						Name  string
					}{
						Color: edge.Node.Color,
						Name:  edge.Node.Name,
					},
				})
			}
		}
	}

	// Aggregate languages
	langMap := make(map[string]*LanguageStatsWithCount)
	for _, edge := range allLangEdges {
		langName := edge.Node.Name
		if existing, ok := langMap[langName]; ok {
			existing.Size += edge.Size
			existing.Count++
		} else {
			langMap[langName] = &LanguageStatsWithCount{
				Name:  edge.Node.Name,
				Color: edge.Node.Color,
				Size:  edge.Size,
				Count: 1,
			}
		}
	}

	// Calculate ranking index
	for _, lang := range langMap {
		lang.RankIndex = math.Pow(float64(lang.Size), sizeWeight) * math.Pow(float64(lang.Count), countWeight)
	}

	// Sort by rank index
	type langPair struct {
		Name  string
		Stats *LanguageStatsWithCount
	}
	var sortedLangs []langPair
	for name, stats := range langMap {
		sortedLangs = append(sortedLangs, langPair{Name: name, Stats: stats})
	}

	sort.Slice(sortedLangs, func(i, j int) bool {
		return sortedLangs[i].Stats.RankIndex > sortedLangs[j].Stats.RankIndex
	})

	// Convert to result format
	result := make(map[string]*LanguageStats)
	for _, pair := range sortedLangs {
		result[pair.Name] = &LanguageStats{
			Name:      pair.Stats.Name,
			Color:     pair.Stats.Color,
			Size:      pair.Stats.Size,
			RepoCount: pair.Stats.Count,
		}
	}

	return result, nil
}
