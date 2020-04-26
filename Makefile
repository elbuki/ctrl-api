include .env

# Prepares the local repository to have a hooks executing
prepare:
	git config core.hooksPath hooks/

# Run the linter service with the current code
lint:
	@docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.24.0-alpine golangci-lint run

# Builds and brings up the project
start:
	docker-compose build
	docker-compose up
