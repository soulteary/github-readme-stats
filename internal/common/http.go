package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GraphQLRequest represents a GraphQL request
type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// GraphQLResponse represents a GraphQL response
type GraphQLResponse struct {
	Data   interface{} `json:"data"`
	Errors []struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"errors"`
}

// Request sends a GraphQL request to GitHub API
func Request(data GraphQLRequest, token string) (*http.Response, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.github.com/graphql", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// RequestREST sends a REST API request to GitHub
func RequestREST(method, url string, token string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/vnd.github.cloak-preview")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
