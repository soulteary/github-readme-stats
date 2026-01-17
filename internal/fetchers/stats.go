package fetchers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/soulteary/github-readme-stats/internal/common"
)

const (
	graphQLURL = "https://api.github.com/graphql"
	restAPIURL = "https://api.github.com/search/commits"
)

// GraphQLReposField is the GraphQL query field for repositories
const GraphQLReposField = `
  repositories(first: 100, ownerAffiliations: OWNER, orderBy: {direction: DESC, field: STARGAZERS}, after: $after) {
    totalCount
    nodes {
      name
      stargazers {
        totalCount
      }
    }
    pageInfo {
      hasNextPage
      endCursor
    }
  }
`

// GraphQLReposQuery is the GraphQL query for repositories
const GraphQLReposQuery = `
  query userInfo($login: String!, $after: String) {
    user(login: $login) {
      %s
    }
  }
`

// GraphQLStatsQuery is the GraphQL query for stats
const GraphQLStatsQuery = `
  query userInfo($login: String!, $after: String, $includeMergedPullRequests: Boolean!, $includeDiscussions: Boolean!, $includeDiscussionsAnswers: Boolean!, $startTime: DateTime) {
    user(login: $login) {
      name
      login
      commits: contributionsCollection (from: $startTime) {
        totalCommitContributions
      }
      reviews: contributionsCollection {
        totalPullRequestReviewContributions
      }
      repositoriesContributedTo(first: 1, contributionTypes: [COMMIT, ISSUE, PULL_REQUEST, REPOSITORY]) {
        totalCount
      }
      pullRequests(first: 1) {
        totalCount
      }
      mergedPullRequests: pullRequests(states: MERGED) @include(if: $includeMergedPullRequests) {
        totalCount
      }
      openIssues: issues(states: OPEN) {
        totalCount
      }
      closedIssues: issues(states: CLOSED) {
        totalCount
      }
      followers {
        totalCount
      }
      repositoryDiscussions @include(if: $includeDiscussions) {
        totalCount
      }
      repositoryDiscussionComments(onlyAnswers: true) @include(if: $includeDiscussionsAnswers) {
        totalCount
      }
      %s
    }
  }
`

// GraphQLResponse represents the GraphQL response structure
type GraphQLResponse struct {
	Data struct {
		User struct {
			Name    string `json:"name"`
			Login   string `json:"login"`
			Commits struct {
				TotalCommitContributions int `json:"totalCommitContributions"`
			} `json:"commits"`
			Reviews struct {
				TotalPullRequestReviewContributions int `json:"totalPullRequestReviewContributions"`
			} `json:"reviews"`
			RepositoriesContributedTo struct {
				TotalCount int `json:"totalCount"`
			} `json:"repositoriesContributedTo"`
			PullRequests struct {
				TotalCount int `json:"totalCount"`
			} `json:"pullRequests"`
			MergedPullRequests struct {
				TotalCount int `json:"totalCount"`
			} `json:"mergedPullRequests"`
			OpenIssues struct {
				TotalCount int `json:"totalCount"`
			} `json:"openIssues"`
			ClosedIssues struct {
				TotalCount int `json:"totalCount"`
			} `json:"closedIssues"`
			Followers struct {
				TotalCount int `json:"totalCount"`
			} `json:"followers"`
			RepositoryDiscussions struct {
				TotalCount int `json:"totalCount"`
			} `json:"repositoryDiscussions"`
			RepositoryDiscussionComments struct {
				TotalCount int `json:"totalCount"`
			} `json:"repositoryDiscussionComments"`
			Repositories struct {
				TotalCount int `json:"totalCount"`
				Nodes      []struct {
					Name       string `json:"name"`
					Stargazers struct {
						TotalCount int `json:"totalCount"`
					} `json:"stargazers"`
				} `json:"nodes"`
				PageInfo struct {
					HasNextPage bool   `json:"hasNextPage"`
					EndCursor   string `json:"endCursor"`
				} `json:"pageInfo"`
			} `json:"repositories"`
		} `json:"user"`
	} `json:"data"`
	Errors []struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"errors"`
}

// StatsFetcherParams represents parameters for fetching stats
type StatsFetcherParams struct {
	Username                  string
	IncludeMergedPullRequests bool
	IncludeDiscussions        bool
	IncludeDiscussionsAnswers bool
	StartTime                 *time.Time
}

// StatsFetcher fetches stats using GraphQL
func StatsFetcher(params StatsFetcherParams) (*GraphQLResponse, error) {
	var allRepos []struct {
		Name       string `json:"name"`
		Stargazers struct {
			TotalCount int `json:"totalCount"`
		} `json:"stargazers"`
	}

	hasNextPage := true
	endCursor := ""
	var firstResponse *GraphQLResponse

	for hasNextPage {
		variables := map[string]interface{}{
			"login":                     params.Username,
			"first":                     100,
			"after":                     endCursor,
			"includeMergedPullRequests": params.IncludeMergedPullRequests,
			"includeDiscussions":        params.IncludeDiscussions,
			"includeDiscussionsAnswers": params.IncludeDiscussionsAnswers,
		}

		if params.StartTime != nil {
			variables["startTime"] = params.StartTime.Format(time.RFC3339)
		}

		var query string
		if endCursor != "" {
			query = fmt.Sprintf(GraphQLReposQuery, GraphQLReposField)
		} else {
			query = fmt.Sprintf(GraphQLStatsQuery, GraphQLReposField)
		}

		req := common.GraphQLRequest{
			Query:     query,
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

		var graphqlResp GraphQLResponse
		if err := json.Unmarshal(body, &graphqlResp); err != nil {
			return nil, err
		}

		if len(graphqlResp.Errors) > 0 {
			return &graphqlResp, nil
		}

		if firstResponse == nil {
			firstResponse = &graphqlResp
		}

		repoNodes := graphqlResp.Data.User.Repositories.Nodes
		allRepos = append(allRepos, repoNodes...)

		// Check if we should continue fetching
		if os.Getenv("FETCH_MULTI_PAGE_STARS") != "true" {
			hasNextPage = false
		} else {
			repoNodesWithStars := 0
			for _, node := range repoNodes {
				if node.Stargazers.TotalCount != 0 {
					repoNodesWithStars++
				}
			}
			hasNextPage = len(repoNodes) == repoNodesWithStars && graphqlResp.Data.User.Repositories.PageInfo.HasNextPage
			endCursor = graphqlResp.Data.User.Repositories.PageInfo.EndCursor
		}
	}

	// Merge all repos into first response
	if firstResponse != nil {
		firstResponse.Data.User.Repositories.Nodes = allRepos
	}

	return firstResponse, nil
}

// FetchTotalCommits fetches total commits using REST API
func FetchTotalCommits(username string) (int, error) {
	url := fmt.Sprintf("%s?q=author:%s", restAPIURL, username)

	resp, err := common.Retryer(func(v interface{}, token string) (*http.Response, error) {
		return common.RequestREST("GET", url, token, nil)
	}, nil, 0)

	if err != nil {
		return 0, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()

	var result struct {
		TotalCount int `json:"total_count"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}

	return result.TotalCount, nil
}

// FetchStats fetches stats for a given username
func FetchStats(
	username string,
	includeAllCommits bool,
	excludeRepo []string,
	includeMergedPullRequests bool,
	includeDiscussions bool,
	includeDiscussionsAnswers bool,
	commitsYear *int,
) (*StatsData, error) {
	if username == "" {
		return nil, common.NewMissingParamError([]string{"username"}, "")
	}

	var startTime *time.Time
	if commitsYear != nil {
		t := time.Date(*commitsYear, 1, 1, 0, 0, 0, 0, time.UTC)
		startTime = &t
	}

	res, err := StatsFetcher(StatsFetcherParams{
		Username:                  username,
		IncludeMergedPullRequests: includeMergedPullRequests,
		IncludeDiscussions:        includeDiscussions,
		IncludeDiscussionsAnswers: includeDiscussionsAnswers,
		StartTime:                 startTime,
	})

	if err != nil {
		return nil, err
	}

	// Handle GraphQL errors
	if len(res.Errors) > 0 {
		if res.Errors[0].Type == "NOT_FOUND" {
			return nil, common.NewCustomError(
				res.Errors[0].Message,
				common.ErrorTypeUserNotFound,
			)
		}
		if res.Errors[0].Message != "" {
			msg := strings.Split(res.Errors[0].Message, "\n")[0]
			if len(msg) > 90 {
				msg = msg[:90]
			}
			return nil, common.NewCustomError(msg, res.Errors[0].Type)
		}
		return nil, common.NewCustomError(
			"Something went wrong while trying to retrieve the stats data using the GraphQL API.",
			common.ErrorTypeGraphQLError,
		)
	}

	user := res.Data.User
	stats := &StatsData{
		Name: user.Name,
	}

	if user.Name == "" {
		stats.Name = user.Login
	}

	// Fetch commits
	if includeAllCommits {
		totalCommits, err := FetchTotalCommits(username)
		if err != nil {
			return nil, err
		}
		stats.TotalCommits = totalCommits
	} else {
		stats.TotalCommits = user.Commits.TotalCommitContributions
	}

	stats.TotalPRs = user.PullRequests.TotalCount
	if includeMergedPullRequests {
		stats.TotalPRsMerged = user.MergedPullRequests.TotalCount
		if user.PullRequests.TotalCount > 0 {
			stats.MergedPRsPercentage = float64(user.MergedPullRequests.TotalCount) / float64(user.PullRequests.TotalCount) * 100
		}
	}
	stats.TotalReviews = user.Reviews.TotalPullRequestReviewContributions
	stats.TotalIssues = user.OpenIssues.TotalCount + user.ClosedIssues.TotalCount
	if includeDiscussions {
		stats.TotalDiscussionsStarted = user.RepositoryDiscussions.TotalCount
	}
	if includeDiscussionsAnswers {
		stats.TotalDiscussionsAnswered = user.RepositoryDiscussionComments.TotalCount
	}
	stats.ContributedTo = user.RepositoriesContributedTo.TotalCount

	// Calculate total stars (excluding excluded repos)
	allExcludedRepos := append(excludeRepo, common.GetExcludeRepositories()...)
	excludedSet := make(map[string]bool)
	for _, repo := range allExcludedRepos {
		excludedSet[repo] = true
	}

	totalStars := 0
	for _, repo := range user.Repositories.Nodes {
		if !excludedSet[repo.Name] {
			totalStars += repo.Stargazers.TotalCount
		}
	}
	stats.TotalStars = totalStars

	// Calculate rank
	rank := common.CalculateRank(common.RankParams{
		AllCommits: includeAllCommits,
		Commits:    float64(stats.TotalCommits),
		PRs:        float64(stats.TotalPRs),
		Issues:     float64(stats.TotalIssues),
		Reviews:    float64(stats.TotalReviews),
		Repos:      float64(user.Repositories.TotalCount),
		Stars:      float64(stats.TotalStars),
		Followers:  float64(user.Followers.TotalCount),
	})

	stats.Rank = RankData{
		Level:      rank.Level,
		Percentile: rank.Percentile,
	}

	return stats, nil
}
