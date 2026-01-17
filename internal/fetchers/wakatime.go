package fetchers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/soulteary/github-readme-stats/internal/common"
)

// FetchWakaTimeStats fetches WakaTime statistics
func FetchWakaTimeStats(username, apiDomain string) (*WakaTimeData, error) {
	if username == "" {
		return nil, common.NewMissingParamError([]string{"username"}, "")
	}

	if apiDomain == "" {
		apiDomain = "wakatime.com"
	}
	apiDomain = strings.TrimSuffix(apiDomain, "/")

	url := fmt.Sprintf("https://%s/api/v1/users/%s/stats?is_including_today=true", apiDomain, username)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, common.NewCustomError(
			fmt.Sprintf("Could not resolve to a User with the login of '%s'", username),
			common.ErrorTypeWakaTimeError,
		)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Data struct {
			Languages []struct {
				Name    string  `json:"name"`
				Percent float64 `json:"percent"`
				Text    string  `json:"text"`
				Hours   int     `json:"hours"`
				Minutes int     `json:"minutes"`
			} `json:"languages"`
			IsCodingActivityVisible bool   `json:"is_coding_activity_visible"`
			IsOtherUsageVisible     bool   `json:"is_other_usage_visible"`
			Range                   string `json:"range"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	// Convert to WakaTimeData
	var languages []WakaTimeLanguage
	for _, lang := range result.Data.Languages {
		languages = append(languages, WakaTimeLanguage{
			Name:    lang.Name,
			Percent: lang.Percent,
			Time:    lang.Text,
			Hours:   lang.Hours,
			Minutes: lang.Minutes,
		})
	}

	return &WakaTimeData{
		Languages:               languages,
		TotalTime:               "", // Will be calculated if needed
		IsCodingActivityVisible: result.Data.IsCodingActivityVisible,
		IsOtherUsageVisible:     result.Data.IsOtherUsageVisible,
		Range:                   result.Data.Range,
	}, nil
}

// FetchWakaTimeStatsFromFile fetches WakaTime statistics from a test data file
func FetchWakaTimeStatsFromFile(testDataPath string) (*WakaTimeData, error) {
	// Try to find test data file
	var dataPath string
	if testDataPath != "" {
		dataPath = testDataPath
	} else {
		// Try to find testdata directory relative to current working directory
		wd, err := os.Getwd()
		if err == nil {
			possiblePaths := []string{
				filepath.Join(wd, "testdata", "wakatime.example.json"),
				filepath.Join(wd, "..", "testdata", "wakatime.example.json"),
				filepath.Join(wd, "github-readme-stats", "testdata", "wakatime.example.json"),
			}
			for _, path := range possiblePaths {
				if _, err := os.Stat(path); err == nil {
					dataPath = path
					break
				}
			}
		}
	}

	if dataPath == "" {
		return nil, fmt.Errorf("test data file not found")
	}

	body, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read test data file: %w", err)
	}

	var result struct {
		Data struct {
			Languages []struct {
				Name    string  `json:"name"`
				Percent float64 `json:"percent"`
				Text    string  `json:"text"`
				Hours   int     `json:"hours"`
				Minutes int     `json:"minutes"`
			} `json:"languages"`
			IsCodingActivityVisible bool   `json:"is_coding_activity_visible"`
			IsOtherUsageVisible     bool   `json:"is_other_usage_visible"`
			Range                   string `json:"range"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse test data: %w", err)
	}

	// Convert to WakaTimeData
	var languages []WakaTimeLanguage
	for _, lang := range result.Data.Languages {
		languages = append(languages, WakaTimeLanguage{
			Name:    lang.Name,
			Percent: lang.Percent,
			Time:    lang.Text,
			Hours:   lang.Hours,
			Minutes: lang.Minutes,
		})
	}

	return &WakaTimeData{
		Languages:               languages,
		TotalTime:               "", // Will be calculated if needed
		IsCodingActivityVisible: result.Data.IsCodingActivityVisible,
		IsOtherUsageVisible:     result.Data.IsOtherUsageVisible,
		Range:                   result.Data.Range,
	}, nil
}
