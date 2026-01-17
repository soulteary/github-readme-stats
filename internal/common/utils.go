package common

import (
	"math"
	"strconv"
	"strings"
	"time"
)

// ParseBoolean returns boolean if value is either "true" or "false" else returns nil
func ParseBoolean(value interface{}) *bool {
	if b, ok := value.(bool); ok {
		return &b
	}

	if str, ok := value.(string); ok {
		lower := strings.ToLower(str)
		if lower == "true" {
			b := true
			return &b
		} else if lower == "false" {
			b := false
			return &b
		}
	}
	return nil
}

// ParseArray parses string to array of strings
func ParseArray(str string) []string {
	if str == "" {
		return []string{}
	}
	return strings.Split(str, ",")
}

// ClampValue clamps the given number between the given range
func ClampValue(number interface{}, min, max float64) float64 {
	var num float64
	switch v := number.(type) {
	case int:
		num = float64(v)
	case int64:
		num = float64(v)
	case float64:
		num = v
	case string:
		parsed, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return min
		}
		num = parsed
	default:
		return min
	}

	if math.IsNaN(num) {
		return min
	}
	return math.Max(min, math.Min(num, max))
}

// LowercaseTrim lowercases and trims string
func LowercaseTrim(name string) string {
	return strings.ToLower(strings.TrimSpace(name))
}

// ChunkArray splits array into chunks
func ChunkArray[T any](arr []T, perChunk int) [][]T {
	if perChunk <= 0 {
		return [][]T{arr}
	}

	var result [][]T
	for i := 0; i < len(arr); i += perChunk {
		end := i + perChunk
		if end > len(arr) {
			end = len(arr)
		}
		result = append(result, arr[i:end])
	}
	return result
}

// DateDiff returns the difference in minutes between two dates
func DateDiff(d1, d2 string) (int, error) {
	// Parse dates
	date1, err := time.Parse(time.RFC3339, d1)
	if err != nil {
		// Try other common formats
		formats := []string{
			"2006-01-02T15:04:05Z07:00",
			"2006-01-02 15:04:05",
			"2006-01-02",
		}
		for _, format := range formats {
			if parsed, err := time.Parse(format, d1); err == nil {
				date1 = parsed
				break
			}
		}
		if date1.IsZero() {
			return 0, err
		}
	}

	date2, err := time.Parse(time.RFC3339, d2)
	if err != nil {
		// Try other common formats
		formats := []string{
			"2006-01-02T15:04:05Z07:00",
			"2006-01-02 15:04:05",
			"2006-01-02",
		}
		for _, format := range formats {
			if parsed, err := time.Parse(format, d2); err == nil {
				date2 = parsed
				break
			}
		}
		if date2.IsZero() {
			return 0, err
		}
	}

	diff := date1.Sub(date2)
	return int(diff.Minutes()), nil
}
