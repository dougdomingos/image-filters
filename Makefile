.DEFAULT_GOAL := help

DIST_DIR = ./bin
BIN_FILE = image-filters
CLI_PATH = ./cli

run:   ## Run the CLI with args (e.g. make run IMG_PATH=img.jpg FILTER=grayscale)
	go run $(CLI_PATH) -img $(IMG_PATH) -filter $(FILTER)

build: ## Build the CLI binary. The binary name can be specified through the "BIN_FILE" flag
	mkdir -p $(DIST_DIR)
	go build -o $(DIST_DIR)/$(BIN_FILE) $(CLI_PATH)

clean: ## Remove old binaries
	rm -r $(DIST_DIR)/*

help: ## Show help for each make command
	@echo 'Makefile commands:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'