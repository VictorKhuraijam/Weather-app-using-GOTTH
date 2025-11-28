.PHONY: install templ dev run build clean docker-build docker-run test fmt lint help

# Install dependencies and tools
install:
	@echo "📦 Installing dependencies..."
	go mod download
	go install github.com/a-h/templ/cmd/templ@latest
	@echo "✅ Dependencies installed"

# Generate templ files
templ:
	@echo "🔨 Generating templ files..."
	templ generate
	@echo "✅ Templ files generated"

# Development mode with hot reload
dev:
	@echo "🚀 Starting development mode with hot reload..."
	templ generate --watch --proxy="http://localhost:8080" --cmd="go run ./main.go"

# Run the application
run: templ
	@echo "🚀 Running application..."
	go run ./main.go

# Build the application
build: templ
	@echo "🔨 Building application..."
	go build -o bin/weather-app ./main.go
	@echo "✅ Binary created at bin/weather-app"

# Clean generated files and build artifacts
clean:
	@echo "🧹 Cleaning..."
	rm -rf bin/
	find . -name "*_templ.go" -type f -delete
	@echo "✅ Cleaned"

# Docker build
docker-build:
	@echo "🐳 Building Docker image..."
	docker build -t weather-app-gotth .
	@echo "✅ Docker image built"

# Docker run
docker-run:
	@echo "🐳 Running Docker container..."
	docker run -p 8080:8080 --env-file .env weather-app-gotth

# Run tests
test:
	@echo "🧪 Running tests..."
	go test -v ./...

# Format code
fmt:
	@echo "✨ Formatting code..."
	go fmt ./...
	templ fmt .
	@echo "✅ Code formatted"

# Lint code
lint:
	@echo "🔍 Linting code..."
	golangci-lint run
	@echo "✅ Linting complete"

# Show help
help:
	@echo "Available commands:"
	@echo "  make install      - Install dependencies and tools"
	@echo "  make templ        - Generate templ files"
	@echo "  make dev          - Run with hot reload (recommended for development)"
	@echo "  make run          - Run the application"
	@echo "  make build        - Build the application binary"
	@echo "  make clean        - Clean generated files and build artifacts"
	@echo "  make docker-build - Build Docker image"
	@echo "  make docker-run   - Run Docker container"
	@echo "  make test         - Run tests"
	@echo "  make fmt          - Format code"
	@echo "  make lint         - Lint code"
	@echo "  make help         - Show this help message"
