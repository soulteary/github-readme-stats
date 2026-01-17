package fetchers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/soulteary/github-readme-stats/internal/common"
)

const repoQuery = `
      fragment RepoInfo on Repository {
        name
        nameWithOwner
        isPrivate
        isArchived
        isTemplate
        stargazers {
          totalCount
        }
        description
        primaryLanguage {
          color
          id
          name
        }
        forkCount
      }
      query getRepo($login: String!, $repo: String!) {
        user(login: $login) {
          repository(name: $repo) {
            ...RepoInfo
          }
        }
        organization(login: $login) {
          repository(name: $repo) {
            ...RepoInfo
          }
        }
      }
    `

// RepoGraphQLResponse represents the GraphQL response for repository
type RepoGraphQLResponse struct {
	Data struct {
		User struct {
			Repository *RepoNode `json:"repository"`
		} `json:"user"`
		Organization struct {
			Repository *RepoNode `json:"repository"`
		} `json:"organization"`
	} `json:"data"`
	Errors []struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"errors"`
}

// RepoNode represents a repository node
type RepoNode struct {
	Name          string `json:"name"`
	NameWithOwner string `json:"nameWithOwner"`
	IsPrivate     bool   `json:"isPrivate"`
	IsArchived    bool   `json:"isArchived"`
	IsTemplate    bool   `json:"isTemplate"`
	Stargazers    struct {
		TotalCount int `json:"totalCount"`
	} `json:"stargazers"`
	Description     string `json:"description"`
	PrimaryLanguage *struct {
		Color string `json:"color"`
		ID    string `json:"id"`
		Name  string `json:"name"`
	} `json:"primaryLanguage"`
	ForkCount int `json:"forkCount"`
}

// FetchRepo fetches repository data
func FetchRepo(username, reponame string) (*RepoData, error) {
	if username == "" && reponame == "" {
		return nil, common.NewMissingParamError([]string{"username", "repo"}, "/api/pin?username=USERNAME&repo=REPO_NAME")
	}
	if username == "" {
		return nil, common.NewMissingParamError([]string{"username"}, "/api/pin?username=USERNAME&repo=REPO_NAME")
	}
	if reponame == "" {
		return nil, common.NewMissingParamError([]string{"repo"}, "/api/pin?username=USERNAME&repo=REPO_NAME")
	}

	variables := map[string]interface{}{
		"login": username,
		"repo":  reponame,
	}

	req := common.GraphQLRequest{
		Query:     repoQuery,
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

	var graphqlResp RepoGraphQLResponse
	if err := json.Unmarshal(body, &graphqlResp); err != nil {
		return nil, err
	}

	// Filter out non-critical GraphQL errors (e.g., "Could not resolve to an Organization")
	// These errors are expected when querying a user (organization query fails) or an org (user query fails)
	// Only check for critical errors like NOT_FOUND, RATE_LIMITED, etc.
	var criticalErrors []struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	}
	for _, err := range graphqlResp.Errors {
		// Ignore errors about resolving to Organization or User, as these are expected
		// when the login is a user (org query fails) or an org (user query fails)
		if !strings.Contains(err.Message, "Could not resolve to an Organization") &&
			!strings.Contains(err.Message, "Could not resolve to a User") {
			criticalErrors = append(criticalErrors, err)
		}
	}

	if len(criticalErrors) > 0 {
		return nil, common.NewCustomError(
			criticalErrors[0].Message,
			common.ErrorTypeGraphQLError,
		)
	}

	data := graphqlResp.Data
	if data.User.Repository == nil && data.Organization.Repository == nil {
		return nil, common.NewCustomError("Not found", common.ErrorTypeUserNotFound)
	}

	var repoNode *RepoNode
	isUser := data.Organization.Repository == nil && data.User.Repository != nil
	isOrg := data.User.Repository == nil && data.Organization.Repository != nil

	if isUser {
		if data.User.Repository == nil || data.User.Repository.IsPrivate {
			return nil, common.NewCustomError("User Repository Not found", common.ErrorTypeUserNotFound)
		}
		repoNode = data.User.Repository
	} else if isOrg {
		if data.Organization.Repository == nil || data.Organization.Repository.IsPrivate {
			return nil, common.NewCustomError("Organization Repository Not found", common.ErrorTypeUserNotFound)
		}
		repoNode = data.Organization.Repository
	} else {
		return nil, common.NewCustomError("Repository Not found", common.ErrorTypeUserNotFound)
	}

	repoData := &RepoData{
		Name:          repoNode.Name,
		NameWithOwner: repoNode.NameWithOwner,
		IsPrivate:     repoNode.IsPrivate,
		IsFork:        false, // Will be set if needed
		IsArchived:    repoNode.IsArchived,
		IsTemplate:    repoNode.IsTemplate,
		Stargazers: StargazersData{
			TotalCount: repoNode.Stargazers.TotalCount,
		},
		Forks: ForksData{
			TotalCount: repoNode.ForkCount,
		},
		Description: repoNode.Description,
	}

	if repoNode.PrimaryLanguage != nil {
		repoData.PrimaryLanguage = &LanguageData{
			Name:  repoNode.PrimaryLanguage.Name,
			Color: repoNode.PrimaryLanguage.Color,
		}
	}

	return repoData, nil
}
