package common

import (
	"net/http"
)

const (
	notWhitelistedUsernameMessage = "This username is not whitelisted"
	notWhitelistedGistMessage     = "This gist ID is not whitelisted"
	blacklistedMessage            = "This username is blacklisted"
)

// GuardAccessResult represents the result of access guard check
type GuardAccessResult struct {
	IsPassed bool
	Response string
}

// GuardAccess guards access using whitelist/blacklist
func GuardAccess(w http.ResponseWriter, id, accessType string, colors map[string]string) GuardAccessResult {
	if accessType != "username" && accessType != "gist" && accessType != "wakatime" {
		return GuardAccessResult{
			IsPassed: false,
			Response: RenderError(ErrorOptions{
				Message: "Invalid access type",
			}),
		}
	}

	isGist := accessType == "gist"
	currentWhitelist := GetGistWhitelist()
	notWhitelistedMsg := notWhitelistedUsernameMessage

	if isGist {
		currentWhitelist = GetGistWhitelist()
		notWhitelistedMsg = notWhitelistedGistMessage
	} else {
		currentWhitelist = GetWhitelist()
	}

	// Check whitelist
	if len(currentWhitelist) > 0 && !IsWhitelisted(id, isGist) {
		response := RenderError(ErrorOptions{
			Message:          notWhitelistedMsg,
			SecondaryMessage: "Please deploy your own instance",
			RenderOptions:    colors,
			ShowRepoLink:     false,
		})
		w.Header().Set("Content-Type", "image/svg+xml")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(response))
		return GuardAccessResult{
			IsPassed: false,
			Response: response,
		}
	}

	// Check blacklist (only for usernames)
	if accessType == "username" && len(GetWhitelist()) == 0 && IsBlacklisted(id) {
		response := RenderError(ErrorOptions{
			Message:          blacklistedMessage,
			SecondaryMessage: "Please deploy your own instance",
			RenderOptions:    colors,
			ShowRepoLink:     false,
		})
		w.Header().Set("Content-Type", "image/svg+xml")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(response))
		return GuardAccessResult{
			IsPassed: false,
			Response: response,
		}
	}

	return GuardAccessResult{
		IsPassed: true,
	}
}
