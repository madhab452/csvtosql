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
	@go run main.go -f=./_examples/BTC-USD-2.csv -dburl="postgres://postgres:postgres@127.0.0.110:5433/csvtosql_db?sslmode=disable"

run_large: ## Run (large files with a lots of data)
	@go run main.go -f=./_examples/BTC-USD-LARGE.csv -dburl="postgres://postgres:postgres@127.0.0.110:5433/csvtosql_db?sslmode=disable"

gen_large: ## Generate large file with a log of data.
	cd _examples && go run main.go

test: ## Run Tests
	go test ./...

lint: ## Run linter
	golangci-lint run

release: ## Generate compiled binaries for different os.
	goreleaser 
