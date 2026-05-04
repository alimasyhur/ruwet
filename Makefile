.PHONY: test test-verbose test-coverage build install clean

# Run tests
test:
	go test ./...

# Run tests with verbose output
test-verbose:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"
	@rm coverage.out

# Build the binary
build:
	go build -o ruwet .

# Install to GOBIN
install:
	go install .

# Clean build artifacts
clean:
	rm -f ruwet coverage.out coverage.html

# Run linter (requires golangci-lint)
lint:
	golangci-lint run

# Run all checks
ci: test lint
