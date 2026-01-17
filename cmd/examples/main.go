package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/soulteary/github-readme-stats/internal/api"
)

type ExampleConfig struct {
	Name        string
	Description string
	Params      map[string]string
	OutputPath  string
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		// .env file is optional
		log.Println("No .env file found, using environment variables")
	}

	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help") {
		printUsage()
		return
	}

	// Check for test mode (no network required)
	if len(os.Args) > 1 && os.Args[1] == "--test" {
		runTestMode()
		return
	}

	// Check for wakatime test mode (use test data)
	if len(os.Args) > 1 && os.Args[1] == "--wakatime-test" {
		runWakaTimeTestMode()
		return
	}

	// 创建示例输出目录
	outputDir := ".github/assets"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// 定义各种示例配置
	examples := []ExampleConfig{
		// Stats Card Examples
		{
			Name:        "stats-basic",
			Description: "Basic stats card with default theme",
			Params: map[string]string{
				"username": "anuraghazra",
				"theme":    "default",
			},
			OutputPath: filepath.Join(outputDir, "stats-basic.svg"),
		},
		{
			Name:        "stats-dark",
			Description: "Stats card with dark theme",
			Params: map[string]string{
				"username": "anuraghazra",
				"theme":    "dark",
			},
			OutputPath: filepath.Join(outputDir, "stats-dark.svg"),
		},
		{
			Name:        "stats-compact",
			Description: "Compact stats card with custom styling",
			Params: map[string]string{
				"username":   "anuraghazra",
				"hide":       "contribs,prs",
				"card_width": "320",
				"theme":      "radical",
			},
			OutputPath: filepath.Join(outputDir, "stats-compact.svg"),
		},
		{
			Name:        "stats-custom",
			Description: "Stats card with custom colors and settings",
			Params: map[string]string{
				"username":     "anuraghazra",
				"show_icons":   "true",
				"title_color":  "ff4757",
				"text_color":   "586069",
				"bg_color":     "282a36",
				"border_color": "ff6b6b",
				"theme":        "merko",
			},
			OutputPath: filepath.Join(outputDir, "stats-custom.svg"),
		},

		// Repo Pin Examples
		{
			Name:        "repo-basic",
			Description: "Basic repository pin card",
			Params: map[string]string{
				"username": "anuraghazra",
				"repo":     "github-readme-stats",
			},
			OutputPath: filepath.Join(outputDir, "repo-basic.svg"),
		},
		{
			Name:        "repo-themed",
			Description: "Repository pin card with custom theme",
			Params: map[string]string{
				"username": "anuraghazra",
				"repo":     "github-readme-stats",
				"theme":    "dark",
			},
			OutputPath: filepath.Join(outputDir, "repo-themed.svg"),
		},

		// Top Languages Examples
		{
			Name:        "top-langs-basic",
			Description: "Basic top languages card",
			Params: map[string]string{
				"username": "anuraghazra",
			},
			OutputPath: filepath.Join(outputDir, "top-langs-basic.svg"),
		},
		{
			Name:        "top-langs-compact",
			Description: "Compact top languages card",
			Params: map[string]string{
				"username":    "anuraghazra",
				"layout":      "compact",
				"langs_count": "8",
				"theme":       "dark",
			},
			OutputPath: filepath.Join(outputDir, "top-langs-compact.svg"),
		},

		// Gist Examples
		{
			Name:        "gist-basic",
			Description: "Basic gist pin card",
			Params: map[string]string{
				"id": "bbfce31e0217a3689c8d961a356cb10d",
			},
			OutputPath: filepath.Join(outputDir, "gist-basic.svg"),
		},

		// WakaTime Examples
		{
			Name:        "wakatime-basic",
			Description: "Basic WakaTime stats card",
			Params: map[string]string{
				"username": "anuraghazra",
			},
			OutputPath: filepath.Join(outputDir, "wakatime-basic.svg"),
		},
		{
			Name:        "wakatime-compact",
			Description: "Compact WakaTime stats card",
			Params: map[string]string{
				"username":    "anuraghazra",
				"layout":      "compact",
				"langs_count": "6",
				"theme":       "dark",
			},
			OutputPath: filepath.Join(outputDir, "wakatime-compact.svg"),
		},
	}

	// 检查是否指定了特定的示例
	var selectedExamples []ExampleConfig
	if len(os.Args) > 1 {
		requestedNames := os.Args[1:]
		for _, name := range requestedNames {
			found := false
			for _, example := range examples {
				if example.Name == name {
					selectedExamples = append(selectedExamples, example)
					found = true
					break
				}
			}
			if !found {
				fmt.Fprintf(os.Stderr, "Error: Example '%s' not found. Available examples:\n", name)
				printAvailableExamples(examples)
				os.Exit(1)
			}
		}
	} else {
		// 生成所有示例
		selectedExamples = examples
	}

	fmt.Printf("Generating %d example(s)...\n", len(selectedExamples))

	generated := 0
	failed := 0

	for _, example := range selectedExamples {
		fmt.Printf("  Generating %s: %s\n", example.Name, example.Description)

		var err error
		switch {
		case example.Name[:5] == "stats":
			err = api.GenerateStatsCard(example.Params, example.OutputPath)
		case example.Name[:4] == "repo":
			err = api.GeneratePinCard(example.Params, example.OutputPath)
		case example.Name[:9] == "top-langs":
			err = api.GenerateTopLangsCard(example.Params, example.OutputPath)
		case example.Name[:4] == "gist":
			err = api.GenerateGistCard(example.Params, example.OutputPath)
		case example.Name[:8] == "wakatime":
			err = api.GenerateWakaTimeCard(example.Params, example.OutputPath)
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "    Error generating %s: %v\n", example.Name, err)
			failed++
		} else {
			fmt.Printf("    ✓ Generated %s\n", example.OutputPath)
			generated++
		}
	}

	fmt.Printf("\nCompleted! Generated: %d, Failed: %d\n", generated, failed)
	if generated > 0 {
		fmt.Printf("Example images saved to: %s/\n", outputDir)
	}
}

func runTestMode() {
	// Load environment variables for test mode as well
	if err := godotenv.Load(); err != nil {
		// .env file is optional
		log.Println("No .env file found, using environment variables")
	}

	fmt.Println("Test Mode: Generating sample error cards (no network required)")

	outputDir := ".github/assets"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Generate test error cards for each type
	testExamples := []struct {
		name string
		fn   func(map[string]string, string) error
	}{
		{"stats-test", api.GenerateStatsCard},
		{"repo-test", api.GeneratePinCard},
		{"top-langs-test", api.GenerateTopLangsCard},
		{"gist-test", api.GenerateGistCard},
		{"wakatime-test", api.GenerateWakaTimeCard},
	}

	for _, test := range testExamples {
		outputPath := filepath.Join(outputDir, test.name+".svg")
		// Use empty params to trigger error (no username/token)
		err := test.fn(map[string]string{}, outputPath)
		// Check if file was created (even with error, it should create an error SVG)
		if _, statErr := os.Stat(outputPath); statErr == nil {
			fmt.Printf("✓ Generated test error card: %s\n", outputPath)
		} else {
			fmt.Printf("✗ Failed to generate test card: %s (error: %v)\n", test.name, err)
		}
	}

	fmt.Println("\nTest completed! Check .github/assets/ directory for error cards.")
	fmt.Println("These demonstrate error handling when API tokens or parameters are missing.")
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `Generate Example Images

Usage:
  %s [EXAMPLE_NAME...]    # Generate specific examples
  %s                      # Generate all examples
  %s --test               # Generate test error cards (no network)
  %s --wakatime-test      # Generate WakaTime test cards using test data (no network)

Description:
  This tool generates example images for all supported card types,
  demonstrating various themes and configurations.

Options:
  EXAMPLE_NAME       Name of specific example to generate (see below)
  --test             Generate error cards for testing (no network access required)
  --wakatime-test    Generate WakaTime test cards using test data (no network access required)

Available Examples:

`, os.Args[0], os.Args[0], os.Args[0])

	printAvailableExamples(getAllExamples())
	fmt.Fprintf(os.Stderr, `
Output:
  All images will be saved as SVG files in the '.github/assets/' directory.

Examples:
  %s stats-basic stats-dark    # Generate only basic and dark stats
  %s                           # Generate all examples
  %s --test                    # Generate test error cards
  %s --wakatime-test           # Generate WakaTime test cards

`, os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}

func printAvailableExamples(examples []ExampleConfig) {
	fmt.Fprintf(os.Stderr, "Stats Cards:\n")
	for _, ex := range examples {
		if ex.Name[:5] == "stats" {
			fmt.Fprintf(os.Stderr, "  %-18s %s\n", ex.Name, ex.Description)
		}
	}

	fmt.Fprintf(os.Stderr, "\nRepository Pins:\n")
	for _, ex := range examples {
		if ex.Name[:4] == "repo" {
			fmt.Fprintf(os.Stderr, "  %-18s %s\n", ex.Name, ex.Description)
		}
	}

	fmt.Fprintf(os.Stderr, "\nTop Languages:\n")
	for _, ex := range examples {
		if ex.Name[:9] == "top-langs" {
			fmt.Fprintf(os.Stderr, "  %-18s %s\n", ex.Name, ex.Description)
		}
	}

	fmt.Fprintf(os.Stderr, "\nGist Pins:\n")
	for _, ex := range examples {
		if ex.Name[:4] == "gist" {
			fmt.Fprintf(os.Stderr, "  %-18s %s\n", ex.Name, ex.Description)
		}
	}

	fmt.Fprintf(os.Stderr, "\nWakaTime Stats:\n")
	for _, ex := range examples {
		if ex.Name[:8] == "wakatime" {
			fmt.Fprintf(os.Stderr, "  %-18s %s\n", ex.Name, ex.Description)
		}
	}
}

func runWakaTimeTestMode() {
	// Load environment variables for test mode as well
	if err := godotenv.Load(); err != nil {
		// .env file is optional
		log.Println("No .env file found, using environment variables")
	}

	fmt.Println("WakaTime Test Mode: Generating sample cards using test data (no network required)")

	outputDir := ".github/assets"
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	// Find test data file
	testDataPath := ""
	wd, err := os.Getwd()
	if err == nil {
		possiblePaths := []string{
			filepath.Join(wd, "testdata", "wakatime.example.json"),
			filepath.Join(wd, "..", "testdata", "wakatime.example.json"),
			filepath.Join(wd, "github-readme-stats", "testdata", "wakatime.example.json"),
		}
		for _, path := range possiblePaths {
			if _, err := os.Stat(path); err == nil {
				testDataPath = path
				break
			}
		}
	}

	if testDataPath == "" {
		fmt.Fprintf(os.Stderr, "Error: test data file not found. Please ensure testdata/wakatime.example.json exists.\n")
		os.Exit(1)
	}

	fmt.Printf("Using test data from: %s\n", testDataPath)

	// Define WakaTime test examples
	testExamples := []ExampleConfig{
		{
			Name:        "wakatime-test-basic",
			Description: "Basic WakaTime stats card (test data)",
			Params: map[string]string{
				"test_data_path": testDataPath,
			},
			OutputPath: filepath.Join(outputDir, "wakatime-test-basic.svg"),
		},
		{
			Name:        "wakatime-test-compact",
			Description: "Compact WakaTime stats card (test data)",
			Params: map[string]string{
				"test_data_path": testDataPath,
				"layout":         "compact",
				"langs_count":    "6",
				"theme":          "dark",
			},
			OutputPath: filepath.Join(outputDir, "wakatime-test-compact.svg"),
		},
		{
			Name:        "wakatime-test-themed",
			Description: "WakaTime stats card with custom theme (test data)",
			Params: map[string]string{
				"test_data_path": testDataPath,
				"theme":          "radical",
				"langs_count":    "5",
			},
			OutputPath: filepath.Join(outputDir, "wakatime-test-themed.svg"),
		},
		{
			Name:        "wakatime-test-hide-progress",
			Description: "WakaTime stats card without progress bars (test data)",
			Params: map[string]string{
				"test_data_path": testDataPath,
				"hide_progress":  "true",
				"theme":          "default",
			},
			OutputPath: filepath.Join(outputDir, "wakatime-test-hide-progress.svg"),
		},
		{
			Name:        "wakatime-test-percent",
			Description: "WakaTime stats card with percent display (test data)",
			Params: map[string]string{
				"test_data_path": testDataPath,
				"display_format": "percent",
				"theme":          "default",
			},
			OutputPath: filepath.Join(outputDir, "wakatime-test-percent.svg"),
		},
		{
			Name:        "wakatime-test-limited",
			Description: "WakaTime stats card with limited languages (test data)",
			Params: map[string]string{
				"test_data_path": testDataPath,
				"langs_count":    "3",
				"theme":          "tokyonight",
			},
			OutputPath: filepath.Join(outputDir, "wakatime-test-limited.svg"),
		},
	}

	fmt.Printf("Generating %d WakaTime test example(s)...\n", len(testExamples))

	generated := 0
	failed := 0

	for _, example := range testExamples {
		fmt.Printf("  Generating %s: %s\n", example.Name, example.Description)

		err := api.GenerateWakaTimeCard(example.Params, example.OutputPath)

		if err != nil {
			fmt.Fprintf(os.Stderr, "    Error generating %s: %v\n", example.Name, err)
			failed++
		} else {
			fmt.Printf("    ✓ Generated %s\n", example.OutputPath)
			generated++
		}
	}

	fmt.Printf("\nCompleted! Generated: %d, Failed: %d\n", generated, failed)
	if generated > 0 {
		fmt.Printf("Test images saved to: %s/\n", outputDir)
	}
}

func getAllExamples() []ExampleConfig {
	return []ExampleConfig{
		{"stats-basic", "Basic stats card with default theme", nil, ""},
		{"stats-dark", "Stats card with dark theme", nil, ""},
		{"stats-compact", "Compact stats card with custom styling", nil, ""},
		{"stats-custom", "Stats card with custom colors and settings", nil, ""},
		{"repo-basic", "Basic repository pin card", nil, ""},
		{"repo-themed", "Repository pin card with custom theme", nil, ""},
		{"top-langs-basic", "Basic top languages card", nil, ""},
		{"top-langs-compact", "Compact top languages card", nil, ""},
		{"gist-basic", "Basic gist pin card", nil, ""},
		{"wakatime-basic", "Basic WakaTime stats card", nil, ""},
		{"wakatime-compact", "Compact WakaTime stats card", nil, ""},
	}
}
