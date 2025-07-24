.DEFAULT_GOAL := help

# Directories and files
DIST_DIR = ./bin
CLI_PATH = ./cmd/cli
API_PATH = ./cmd/api

# Parameters
OUT_DIR = ./output
IMG_SIZE = 5000

run-cli:   ## Run the CLI with args (e.g. make run IMG=img.jpg FILTER=grayscale)
	go run $(CLI_PATH) -img $(IMG) -outDir $(OUT_DIR) -filter $(FILTER)

run-api:   ## Start the REST API server
	go run $(API_PATH)

list:      ## List the avaliable filter pipelines
	go run $(CLI_PATH) --list

bench:     ## Run a benchmark of a specific filter in both serial and concurrent modes
	go test -bench=. -run=^$$ -benchmem ./engines -args -filter $(FILTER) -imageSize $(IMG_SIZE)

build-cli: ## Build the CLI binary. The binary will be named "image-filters-cli"
	mkdir -p $(DIST_DIR)
	go build -o $(DIST_DIR)/image-filters-cli $(CLI_PATH)

build-api: ## Build the API binary. The binary will be named "image-filters-api"
	mkdir -p $(DIST_DIR)
	go build -o $(DIST_DIR)/image-filters-api $(API_PATH)

clean:     ## Remove old binaries
	rm -r $(DIST_DIR)/*

help:      ## Show help for each make command
	@echo 'Makefile commands:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'