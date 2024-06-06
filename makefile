# Makefile

# Define variables
SRC_DIR := ./src
MAIN := main.go
APP_NAME := GO

# Default target
.PHONY: start
start:
	@echo "Starting $(APP_NAME)"
	@nodemon --exec go run ./$(MAIN) --signal SIGTERM

# Target to build the executable
.PHONY: build
build:
	@echo "Building $(APP_NAME)"
	@go build -o $(APP_NAME) $(SRC_DIR)/$(MAIN)

# Target to run the built executable
.PHONY: run
run: build
	@echo "Running $(APP_NAME)"
	@./$(APP_NAME)

# Target to clean up built files
.PHONY: clean
clean:
	@echo "Cleaning up"
	@rm -f $(APP_NAME)
