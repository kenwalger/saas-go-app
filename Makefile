.PHONY: build run test clean deps migrate

# Build the application
build:
	go build -o bin/server ./cmd/server

# Run the application
run:
	go run ./cmd/server

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Install dependencies
deps:
	go mod download
	go mod tidy

# Run database migrations (creates tables)
migrate:
	go run ./cmd/server -migrate || echo "Note: Migration is handled automatically on startup"

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run ./... || echo "Note: Install golangci-lint for linting"

# Start frontend dev server (requires npm/yarn)
frontend-dev:
	cd web/frontend && npm run dev

# Build frontend
frontend-build:
	cd web/frontend && npm run build

