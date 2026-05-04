.PHONY: test test-verbose test-coverage build install clean

# Run tests
test:
	go test ./...

# Run tests with verbose output
test-verbose:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -cover ./...
	@echo "Coverage checked"

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
