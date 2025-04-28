# Makefile

APP_NAME := itinerary
BIN_DIR := bin

.PHONY: all build run test tidy clean docker

all: build

build:
	@echo "→ Building binary..."
	GOOS=$(shell go env GOOS) \
	GOARCH=$(shell go env GOARCH) \
	go build -o $(BIN_DIR)/$(APP_NAME) cmd/server/main.go

run: build
	@echo "→ Running locally..."
	./$(BIN_DIR)/$(APP_NAME)

test:
	@echo "→ Running tests..."
	go test ./internal/...

tidy:
	@echo "→ Tidying modules..."
	go mod tidy

clean:
	@echo "→ Cleaning up..."
	rm -rf $(BIN_DIR)

docker:
	@echo "→ Building Docker image..."
	docker build -t $(APP_NAME)-service .