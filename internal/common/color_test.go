package common

import (
	"strings"
	"testing"
)

func TestGetCardColors_ThemeBgColor(t *testing.T) {
	tests := []struct {
		name     string
		args     map[string]string
		theme    string
		wantBg   string
		wantType string // "string" or "array"
	}{
		{
			name:     "gruvbox theme should have bg color with # prefix",
			args:     map[string]string{"theme": "gruvbox"},
			theme:    "gruvbox",
			wantBg:   "#282828",
			wantType: "string",
		},
		{
			name:     "default theme should have bg color with # prefix",
			args:     map[string]string{"theme": "default"},
			theme:    "default",
			wantBg:   "#fffefe",
			wantType: "string",
		},
		{
			name:     "dark theme should have bg color with # prefix",
			args:     map[string]string{"theme": "dark"},
			theme:    "dark",
			wantBg:   "#151515",
			wantType: "string",
		},
		{
			name:     "custom bg_color should have # prefix",
			args:     map[string]string{"theme": "gruvbox", "bg_color": "ff0000"},
			theme:    "gruvbox",
			wantBg:   "#ff0000",
			wantType: "string",
		},
		{
			name:     "custom bg_color with # prefix should remain",
			args:     map[string]string{"theme": "gruvbox", "bg_color": "#ff0000"},
			theme:    "gruvbox",
			wantBg:   "#ff0000",
			wantType: "string",
		},
		{
			name:     "custom bg_color without # prefix should add #",
			args:     map[string]string{"theme": "gruvbox", "bg_color": "ff0000"},
			theme:    "gruvbox",
			wantBg:   "#ff0000",
			wantType: "string",
		},
		{
			name:     "no theme specified should use default",
			args:     map[string]string{},
			theme:    "default",
			wantBg:   "#fffefe",
			wantType: "string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			colors := GetCardColors(tt.args)

			// Check if bgColor is a string
			bgColorStr, ok := colors.BgColor.(string)
			if !ok {
				// If it's an array (gradient), that's also valid
				if _, isArray := colors.BgColor.([]string); isArray {
					if tt.wantType == "array" {
						return // Test passed for gradient
					}
					t.Errorf("GetCardColors() BgColor = %v (array), want string %v", colors.BgColor, tt.wantBg)
					return
				}
				t.Errorf("GetCardColors() BgColor = %v (type %T), want string", colors.BgColor, colors.BgColor)
				return
			}

			// Check if it has # prefix
			if !strings.HasPrefix(bgColorStr, "#") {
				t.Errorf("GetCardColors() BgColor = %v, want string with # prefix", bgColorStr)
			}

			// Check if it matches expected value
			if bgColorStr != tt.wantBg {
				t.Errorf("GetCardColors() BgColor = %v, want %v", bgColorStr, tt.wantBg)
			}
		})
	}
}

func TestGetCardColors_AllColorsHavePrefix(t *testing.T) {
	args := map[string]string{
		"theme": "gruvbox",
	}
	colors := GetCardColors(args)

	// Check all string colors have # prefix
	if !strings.HasPrefix(colors.TitleColor, "#") {
		t.Errorf("TitleColor = %v, want string with # prefix", colors.TitleColor)
	}
	if !strings.HasPrefix(colors.IconColor, "#") {
		t.Errorf("IconColor = %v, want string with # prefix", colors.IconColor)
	}
	if !strings.HasPrefix(colors.TextColor, "#") {
		t.Errorf("TextColor = %v, want string with # prefix", colors.TextColor)
	}
	if !strings.HasPrefix(colors.BorderColor, "#") {
		t.Errorf("BorderColor = %v, want string with # prefix", colors.BorderColor)
	}
	if !strings.HasPrefix(colors.RingColor, "#") {
		t.Errorf("RingColor = %v, want string with # prefix", colors.RingColor)
	}

	// Check BgColor
	bgColorStr, ok := colors.BgColor.(string)
	if ok && !strings.HasPrefix(bgColorStr, "#") {
		t.Errorf("BgColor = %v, want string with # prefix", bgColorStr)
	}
}

func TestGetCardColors_GradientBgColor(t *testing.T) {
	args := map[string]string{
		"theme":    "gruvbox",
		"bg_color": "90,ff0000,00ff00,0000ff", // Gradient
	}
	colors := GetCardColors(args)

	// BgColor should be an array for gradients
	gradient, ok := colors.BgColor.([]string)
	if !ok {
		t.Errorf("GetCardColors() BgColor = %v (type %T), want []string for gradient", colors.BgColor, colors.BgColor)
		return
	}

	if len(gradient) < 2 {
		t.Errorf("GetCardColors() BgColor gradient = %v, want at least 2 colors", gradient)
	}
}

func TestFallbackColor_BgColorPrefix(t *testing.T) {
	tests := []struct {
		name        string
		color       string
		fallback    interface{}
		wantHasHash bool
		wantType    string
	}{
		{
			name:        "empty color with string fallback should add #",
			color:       "",
			fallback:    "282828",
			wantHasHash: true,
			wantType:    "string",
		},
		{
			name:        "empty color with # fallback should keep #",
			color:       "",
			fallback:    "#282828",
			wantHasHash: true,
			wantType:    "string",
		},
		{
			name:        "valid hex color should add #",
			color:       "ff0000",
			fallback:    "282828",
			wantHasHash: true,
			wantType:    "string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FallbackColor(tt.color, tt.fallback)

			if tt.wantType == "string" {
				resultStr, ok := result.(string)
				if !ok {
					t.Errorf("FallbackColor() = %v (type %T), want string", result, result)
					return
				}

				if tt.wantHasHash && !strings.HasPrefix(resultStr, "#") {
					t.Errorf("FallbackColor() = %v, want string with # prefix", resultStr)
				}
			}
		})
	}
}
