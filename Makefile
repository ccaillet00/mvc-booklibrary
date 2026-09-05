# ==============================================================================
# Makefile for mvc-booklibrary
# Usage: make <target>  (run `make` alone to see all targets)
# ==============================================================================

BINARY  := bin/booklibrary
MODE    ?= cli

.PHONY: help build test test-verbose vet run run-http run-both docker-up docker-down docker-build clean

help: ## Show all available targets
	@grep -E '^[a-zA-Z0-9_-]+:.*?## ' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-14s\033[0m %s\n", $$1, $$2}'
	@echo ""
	@echo "  Variables: MODE=cli|http|both (default: cli)"

build: ## Compile the binary into bin/
	@mkdir -p bin
	go build -o $(BINARY) .
	@echo "Built $(BINARY)"

test: ## Run all tests
	go test ./...

test-verbose: ## Run all tests with verbose output, race detector and coverage
	go test ./... -v -race -cover

vet: ## Run go vet for static analysis
	go vet ./...

run-http: ## Run the app as HTTP server on :8080
	go run . -mode http

run-gin-http: ## Run the app as an HTTP server using Gin on :8080
	go run . -mode gin-http

run-terminal: ## Run the app in CLI mode
	go run . -mode cli

docker-build: ## Build the Docker image
	docker compose build

docker-up: ## Build and start the Docker container
	docker compose up --build

docker-down: ## Stop and remove the Docker container
	docker compose down

clean: ## Remove build artifacts and local data
	rm -rf bin/ books.json
	@echo "Cleaned."
