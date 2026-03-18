# Guto CLI Makefile

BINARY_NAME=guto
INSTALL_PATH=/usr/local/bin

.PHONY: all build install clean help dev test

all: build

build: ## Build the binary locally
	@echo "🔨 Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) main.go

install: build ## Install the binary to $(INSTALL_PATH)
	@echo "🚀 Installing $(BINARY_NAME) to $(INSTALL_PATH)..."
	sudo mv $(BINARY_NAME) $(INSTALL_PATH)/

test: ## Run tests with coverage
	@echo "🧪 Running tests..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

clean: ## Remove the local binary and Output folder
	@echo "🧹 Cleaning up..."
	rm -f $(BINARY_NAME)
	rm -rf Output/
	rm -f coverage.out

dev: ## Run the project directly for development
	go run main.go

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
