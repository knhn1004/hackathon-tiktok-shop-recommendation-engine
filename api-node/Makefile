build:
	@go build -o bin/api-node main.go

run: build
	@cp .env.local bin/.env.local && ./bin/api-node
watch:
	@air