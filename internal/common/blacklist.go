package common

var blacklist = []string{
	"renovate-bot",
	"technote-space",
	"sw-yx",
	"YourUsername",
	"[YourUsername]",
}

// IsBlacklisted checks if a username is blacklisted
func IsBlacklisted(username string) bool {
	for _, item := range blacklist {
		if item == username {
			return true
		}
	}
	return false
}
