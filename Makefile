# Makefile for building secwrapper on Windows

# Define the Go binary name and platform-specific variables
BINARY_NAME=checkexec
GOOS=windows
GOARCH=amd64
BUILD_DIR=build

# Build command
build:
	@echo "Building $(BINARY_NAME) for $(GOOS)/$(GOARCH)..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BUILD_DIR)/$(BINARY_NAME).exe ./...

# Clean build files
clean:
	@echo "Cleaning build directory..."
	rm -rf $(BUILD_DIR)

# Run the application
run: build
	@echo "Running $(BINARY_NAME)..."
	$(BUILD_DIR)/$(BINARY_NAME).exe

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy

# Default target
all: deps clean build

.PHONY: build clean run deps all
