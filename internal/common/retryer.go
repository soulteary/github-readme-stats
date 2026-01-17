package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// FetcherFunction represents a function that fetches data
type FetcherFunction func(variables interface{}, token string) (*http.Response, error)

// GetPATCount returns the number of GitHub API tokens available
func GetPATCount() int {
	count := 0
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, "PAT_") {
			count++
		}
	}
	return count
}

// GetRetries returns the number of retries based on environment
func GetRetries() int {
	if os.Getenv("NODE_ENV") == "test" {
		return 7
	}
	return GetPATCount()
}

// GetPAT returns a PAT token by index (1-based)
func GetPAT(index int) string {
	return os.Getenv(fmt.Sprintf("PAT_%d", index))
}

// Retryer tries to execute the fetcher function until it succeeds or max retries is reached
func Retryer(fetcher FetcherFunction, variables interface{}, retries int) (*http.Response, error) {
	maxRetries := GetRetries()
	if maxRetries == 0 {
		return nil, NewCustomError("No GitHub API tokens found", ErrorTypeNoTokens)
	}

	if retries > maxRetries {
		return nil, NewCustomError(
			"Downtime due to GitHub API rate limiting",
			ErrorTypeMaxRetry,
		)
	}

	token := GetPAT(retries + 1)
	if token == "" {
		return nil, NewCustomError("No GitHub API tokens found", ErrorTypeNoTokens)
	}

	response, err := fetcher(variables, token)
	if err != nil {
		// Network/unexpected error → let caller treat as failure
		return nil, err
	}

	// Check for rate limiting in GraphQL errors
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	response.Body.Close()

	var graphqlResp GraphQLResponse
	if err := json.Unmarshal(body, &graphqlResp); err == nil {
		if len(graphqlResp.Errors) > 0 {
			errorType := graphqlResp.Errors[0].Type
			errorMsg := graphqlResp.Errors[0].Message
			isRateLimited := errorType == "RATE_LIMITED" || regexp.MustCompile(`(?i)rate limit`).MatchString(errorMsg)

			if isRateLimited {
				// Retry with next token
				return Retryer(fetcher, variables, retries+1)
			}
		}
	}

	// Check for bad credentials or account suspended in REST API errors
	if response.StatusCode == 401 || response.StatusCode == 403 {
		bodyStr := string(body)
		isBadCredential := strings.Contains(bodyStr, "Bad credentials")
		isAccountSuspended := strings.Contains(bodyStr, "Sorry. Your account was suspended.")

		if isBadCredential || isAccountSuspended {
			// Retry with next token
			return Retryer(fetcher, variables, retries+1)
		}
	}

	// Create a new response reader for the body
	response.Body = io.NopCloser(strings.NewReader(string(body)))
	return response, nil
}
