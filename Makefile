# Variables
APP_NAME = h4x0r_b41t
GO_FILES = $(shell find . -type f -name '*.go')

# Default target: Build the application
all: build

# Build the Go application
build:
	@echo "Building $(APP_NAME)..."
	GOOS=linux GOARCH=amd64 go build -o $(APP_NAME)

# Run the application
run: build
	@echo "Running $(APP_NAME)..."
	./$(APP_NAME)

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)

# Format the Go source code
fmt:
	@echo "Formatting Go source files..."
	go fmt ./...

# Run Go tests
test:
	@echo "Running tests..."
	go test ./...

# Install project dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy

# Build and run with Docker Compose
docker:
	@echo "Building and running with Docker Compose..."
	docker-compose up --build

# Stop and remove Docker containers
docker-clean:
	@echo "Stopping and cleaning up Docker containers..."
	docker-compose down

# Help message
help:
	@echo "Makefile for $(APP_NAME)"
	@echo "Usage:"
	@echo "  make             Build the application"
	@echo "  make build       Build the Go application"
	@echo "  make run         Run the application"
	@echo "  make clean       Clean up build artifacts"
	@echo "  make fmt         Format the Go source code"
	@echo "  make test        Run Go tests"
	@echo "  make deps        Install project dependencies"
	@echo "  make docker      Build and run with Docker Compose"
	@echo "  make docker-clean Stop and remove Docker containers"

# Phony targets
.PHONY: all build run clean fmt test deps docker docker-clean help
