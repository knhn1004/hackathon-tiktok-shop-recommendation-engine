build:
	@echo "Building api-node..."
	@go build -o bin/api-node main.go

run: build
	@echo "Starting api-node..."
	@./bin/api-node

seed: build
	@echo "Seeding the database..."
	@go run cmd/seed.go

watch:
	@echo "Starting api-node in watch mode..."
	@air
