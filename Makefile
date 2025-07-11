# Makefile for Go project formatting and linting

# Default target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  fmt FILE=<file>     - Format a specific Go file"
	@echo "  lint FILE=<file>    - Lint a specific Go file"
	@echo "  check FILE=<file>   - Run both fmt and lint on a file"
	@echo "  check-all           - Run fmt and lint on all Go files"
	@echo ""
	@echo "Usage examples:"
	@echo "  make fmt FILE=main.go"
	@echo "  make lint FILE=main.go"
	@echo "  make check FILE=main.go"
	@echo "  make check-all"

# Format a specific file
.PHONY: fmt
fmt:
	@if [ -z "$(FILE)" ]; then \
		echo "Error: FILE parameter is required. Usage: make fmt FILE=<filename>"; \
		exit 1; \
	fi
	@echo "Formatting $(FILE)..."
	go fmt $(FILE)

# Lint a specific file
.PHONY: lint
lint:
	@if [ -z "$(FILE)" ]; then \
		echo "Error: FILE parameter is required. Usage: make lint FILE=<filename>"; \
		exit 1; \
	fi
	@echo "Linting $(FILE)..."
	golangci-lint run $(FILE)

# Run both fmt and lint on a file
.PHONY: check
check: fmt lint
	@echo "Formatting and linting complete for $(FILE)"

# Run fmt and lint on all Go files
.PHONY: check-all
check-all:
	@echo "Formatting all Go files..."
	go fmt ./...
	@echo "Linting all Go files..."
	golangci-lint run ./...
	@echo "Formatting and linting complete for all files"