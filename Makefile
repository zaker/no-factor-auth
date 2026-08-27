BINARY_NAME := no-factor-auth
BIN_DIR := .tmp/bin

.PHONY: build test
build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME) .

test:
	go test ./...

