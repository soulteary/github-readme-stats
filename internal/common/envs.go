package common

import (
	"os"
	"strings"
)

var (
	whitelist           []string
	gistWhitelist       []string
	excludeRepositories []string
)

func init() {
	loadEnvVars()
}

func loadEnvVars() {
	if wl := os.Getenv("WHITELIST"); wl != "" {
		whitelist = strings.Split(wl, ",")
		for i := range whitelist {
			whitelist[i] = strings.TrimSpace(whitelist[i])
		}
	} else {
		whitelist = []string{}
	}

	if gwl := os.Getenv("GIST_WHITELIST"); gwl != "" {
		gistWhitelist = strings.Split(gwl, ",")
		for i := range gistWhitelist {
			gistWhitelist[i] = strings.TrimSpace(gistWhitelist[i])
		}
	} else {
		gistWhitelist = []string{}
	}

	if er := os.Getenv("EXCLUDE_REPO"); er != "" {
		excludeRepositories = strings.Split(er, ",")
		for i := range excludeRepositories {
			excludeRepositories[i] = strings.TrimSpace(excludeRepositories[i])
		}
	} else {
		excludeRepositories = []string{}
	}
}

// GetWhitelist returns the whitelist
func GetWhitelist() []string {
	return whitelist
}

// GetGistWhitelist returns the gist whitelist
func GetGistWhitelist() []string {
	return gistWhitelist
}

// GetExcludeRepositories returns the list of excluded repositories
func GetExcludeRepositories() []string {
	return excludeRepositories
}

// IsWhitelisted checks if an ID is in the whitelist
func IsWhitelisted(id string, isGist bool) bool {
	var list []string
	if isGist {
		list = gistWhitelist
	} else {
		list = whitelist
	}

	if len(list) == 0 {
		return true // No whitelist means all are allowed
	}

	for _, item := range list {
		if strings.TrimSpace(item) == id {
			return true
		}
	}
	return false
}
