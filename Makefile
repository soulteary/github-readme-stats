# Makefile for github-readme-stats

.PHONY: help build build-server build-examples clean test examples run-server

# Default target
help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build all binaries
	@echo "Building all binaries..."
	go build -o bin/server ./cmd/server
	go build -o bin/examples ./cmd/examples

build-server: ## Build the server binary
	@echo "Building server..."
	go build -o bin/server ./cmd/server

build-examples: ## Build the examples generator binary
	@echo "Building examples generator..."
	go build -o bin/examples ./cmd/examples

clean: ## Clean build artifacts
	@echo "Cleaning..."
	rm -rf bin/
	rm -f .github/assets/*-test.svg .github/assets/stats-*.svg .github/assets/repo-*.svg .github/assets/top-langs-*.svg .github/assets/gist-*.svg .github/assets/wakatime-*.svg

test: ## Run tests
	@echo "Running tests..."
	go test ./...

examples: ## Generate example images (requires network access)
	@echo "Generating example images..."
	./bin/examples

run-server: ## Run the server (requires network access)
	@echo "Starting server..."
	./bin/server

# Create bin directory if it doesn't exist
bin:
	mkdir -p bin

# Ensure bin directory exists before building
build-server: bin
build-examples: bin