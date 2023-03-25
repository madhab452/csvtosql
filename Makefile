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
	@go run main.go -f=./_examples/BTC-USD-2.csv -DBURL="postgres://postgres:postgres@127.0.0.1:5440/csvtosql_db?sslmode=disable"

gen_large: ## Generate large file with a lots of data.
	cd _examples && go run main.go

test: ## Run Tests
	go test ./...

lint: ## Run linter
	golangci-lint run

release: ## Generate compiled binaries for different os.
	goreleaser 

start-postgres: ## start postgres container
	@docker compose up -d 

stop-postgres: ## stop postgres container
	@docker compose down

connect-db: ## connect database
	@PGPASSWORD=postgres psql -h localhost -p 5440 -U postgres
