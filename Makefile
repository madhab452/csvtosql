export DBURL=postgres://postgres:postgres@127.0.0.105:5432/csv_btc?sslmode=disable

.PHONY: help # include all targets 

# ------------------
# Help
# ------------------
help: ## Show command list
	@echo "Usage:"
	@echo " make [target]"
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-20s\033[0m %s\n", $$1, $$2}'

run: ## Run service
	go run main.go -fname=./csvs/BTC-USD.csv
	