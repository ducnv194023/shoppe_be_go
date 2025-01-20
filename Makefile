# Variables
include .env

APP_NAME=shoppe_be_go
DOCKER_COMPOSE=docker-compose.yml

# Go related variables
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin
GOBUILD=$(GOBASE)/build

# Build flags
LDFLAGS=-w -s

.PHONY: all build clean test docker-build docker-up docker-down lint help

# Build binary
build:
	@echo "Building $(APP_NAME)..."
	@go build -ldflags "$(LDFLAGS)" -o $(GOBIN)/$(APP_NAME) ./cmd/server

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf $(GOBIN)/* $(GOBUILD)/*
	@go clean

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

# Run linter
lint:
	@echo "Running linter..."
	@golangci-lint run ./...

# Docker commands
docker-build:
	@echo "Building Docker images..."
	@docker-compose -f $(DOCKER_COMPOSE) build

docker-up:
	@echo "Starting Docker containers..."
	@docker-compose -f $(DOCKER_COMPOSE) up -d

docker-down:
	@echo "Stopping Docker containers..."
	@docker-compose -f $(DOCKER_COMPOSE) down

docker-logs:
	@echo "Showing Docker logs..."
	@docker-compose -f $(DOCKER_COMPOSE) logs -f

# Database commands
db-migrate:
	@echo "Running database migrations..."
	@go run ./cmd/migration

# Run application locally
run:
	@echo "Running application..."
	@go run ./cmd/server/main.go

# Help command
help:
	@echo "Available commands:"
	@echo "  make build         - Build the application"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make test          - Run tests"
	@echo "  make test-coverage - Run tests with coverage"
	@echo "  make lint          - Run linter"
	@echo "  make docker-build  - Build Docker images"
	@echo "  make docker-up     - Start Docker containers"
	@echo "  make docker-down   - Stop Docker containers"
	@echo "  make docker-logs   - Show Docker logs"
	@echo "  make db-migrate    - Run database migrations"
	@echo "  make run          - Run application locally"

# Default target
all: clean build

# Database migrations
migrate-up:
	@echo "Running migrations up..."
	@goose -v -dir internal/migrations mysql "$(DB_URL)" up

migrate-down:
	@echo "Running migrations down..."
	@goose -dir internal/migrations mysql "$(DB_URL)" down

migrate-status:
	@echo "Checking migration status..."
	@goose -dir internal/migrations mysql "$(DB_URL)" status

migrate-create:
	@read -p "Enter migration name: " name; \
	goose -dir internal/migrations create $$name sql

# SQLC commands
sqlc-gen:
	@echo "Generating SQLC code..."
	@sqlc generate