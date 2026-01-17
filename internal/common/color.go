package common

import (
	"regexp"
	"strings"

	"github.com/soulteary/github-readme-stats/internal/themes"
)

var hexColorRegex = regexp.MustCompile(`^([A-Fa-f0-9]{8}|[A-Fa-f0-9]{6}|[A-Fa-f0-9]{3}|[A-Fa-f0-9]{4})$`)

// IsValidHexColor checks if a string is a valid hex color
func IsValidHexColor(hexColor string) bool {
	return hexColorRegex.MatchString(hexColor)
}

// IsValidGradient checks if the given array of colors is a valid gradient
func IsValidGradient(colors []string) bool {
	if len(colors) <= 2 {
		return false
	}
	for i := 1; i < len(colors); i++ {
		if !IsValidHexColor(colors[i]) {
			return false
		}
	}
	return true
}

// FallbackColor retrieves a gradient if color has more than one valid hex codes else a single color
func FallbackColor(color string, fallbackColor interface{}) interface{} {
	// If color is empty, return fallback with # prefix if it's a string
	if color == "" {
		if fallbackStr, ok := fallbackColor.(string); ok {
			if !strings.HasPrefix(fallbackStr, "#") {
				return "#" + fallbackStr
			}
			return fallbackStr
		}
		return fallbackColor
	}

	// If color already has # prefix, remove it for validation
	colorWithoutHash := color
	if strings.HasPrefix(color, "#") {
		colorWithoutHash = color[1:]
	}

	var gradient []string
	colors := strings.Split(colorWithoutHash, ",")
	if len(colors) > 1 && IsValidGradient(colors) {
		gradient = colors
	}

	if gradient != nil {
		return gradient
	}

	if IsValidHexColor(colorWithoutHash) {
		// If original color had #, return as is, otherwise add #
		if strings.HasPrefix(color, "#") {
			return color
		}
		return "#" + colorWithoutHash
	}

	// If fallback is a string and doesn't start with #, add it
	if fallbackStr, ok := fallbackColor.(string); ok {
		if !strings.HasPrefix(fallbackStr, "#") {
			return "#" + fallbackStr
		}
		return fallbackStr
	}

	return fallbackColor
}

// CardColors represents card colors
type CardColors struct {
	TitleColor  string
	IconColor   string
	TextColor   string
	BgColor     interface{} // string or []string for gradient
	BorderColor string
	RingColor   string
}

// GetCardColors returns theme based colors with proper overrides and defaults
func GetCardColors(args map[string]string) CardColors {
	defaultTheme := themes.Themes["default"]
	themeName := args["theme"]
	isThemeProvided := themeName != ""

	var selectedTheme themes.Theme
	if isThemeProvided {
		selectedTheme = themes.GetTheme(themeName)
	} else {
		selectedTheme = defaultTheme
	}

	defaultBorderColor := selectedTheme.BorderColor
	if defaultBorderColor == "" {
		defaultBorderColor = defaultTheme.BorderColor
	}

	// Get the color provided by the user else the theme color
	// finally if both colors are invalid fallback to default theme
	titleColor := FallbackColor(
		args["title_color"],
		selectedTheme.TitleColor,
	)
	if titleColor == nil {
		titleColor = "#" + defaultTheme.TitleColor
	}

	// Get ring color - use titleColor as fallback (like JavaScript version)
	ringColorInput := args["ring_color"]
	if ringColorInput == "" {
		ringColorInput = selectedTheme.RingColor
	}
	ringColor := FallbackColor(
		ringColorInput,
		titleColor, // Use titleColor as fallback, not selectedTheme.RingColor
	)
	if ringColor == nil {
		ringColor = titleColor
	} else if ringColorStr, ok := ringColor.(string); ok && (ringColorStr == "" || ringColorStr == "#") {
		ringColor = titleColor
	}

	iconColor := FallbackColor(
		args["icon_color"],
		selectedTheme.IconColor,
	)
	if iconColor == nil {
		iconColor = "#" + defaultTheme.IconColor
	}

	textColor := FallbackColor(
		args["text_color"],
		selectedTheme.TextColor,
	)
	if textColor == nil {
		textColor = "#" + defaultTheme.TextColor
	}

	bgColor := FallbackColor(
		args["bg_color"],
		selectedTheme.BgColor,
	)
	if bgColor == nil {
		bgColor = "#" + defaultTheme.BgColor
	}
	// Ensure bgColor string has # prefix (gradients are arrays, skip them)
	if bgColorStr, ok := bgColor.(string); ok {
		if bgColorStr != "" && !strings.HasPrefix(bgColorStr, "#") {
			bgColor = "#" + bgColorStr
		}
	}

	borderColor := FallbackColor(
		args["border_color"],
		defaultBorderColor,
	)
	if borderColor == nil {
		borderColor = "#" + defaultBorderColor
	}

	// Convert to strings and ensure they have # prefix
	titleColorStr, _ := titleColor.(string)
	if titleColorStr != "" && !strings.HasPrefix(titleColorStr, "#") {
		titleColorStr = "#" + titleColorStr
	}

	ringColorStr, _ := ringColor.(string)
	if ringColorStr == "" || ringColorStr == "#" {
		// If ringColor is empty or invalid, use titleColor as fallback
		ringColorStr = titleColorStr
	} else if !strings.HasPrefix(ringColorStr, "#") {
		ringColorStr = "#" + ringColorStr
	}

	iconColorStr, _ := iconColor.(string)
	if iconColorStr != "" && !strings.HasPrefix(iconColorStr, "#") {
		iconColorStr = "#" + iconColorStr
	}

	textColorStr, _ := textColor.(string)
	if textColorStr != "" && !strings.HasPrefix(textColorStr, "#") {
		textColorStr = "#" + textColorStr
	}

	borderColorStr, _ := borderColor.(string)
	if borderColorStr != "" && !strings.HasPrefix(borderColorStr, "#") {
		borderColorStr = "#" + borderColorStr
	}

	return CardColors{
		TitleColor:  titleColorStr,
		IconColor:   iconColorStr,
		TextColor:   textColorStr,
		BgColor:     bgColor,
		BorderColor: borderColorStr,
		RingColor:   ringColorStr,
	}
}
