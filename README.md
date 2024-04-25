# madhab452/csvtosql

csvtosql is a CLI tool designed to simplify the data migration process from Excel to SQL databases. Often, product owners and users prefer working with Excel, while developers like myself find it more efficient to perform queries on SQL databases. By automating the conversion of CSV data, which can be exported from Excel, into SQL database tables, csvtosql ensures a seamless and efficient data transfer process, making it both simple and elegant for developers to use.

## How to use it?
Download the latest binary from releases and use this tool.
At very basic level this you can run 

> ./csvtosql -f=./_examples/BTC-USD-2.csv -dburl="postgres://postgres:postgres@127.0.0.110:5433/csvtosql_db?sslmode=disable"

## Future improvement:
- Support for mysql and elastic search.
- Support for postgres column type.
