# Makefile

# Define the Go binary name
BINARY_NAME=packonix.bin

# Build the Go application
build:
	go build -o $(BINARY_NAME)

# Run the Go application
run: build
	./$(BINARY_NAME)

# Clean the build artifacts
clean:
	rm -f $(BINARY_NAME)

# Run tests
test:
	go test ./...

# Install dependencies
deps:
	go mod tidy

.PHONY: build run clean test deps
