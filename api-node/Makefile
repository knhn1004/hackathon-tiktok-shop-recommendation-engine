build:
	@echo "Building api-node..."
	@go build -o bin/api-node main.go

run: build
	@echo "Starting api-node..."
	@./bin/api-node

seed: build
	@echo "Seeding the database..."
	@go run -tags seed cmd/seed/main.go

purge-db:
	@echo "Purging the database..."
	@go run -tags purge_db cmd/purge_db/main.go

watch:
	@echo "Starting api-node in watch mode..."
	@air
