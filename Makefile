SRC_DIR := /Users/lukawivilashvili/GolandProjects/TrophyScrape/src
BUILD_DIR := bin
APP_NAME := TrophyScrape

build:
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(wildcard $(SRC_DIR)/*.go)

run: build
	@./$(BUILD_DIR)/$(APP_NAME)

test:
	@go test -v ./...
