package common

import (
	"fmt"
	"math"
	"strings"
)

// KFormatter formats a number with suffix k (thousands) precise to given decimal places
func KFormatter(num float64, precision *int) interface{} {
	abs := math.Abs(num)
	sign := 1.0
	if num < 0 {
		sign = -1.0
	}

	// If number is less than 1000, return the original number regardless of precision
	if abs < 1000 {
		return int(sign * abs)
	}

	if precision != nil && !math.IsNaN(float64(*precision)) {
		prec := *precision
		if prec < 0 {
			prec = 0
		}
		if prec > 2 {
			prec = 2
		}
		return fmt.Sprintf("%.*fk", prec, sign*(abs/1000))
	}

	return fmt.Sprintf("%.1fk", sign*(abs/1000))
}

// FormatBytes converts bytes to a human-readable string representation
func FormatBytes(bytes int64) (string, error) {
	if bytes < 0 {
		return "", fmt.Errorf("bytes must be a non-negative number")
	}

	if bytes == 0 {
		return "0 B", nil
	}

	sizes := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	base := 1024.0
	i := int(math.Floor(math.Log(float64(bytes)) / math.Log(base)))

	if i >= len(sizes) {
		return "", fmt.Errorf("bytes is too large to convert to a human-readable string")
	}

	return fmt.Sprintf("%.1f %s", float64(bytes)/math.Pow(base, float64(i)), sizes[i]), nil
}

// WrapTextMultiline splits text over multiple lines based on the card width
func WrapTextMultiline(text string, width int, maxLines int) []string {
	if width <= 0 {
		width = 59
	}
	if maxLines <= 0 {
		maxLines = 3
	}

	fullWidthComma := "，"
	encoded := EncodeHTML(text)
	isChinese := strings.Contains(encoded, fullWidthComma)

	var wrapped []string

	if isChinese {
		wrapped = strings.Split(encoded, fullWidthComma)
	} else {
		// Simple word wrap implementation
		words := strings.Fields(encoded)
		var line strings.Builder
		for _, word := range words {
			if line.Len()+len(word)+1 > width && line.Len() > 0 {
				wrapped = append(wrapped, strings.TrimSpace(line.String()))
				line.Reset()
			}
			if line.Len() > 0 {
				line.WriteString(" ")
			}
			line.WriteString(word)
		}
		if line.Len() > 0 {
			wrapped = append(wrapped, strings.TrimSpace(line.String()))
		}
	}

	// Trim and limit to maxLines
	var lines []string
	for i, line := range wrapped {
		if i >= maxLines {
			break
		}
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			lines = append(lines, trimmed)
		}
	}

	// Add "..." to the last line if text exceeds maxLines
	if len(wrapped) > maxLines && len(lines) > 0 {
		lines[len(lines)-1] += "..."
	}

	return lines
}

// truncateLineWithEllipsis truncates a line to fit within maxWidth and adds ellipsis
func truncateLineWithEllipsis(line string, maxWidth float64, fontSize float64) string {
	ellipsisWidth := MeasureText("...", fontSize)
	maxContentWidth := maxWidth - ellipsisWidth

	// If line already fits, return as is
	if MeasureText(line, fontSize) <= maxWidth {
		return line
	}

	// Use binary search to find the optimal truncation point
	// Start with a reasonable estimate
	runes := []rune(line)
	left := 0
	right := len(runes)
	best := 0

	// Binary search for the best truncation point
	for left <= right {
		mid := (left + right) / 2
		if mid >= len(runes) {
			mid = len(runes) - 1
		}
		if mid < 0 {
			break
		}

		testStr := string(runes[:mid])
		testWidth := MeasureText(testStr, fontSize)

		if testWidth <= maxContentWidth {
			best = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// If we found a good point, use it
	if best > 0 {
		result := string(runes[:best]) + "..."
		// Final check to ensure it fits
		if MeasureText(result, fontSize) <= maxWidth {
			return result
		}
		// If still too long, reduce by one character
		if best > 1 {
			result = string(runes[:best-1]) + "..."
			if MeasureText(result, fontSize) <= maxWidth {
				return result
			}
		}
	}

	// Fallback: try removing characters one by one
	for len(line) > 0 {
		testStr := line + "..."
		if MeasureText(testStr, fontSize) <= maxWidth {
			return testStr
		}
		if len(line) > 1 {
			line = line[:len(line)-1]
		} else {
			return "..."
		}
	}
	return "..."
}

// WrapTextByPixelWidth splits text over multiple lines based on pixel width
// This function uses actual text measurement to ensure proper wrapping
func WrapTextByPixelWidth(text string, maxWidth float64, fontSize float64, maxLines int) []string {
	if maxWidth <= 0 {
		maxWidth = 350
	}
	if fontSize <= 0 {
		fontSize = 13
	}
	if maxLines <= 0 {
		maxLines = 3
	}

	// Work with original text (not encoded) for accurate measurement
	// We'll encode each line after wrapping
	words := strings.Fields(text)
	if len(words) == 0 {
		return []string{EncodeHTML("")}
	}

	var lines []string
	var currentLine strings.Builder
	lastProcessedIndex := -1

	for i, word := range words {
		lastProcessedIndex = i

		// Check if current line already exceeds width (before adding new word)
		// This handles the case where the line itself is too long
		// Use a stricter check: if line is >= 95% of maxWidth, save it to prevent overflow
		if currentLine.Len() > 0 {
			currentLineStr := currentLine.String()
			currentLineWidth := MeasureText(currentLineStr, fontSize)
			// Use 95% threshold to ensure we don't exceed
			if currentLineWidth >= maxWidth*0.95 {
				// Current line is close to or exceeds width, save it and start new line
				currentLineStr = strings.TrimSpace(currentLineStr)
				currentLineStr = truncateLineWithEllipsis(currentLineStr, maxWidth, fontSize)
				lines = append(lines, currentLineStr)
				currentLine.Reset()

				// Check if we've reached max lines
				if len(lines) >= maxLines {
					// Encode all lines and return
					encodedLines := make([]string, len(lines))
					for j, line := range lines {
						encodedLines[j] = EncodeHTML(line)
					}
					return encodedLines
				}
			}
		}

		// Check if we've reached max lines
		if len(lines) >= maxLines {
			// Need to truncate the last line if there's more content
			if currentLine.Len() > 0 {
				currentLineStr := strings.TrimSpace(currentLine.String())
				currentLineStr = truncateLineWithEllipsis(currentLineStr, maxWidth, fontSize)
				lines = append(lines, currentLineStr)
			}
			// Check if there are remaining words
			remainingWords := words[lastProcessedIndex+1:]
			if len(remainingWords) > 0 && len(lines) > 0 {
				// Truncate last line with ellipsis (remove existing ellipsis first if present)
				lastLine := strings.TrimSuffix(lines[len(lines)-1], "...")
				lastLine = truncateLineWithEllipsis(lastLine, maxWidth, fontSize)
				lines[len(lines)-1] = lastLine
			}
			// Encode all lines and return
			encodedLines := make([]string, len(lines))
			for j, line := range lines {
				encodedLines[j] = EncodeHTML(line)
			}
			return encodedLines
		}

		// Build test line to check if adding this word would exceed width
		testLine := currentLine.String()
		if testLine != "" {
			testLine += " " + word
		} else {
			testLine = word
		}

		lineWidth := MeasureText(testLine, fontSize)

		// If adding this word would exceed the width, save current line and start new one
		// Use 85% threshold to ensure we wrap earlier and don't exceed
		threshold := maxWidth * 0.85
		if lineWidth > threshold {
			// Save current line if it has content
			if currentLine.Len() > 0 {
				currentLineStr := strings.TrimSpace(currentLine.String())
				// Ensure current line doesn't exceed width (truncate if needed)
				currentLineStr = truncateLineWithEllipsis(currentLineStr, maxWidth, fontSize)
				lines = append(lines, currentLineStr)
				currentLine.Reset()

				// Check if we've reached max lines
				if len(lines) >= maxLines {
					// Truncate last line if there are remaining words
					if len(words) > i && len(lines) > 0 {
						lastLine := lines[len(lines)-1]
						// Remove ellipsis if present to recalculate
						lastLine = strings.TrimSuffix(lastLine, "...")
						lastLine = truncateLineWithEllipsis(lastLine, maxWidth, fontSize)
						lines[len(lines)-1] = lastLine
					}
					// Encode all lines and return
					encodedLines := make([]string, len(lines))
					for j, line := range lines {
						encodedLines[j] = EncodeHTML(line)
					}
					return encodedLines
				}
			}

			// Check if the word itself exceeds width
			wordWidth := MeasureText(word, fontSize)
			if wordWidth > maxWidth {
				// Word is too long, need to truncate it
				if len(lines) >= maxLines-1 {
					// Last line, truncate with ellipsis
					truncated := truncateLineWithEllipsis(word, maxWidth, fontSize)
					lines = append(lines, truncated)
					// Encode all lines and return
					encodedLines := make([]string, len(lines))
					for j, line := range lines {
						encodedLines[j] = EncodeHTML(line)
					}
					return encodedLines
				}
				// Not last line, add truncated word
				truncated := truncateLineWithEllipsis(word, maxWidth, fontSize)
				lines = append(lines, truncated)
				currentLine.Reset()
			} else {
				// Word fits on new line, add it
				currentLine.WriteString(word)
			}
			continue
		}

		// Word fits on current line, add it
		if currentLine.Len() > 0 {
			currentLine.WriteString(" ")
		}
		currentLine.WriteString(word)

		// Critical check: verify current line doesn't exceed width after adding the word
		// This is the most important check to ensure we never exceed maxWidth
		// Use 85% threshold to ensure we wrap earlier
		currentLineStr := currentLine.String()
		currentLineWidth := MeasureText(currentLineStr, fontSize)
		threshold = maxWidth * 0.85
		if currentLineWidth > threshold {
			// Current line exceeds width, need to save it and start a new line
			// Remove the last word we just added
			wordsInLine := strings.Fields(currentLineStr)
			if len(wordsInLine) > 1 {
				// Rebuild line without the last word
				currentLine.Reset()
				for j := 0; j < len(wordsInLine)-1; j++ {
					if j > 0 {
						currentLine.WriteString(" ")
					}
					currentLine.WriteString(wordsInLine[j])
				}
				currentLineStr = strings.TrimSpace(currentLine.String())
				currentLineStr = truncateLineWithEllipsis(currentLineStr, maxWidth, fontSize)
				lines = append(lines, currentLineStr)
				currentLine.Reset()

				// Check if we've reached max lines
				if len(lines) >= maxLines {
					// Encode all lines and return
					encodedLines := make([]string, len(lines))
					for j, line := range lines {
						encodedLines[j] = EncodeHTML(line)
					}
					return encodedLines
				}

				// Add the word that didn't fit to a new line
				currentLine.WriteString(wordsInLine[len(wordsInLine)-1])
			} else {
				// Single word that's too long, truncate it
				currentLine.Reset()
				truncated := truncateLineWithEllipsis(word, maxWidth, fontSize)
				lines = append(lines, truncated)

				// Check if we've reached max lines
				if len(lines) >= maxLines {
					// Encode all lines and return
					encodedLines := make([]string, len(lines))
					for j, line := range lines {
						encodedLines[j] = EncodeHTML(line)
					}
					return encodedLines
				}
			}
		}
	}

	// Add remaining content
	if currentLine.Len() > 0 {
		lineStr := strings.TrimSpace(currentLine.String())
		// Ensure line doesn't exceed width (truncate if needed)
		lineStr = truncateLineWithEllipsis(lineStr, maxWidth, fontSize)

		if len(lines) < maxLines {
			lines = append(lines, lineStr)
		} else {
			// We've reached max lines, check if there's more content
			remainingWords := words[lastProcessedIndex+1:]
			if len(remainingWords) > 0 && len(lines) > 0 {
				// Need to add ellipsis to last line
				lastLine := strings.TrimSuffix(lines[len(lines)-1], "...")
				lastLine = truncateLineWithEllipsis(lastLine, maxWidth, fontSize)
				lines[len(lines)-1] = lastLine
			}
		}
	}

	// Encode all lines before returning
	encodedLines := make([]string, len(lines))
	for i, line := range lines {
		encodedLines[i] = EncodeHTML(line)
	}
	return encodedLines
}
