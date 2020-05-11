# Prepares the local repository to have a hooks executing
prepare:
	git config core.hooksPath hooks/

	# Install linters
	go get github.com/walle/lll/...
	go get github.com/client9/misspell/cmd/misspell

# Run the linter service with the current code
lint:
	@golangci-lint run

# Performs the execution of unit tests and coverage analysis
tests:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

# Builds and brings up the project
start:
	docker-compose build
	docker-compose up
