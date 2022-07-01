# madhab452/csvtosql

csvtosql is a cli tool to convert csv data to sql db tables,
Oftentimes we have to work with data in Excel because product owner and users are more comfortable with excel. But, for a developer like me feel more comfortable to run sql queries. This will automate your data migration from excel to sql.

## How to use it?
Download the latest binary from releases and use this tool.
At very basic level this you can run 

```./csvtosql -f=./_examples/BTC-USD-2.csv -dburl="postgres://postgres:postgres@127.0.0.110:5433/csvtosql_db?sslmode=disable"```

## Future improvement:
- Support for mysql and elastic search.
- Support for postgres column type.
