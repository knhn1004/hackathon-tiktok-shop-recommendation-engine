# Database connection details
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_USER ?= postgres
DB_PASSWORD ?= postgres
DB_NAME ?= postgres

.PHONY: build run watch migrate generate-migration

db-gen:
	@echo "Generating migration from models..."
	@go run cmd/migration_generator/main.go

migrate:
	@echo "Running database migrations..."
	@go run cmd/migrate/main.go

build:
	@echo "Building api-node..."
	@go build -o bin/api-node main.go

run: migrate build
	@echo "Starting api-node..."
	@./bin/api-node

watch: migrate
	@echo "Starting api-node in watch mode..."
	@air
