# Database connection details
DB_HOST ?= localhost
DB_PORT ?= 5432
DB_USER ?= postgres
DB_PASSWORD ?= postgres
DB_NAME ?= postgres

.PHONY: build run watch

build:
	@echo "Building api-node..."
	@go build -o bin/api-node main.go

run: build
	@echo "Starting api-node..."
	@./bin/api-node

watch:
	@echo "Starting api-node in watch mode..."
	@air
