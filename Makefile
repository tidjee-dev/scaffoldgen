.PHONY: build test clean version bump-patch bump-minor bump-major

# Default target
all: build

# Build the application
build:
	go build -ldflags="-X github.com/tidjee-dev/scaffoldgen/internal/app.Version=$(shell cat version.json | jq -r .version) -X github.com/tidjee-dev/scaffoldgen/internal/app.BuildDate=$(shell cat version.json | jq -r .build.date) -X github.com/tidjee-dev/scaffoldgen/internal/app.BuildCommit=$(shell cat version.json | jq -r .build.commit) -X github.com/tidjee-dev/scaffoldgen/internal/app.GoVersion=$(shell cat version.json | jq -r .metadata.go_version)" -o scaffoldgen ./cmd/scaffoldgen

# Run tests
test:
	go test ./...

# Clean build artifacts
clean:
	rm -f scaffoldgen

# Update version (requires version argument)
version:
	@if [ -z "$(VERSION)" ]; then \
		echo "Usage: make version VERSION=1.0.1"; \
		exit 1; \
	fi
	go run scripts/update-version.go $(VERSION)

# Bump patch version (1.0.0 -> 1.0.1)
bump-patch:
	@./scripts/bump-version.sh patch

# Bump minor version (1.0.0 -> 1.1.0)
bump-minor:
	@./scripts/bump-version.sh minor

# Bump major version (1.0.0 -> 2.0.0)
bump-major:
	@./scripts/bump-version.sh major

# Install dependencies
deps:
	go mod download
	go mod tidy

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	go vet ./...

# Run all checks
check: fmt lint test

# Build with version info
build-with-version: build
	@echo "Built with version: $$(./scaffoldgen --version)"
