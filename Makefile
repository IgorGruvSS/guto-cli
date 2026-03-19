# Guto CLI Makefile

BINARY_NAME=guto
INSTALL_PATH=/usr/local/bin

.PHONY: all build install uninstall clean help dev test setup deps whisper post-install

all: build

build: ## Build the binary locally
	@echo "🔨 Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) main.go

install: build ## Install the binary to $(INSTALL_PATH)
	@echo "🚀 Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
	sudo mv $(BINARY_NAME) $(INSTALL_PATH)/

uninstall: ## Remove the binary from $(INSTALL_PATH)
	@echo "🗑️ Uninstalling $(BINARY_NAME) from $(INSTALL_PATH)..."
	sudo rm -f $(INSTALL_PATH)/$(BINARY_NAME)

deps: ## Install system dependencies (requires sudo)
	@bash scripts/setup_deps.sh

whisper: ## Configure Whisper environment (requires sudo)
	@sudo bash scripts/setup_whisper.sh

post-install: ## Show post-installation instructions
	@bash scripts/post_install.sh

setup: deps whisper install post-install ## Full system setup (dependencies + whisper + binary)

tools: ## Install development tools (golangci-lint, govulncheck)
	@echo "🛠️ Installing development tools..."
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.5

fmt: ## Format code
	@echo "🧹 Formatting code..."
	go fmt ./...

lint: ## Run linter
	@echo "🔍 Running linter..."
	golangci-lint run ./...

vuln: ## Check for vulnerabilities
	@echo "🛡️ Checking for vulnerabilities..."
	-govulncheck ./...

check-coverage: ## Run tests and enforce 75% coverage threshold (excludes internal/)
	@echo "🧪 Running tests with coverage..."
	@go test -v -coverprofile=coverage.out ./cmd/...
	@go tool cover -func=coverage.out
	@head -1 coverage.out > coverage_filtered.out
	@grep -v "^mode:" coverage.out | grep -v "/internal/" >> coverage_filtered.out
	@COVERAGE=$$(go tool cover -func=coverage_filtered.out | grep total | awk '{print $$3}' | sed 's/%//'); \
	echo "Coverage (excluding internal/): $$COVERAGE%"; \
	if [ $$(echo "$$COVERAGE < 75" | bc -l) -eq 1 ]; then \
		echo "❌ Coverage is below 75%"; exit 1; \
	fi

ci: fmt lint vuln check-coverage ## Run all CI checks locally

hooks: ## Install git hooks
	@bash scripts/setup_hooks.sh

test: ## Run all tests
	@echo "🧪 Running tests..."
	go test -v ./...

coverage: ## Run tests with coverage
	@echo "🧪 Running tests with coverage..."
	-go test -v -coverprofile=coverage.out ./cmd/...
	-go tool cover -func=coverage.out

clean: ## Remove the local binary and Output folder
	@echo "🧹 Cleaning up..."
	rm -f $(BINARY_NAME)
	rm -rf Output/
	rm -f coverage.out

dev: ## Run the project directly for development
	go run main.go

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
