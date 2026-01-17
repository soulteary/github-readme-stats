package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/soulteary/github-readme-stats/internal/api"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		// .env file is optional
		log.Println("No .env file found, using environment variables")
	}

	// Check if running in CLI mode
	if len(os.Args) > 1 {
		if err := runCLI(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	// Set Gin mode
	if os.Getenv("NODE_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router
	r := gin.Default()

	// API routes
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/", api.StatsHandler)
		apiGroup.GET("/pin", api.PinHandler)
		apiGroup.GET("/top-langs", api.TopLangsHandler)
		apiGroup.GET("/gist", api.GistHandler)
		apiGroup.GET("/wakatime", api.WakaTimeHandler)
	}

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	// Start server
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	log.Printf("Server running on port %s", port)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}

func runCLI() error {
	var cardType, outputPath string
	var params map[string]string

	// Check for help flag
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help") {
		printUsage()
		return nil
	}

	// Check if first argument is a URL
	if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "/api") {
		// URL format
		urlStr := os.Args[1]
		parsedURL, err := url.Parse(urlStr)
		if err != nil {
			return fmt.Errorf("invalid URL: %v", err)
		}

		// Determine card type from path
		path := parsedURL.Path
		switch path {
		case "/api", "/api/":
			cardType = "stats"
		case "/api/pin":
			cardType = "pin"
		case "/api/top-langs":
			cardType = "top-langs"
		case "/api/gist":
			cardType = "gist"
		case "/api/wakatime":
			cardType = "wakatime"
		default:
			return fmt.Errorf("unknown API path: %s", path)
		}

		// Parse query parameters
		params = make(map[string]string)
		for key, values := range parsedURL.Query() {
			if len(values) > 0 {
				params[key] = values[0]
			}
		}

		// Check for --output flag in remaining args
		flagSet := flag.NewFlagSet("", flag.ContinueOnError)
		flagSet.StringVar(&outputPath, "output", "", "Output file path (default: stdout)")
		if len(os.Args) > 2 {
			flagSet.Parse(os.Args[2:])
		}
	} else {
		// Flag format - parse all arguments manually for flexibility
		args := os.Args[1:]
		params = make(map[string]string)

		// Parse arguments
		for i := 0; i < len(args); i++ {
			arg := args[i]
			if strings.HasPrefix(arg, "--") {
				if strings.Contains(arg, "=") {
					// --key=value format
					parts := strings.SplitN(arg[2:], "=", 2)
					if len(parts) == 2 {
						key := parts[0]
						value := parts[1]
						if key == "type" {
							cardType = value
						} else if key == "output" {
							outputPath = value
						} else {
							params[key] = value
						}
					}
				} else {
					// --key value format
					key := arg[2:]
					if key == "type" {
						if i+1 < len(args) {
							cardType = args[i+1]
							i++
						}
					} else if key == "output" {
						if i+1 < len(args) {
							outputPath = args[i+1]
							i++
						}
					} else {
						if i+1 < len(args) && !strings.HasPrefix(args[i+1], "--") {
							params[key] = args[i+1]
							i++
						} else {
							// Boolean flag without value
							params[key] = "true"
						}
					}
				}
			} else if !strings.HasPrefix(arg, "-") {
				// Positional argument - treat as error or ignore
				return fmt.Errorf("unexpected positional argument: %s (use --key=value format)", arg)
			}
		}

		if cardType == "" {
			return fmt.Errorf("--type is required (options: stats, pin, top-langs, gist, wakatime)")
		}
	}

	// Generate card based on type
	switch cardType {
	case "stats":
		return api.GenerateStatsCard(params, outputPath)
	case "pin":
		return api.GeneratePinCard(params, outputPath)
	case "top-langs":
		return api.GenerateTopLangsCard(params, outputPath)
	case "gist":
		return api.GenerateGistCard(params, outputPath)
	case "wakatime":
		return api.GenerateWakaTimeCard(params, outputPath)
	default:
		return fmt.Errorf("unknown card type: %s (options: stats, pin, top-langs, gist, wakatime)", cardType)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `Usage:
  %s [OPTIONS]                    # Start HTTP server
  %s "/api?username=xxx&..."     # Generate card from URL
  %s --type=TYPE [OPTIONS]        # Generate card with flags

URL Format:
  /api?username=xxx&theme=dark
  /api/pin?username=xxx&repo=xxx
  /api/top-langs?username=xxx
  /api/gist?id=xxx
  /api/wakatime?username=xxx

Flag Format:
  --type=TYPE                     # Card type (required): stats, pin, top-langs, gist, wakatime
  --output=FILE                   # Output file path (default: stdout)
  --username=USER                 # GitHub username
  --theme=THEME                   # Theme name
  [other parameters...]           # Any other API parameters

Examples:
  %s "/api?username=anuraghazra&hide=contribs,prs" --output=stats.svg
  %s --type=stats --username=soulteary --theme=dark --output=stats.svg
  %s --type=pin --username=soulteary --repo=github-readme-stats
  %s "/api/top-langs?username=soulteary&theme=dark"

`, os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0], os.Args[0])
}
