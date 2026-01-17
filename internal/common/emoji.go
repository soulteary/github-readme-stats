package common

import (
	"regexp"
	"strings"
)

// EmojiMap contains common emoji mappings
// This is a subset of common emojis used in GitHub descriptions
var EmojiMap = map[string]string{
	":heart:":                     "❤️",
	":sparkles:":                  "✨",
	":star:":                      "⭐",
	":rocket:":                    "🚀",
	":fire:":                      "🔥",
	":zap:":                       "⚡",
	":tada:":                      "🎉",
	":white_check_mark:":          "✅",
	":x:":                         "❌",
	":warning:":                   "⚠️",
	":bug:":                       "🐛",
	":art:":                       "🎨",
	":memo:":                      "📝",
	":lipstick:":                  "💄",
	":rotating_light:":            "🚨",
	":construction:":              "🚧",
	":green_heart:":               "💚",
	":arrow_down:":                "⬇️",
	":arrow_up:":                  "⬆️",
	":pushpin:":                   "📌",
	":construction_worker:":       "👷",
	":chart_with_upwards_trend:":  "📈",
	":hammer:":                    "🔨",
	":package:":                   "📦",
	":bento:":                     "🍱",
	":ok_hand:":                   "👌",
	":boom:":                      "💥",
	":wastebasket:":               "🗑️",
	":lock:":                      "🔒",
	":apple:":                     "🍎",
	":penguin:":                   "🐧",
	":checkered_flag:":            "🏁",
	":robot:":                     "🤖",
	":green_apple:":               "🍏",
	":bookmark:":                  "🔖",
	":recycle:":                   "♻️",
	":white_circle:":              "⚪",
	":heavy_minus_sign:":          "➖",
	":heavy_plus_sign:":           "➕",
	":wrench:":                    "🔧",
	":globe_with_meridians:":      "🌐",
	":pencil2:":                   "✏️",
	":hankey:":                    "💩",
	":rewind:":                    "⏪",
	":twisted_rightwards_arrows:": "🔀",
	":truck:":                     "🚚",
	":page_facing_up:":            "📄",
	":busts_in_silhouette:":       "👥",
	":children_crossing:":         "🚸",
	":building_construction:":     "🏗️",
	":iphone:":                    "📱",
	":clown_face:":                "🤡",
	":egg:":                       "🥚",
	":see_no_evil:":               "🙈",
	":camera_flash:":              "📸",
	":alembic:":                   "⚗️",
	":mag:":                       "🔍",
	":wheel_of_dharma:":           "☸️",
	":label:":                     "🏷️",
	":seedling:":                  "🌱",
	":triangular_flag_on_post:":   "🚩",
	":goal_net:":                  "🥅",
	":dizzy:":                     "💫",
	":monocle_face:":              "🧐",
	":stethoscope:":               "🩺",
	":bricks:":                    "🧱",
	":technologist:":              "🧑‍💻",
	":money_with_wings:":          "💸",
}

// ParseEmojis parses emoji from string, replacing :emoji: format with actual emoji
func ParseEmojis(str string) string {
	if str == "" {
		return ""
	}

	// Regular expression to match :emoji: format
	re := regexp.MustCompile(`:\w+:`)

	result := re.ReplaceAllStringFunc(str, func(match string) string {
		// Check if emoji exists in map
		if emoji, ok := EmojiMap[strings.ToLower(match)]; ok {
			return emoji
		}
		// If not found, return empty string (remove the :emoji: text)
		return ""
	})

	return result
}
