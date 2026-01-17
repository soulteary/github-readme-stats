package common

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const (
	Min  = 60
	Hour = 60 * Min
	Day  = 24 * Hour
)

// Durations
const (
	OneMinute      = Min
	FiveMinutes    = 5 * Min
	TenMinutes     = 10 * Min
	FifteenMinutes = 15 * Min
	ThirtyMinutes  = 30 * Min
	TwoHours       = 2 * Hour
	FourHours      = 4 * Hour
	SixHours       = 6 * Hour
	EightHours     = 8 * Hour
	TwelveHours    = 12 * Hour
	OneDay         = Day
	TwoDay         = 2 * Day
	SixDay         = 6 * Day
	TenDay         = 10 * Day
)

// CacheTTL contains default cache TTL values
type CacheTTL struct {
	Default int
	Min     int
	Max     int
}

var (
	StatsCardTTL = CacheTTL{
		Default: OneDay,
		Min:     TwelveHours,
		Max:     TwoDay,
	}
	TopLangsCardTTL = CacheTTL{
		Default: SixDay,
		Min:     TwoDay,
		Max:     TenDay,
	}
	PinCardTTL = CacheTTL{
		Default: TenDay,
		Min:     OneDay,
		Max:     TenDay,
	}
	GistCardTTL = CacheTTL{
		Default: TwoDay,
		Min:     OneDay,
		Max:     TenDay,
	}
	WakaTimeCardTTL = CacheTTL{
		Default: OneDay,
		Min:     TwelveHours,
		Max:     TwoDay,
	}
	ErrorTTL = TenMinutes
)

// ResolveCacheSeconds resolves the cache seconds based on requested, default, min, and max values
func ResolveCacheSeconds(requested int, def, min, max int) int {
	cacheSeconds := ClampValue(requested, float64(min), float64(max))
	if requested == 0 || requested < min {
		cacheSeconds = float64(def)
	}

	// Check environment variable
	if envCacheSeconds := os.Getenv("CACHE_SECONDS"); envCacheSeconds != "" {
		if envVal, err := strconv.Atoi(envCacheSeconds); err == nil {
			cacheSeconds = float64(envVal)
		}
	}

	return int(cacheSeconds)
}

// DisableCaching disables caching by setting appropriate headers
func DisableCaching(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate, max-age=0, s-maxage=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}

// SetCacheHeaders sets the Cache-Control headers on the response
func SetCacheHeaders(w http.ResponseWriter, cacheSeconds int) {
	if cacheSeconds < 1 || os.Getenv("NODE_ENV") == "development" {
		DisableCaching(w)
		return
	}

	cacheControl := fmt.Sprintf("max-age=%d, s-maxage=%d, stale-while-revalidate=%d", cacheSeconds, cacheSeconds, OneDay)
	w.Header().Set("Cache-Control", cacheControl)
}

// SetErrorCacheHeaders sets the Cache-Control headers for error responses
func SetErrorCacheHeaders(w http.ResponseWriter) {
	envCacheSeconds := os.Getenv("CACHE_SECONDS")
	if envCacheSeconds != "" {
		if envVal, err := strconv.Atoi(envCacheSeconds); err == nil && envVal < 1 {
			DisableCaching(w)
			return
		}
	}

	if os.Getenv("NODE_ENV") == "development" {
		DisableCaching(w)
		return
	}

	// Use lower cache period for errors
	cacheControl := fmt.Sprintf("max-age=%d, s-maxage=%d, stale-while-revalidate=%d", ErrorTTL, ErrorTTL, OneDay)
	w.Header().Set("Cache-Control", cacheControl)
}
