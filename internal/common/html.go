package common

import (
	"strconv"
	"strings"
)

// EncodeHTML encodes string as HTML
// See https://stackoverflow.com/a/48073476/10629172
// Avoids double-encoding HTML entities by checking if '&' is followed by '#'
func EncodeHTML(str string) string {
	// First, remove backspace characters
	str = strings.ReplaceAll(str, "\u0008", "")

	var result strings.Builder
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		r := runes[i]

		// Check if character needs encoding (Unicode range \u00A0-\u9999 or <>&)
		if (r >= 0x00A0 && r <= 0x9999) || r == '<' || r == '>' {
			result.WriteString("&#")
			result.WriteString(strconv.Itoa(int(r)))
			result.WriteString(";")
		} else if r == '&' {
			// For '&', check if it's part of an existing HTML entity (followed by '#')
			// If so, don't encode it to avoid double-encoding
			if i+1 < len(runes) && runes[i+1] == '#' {
				result.WriteRune(r)
			} else {
				result.WriteString("&#38;")
			}
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}
