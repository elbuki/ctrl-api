# Performs the execution of unit tests and coverage analysis
coverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

# Run the linter service with the current code
lint:
	@golangci-lint run

# Prepares the local repository to have a hooks executing
prepare:
	git config core.hooksPath hooks/
	# Install linter
	GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.26.0

# Builds and brings up the project
start:
	./build/src

# Run the tests to ensure everything's alright
test:
	@go test ./...
