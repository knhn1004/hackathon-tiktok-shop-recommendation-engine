build:
	@echo "Building api-node..."
	@go build -o bin/api-node main.go

run: build add-vector-extension
	@echo "Starting api-node..."
	@./bin/api-node

seed: build
	@echo "Seeding the database..."
	@go run -tags seed cmd/seed/main.go

purge-db:
	@echo "Purging the database..."
	@go run -tags purge_db cmd/purge_db/main.go

add-vector-extension:
	@echo "Adding vector extension to the database..."
	@PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -U $(DB_USER) -d $(DB_NAME) -c "CREATE EXTENSION IF NOT EXISTS vector;"

watch: add-vector-extension
	@echo "Starting api-node in watch mode..."
	@air
