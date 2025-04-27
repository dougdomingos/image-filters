.DEFAULT_GOAL := help

# Directories and files
DIST_DIR = ./bin
BIN_FILE = image-filters
CLI_PATH = ./cli

# Parameters
MODE = serial
OUTPUT_DIR = ./output
IMG_SIZE = 5000

run:   ## Run the CLI with args (e.g. make run IMG_PATH=img.jpg FILTER=grayscale MODE=serial)
	go run $(CLI_PATH) -img $(IMG_PATH) -outputDir $(OUTPUT_DIR) -filter $(FILTER) -mode $(MODE)

bench: ## Run a benchmark of a specific filter in both serial and concurrent modes
	go test -bench=. -run=^$$ -benchmem ./engines -args -filter $(FILTER) -imageSize $(IMG_SIZE)

build: ## Build the CLI binary. The binary name can be specified through the "BIN_FILE" flag
	mkdir -p $(DIST_DIR)
	go build -o $(DIST_DIR)/$(BIN_FILE) $(CLI_PATH)

clean: ## Remove old binaries
	rm -r $(DIST_DIR)/*

help:  ## Show help for each make command
	@echo 'Makefile commands:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'