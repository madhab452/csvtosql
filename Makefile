.PHONY: help # include all targets 

# ------------------
# Help
# ------------------
help: ## Show command list
	@echo "Usage:"
	@echo " make [target]"
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-20s\033[0m %s\n", $$1, $$2}'

run: ## Run command
	. ./.env.sh; \
	go run main.go -fname=./_examples/BTC-USD-2.csv

run2: ## Run (large files with a lots of data)
	. ./.env.sh; \
	go run main.go -fname=./_examples/BTC-USD-LARGE.csv
	
test: ## Run Tests
	. ./.env.sh
	go test ./...