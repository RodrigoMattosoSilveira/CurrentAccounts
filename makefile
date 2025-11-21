# Project variables
BINARY_NAME=web
CMD_DIR=cmd/web
BUILD_DIR=bin

# Default target
all: build

## Build the application
build:
	@echo "▶ Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	GO111MODULE=on go build -o $(BUILD_DIR)/$(BINARY_NAME) $(CMD_DIR)

## Run the application
run: build
	@echo "▶ Running $(BINARY_NAME)..."
	@./$(BUILD_DIR)/$(BINARY_NAME)

## Run all tests
test:
	@echo "▶ Running tests..."
	go test -v ./... | grep -E 'PASS|FAIL' | tr -d ' '

# Run all tests
test_g:
	@echo "▶ Running tests..."
	UPDATE_GOLDEN_FILES=true go test -v ./... | grep -E 'PASS|FAIL' | tr -d ' '

## Run all tests with coverage
test_c:
	@echo "▶ Running tests with coverage ..."
	@go test ./... -cover

## Run all tests with verbose
test_v:
	@echo "▶ Running tests with verbose ..."
	@go test ./... -v
# Tidy dependencies
tidy:
	@echo "▶ Tidying go.mod..."
	@go mod tidy

## Clean build artifacts
clean:
	@echo "▶ Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)

## Show help
help:
	@echo ""
	@echo "Available commands:"
	@echo "  make build   - Build the binary"
	@echo "  make run     - Build and run the binary"
	@echo "  make test    - Run tests"
	@echo "  make test_c  - Run tests with coverage"
	@echo "  make test_g  - Run tests with golden files"
	@echo "  make test_v  - Run tests with golden verbose"
	@echo "  make tidy    - Run go mod tidy"
	@echo "  make clean   - Remove build directory"
	@echo "  make help    - Show this help message"
	@echo ""
