# madhab452/csvtosql

csvtosql is a cli tool to convert csv data to sql db tables,
Oftentimes we have to work with data in Excel because product owner and users are more comfortable with excel. But, for a developer like me feel more comfortable to run sql queries. This will automate your data migration from excel to sql.

## How to run?

Clone/Fork the repo, rename .env.sh.example to .env.sh and ajust environment variable to your machine.
run `make run `

#### TODO

- [ ] Improve speed for large files.
- [ ] Handle more edge cases, for example when no of cols is greater than header
- [ ] Add mysql as optional db
- [ ] Customize linter rule
	example https://raw.githubusercontent.com/golangci/golangci-lint/master/.golangci.yml
