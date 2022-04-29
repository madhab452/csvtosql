package cmd

import (
	"context"
	"database/sql"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	_ "github.com/lib/pq"

	"github.com/madhab452/csvtosql/cmd/sqlgen"
	"github.com/madhab452/csvtosql/cmd/util"
)

var filePattern = regexp.MustCompile(".csv$")

type CsvToSql struct {
	Fname string
	DB    *sql.DB
}

type Option struct {
	Fname string
	DBURL string
}

// New Creates an instance of CsvToSql
func New(ctx context.Context, log *log.Logger, opts *Option) (*CsvToSql, error) {
	fname := flag.String("fname", "", "csv file")
	flag.Parse()

	if !filePattern.MatchString(*fname) {
		return nil, fmt.Errorf("invalid filename, expected csv file")
	}

	db, err := sql.Open("postgres", opts.DBURL)
	if err != nil {
		return nil, fmt.Errorf("sql.Open(): %w", err)
	}

	return &CsvToSql{
		Fname: *fname,
		DB:    db,
	}, nil
}

// Do validates csv data, prepare sql query and runs query against specified db
func (cs *CsvToSql) Do() error {
	f, err := os.Open(cs.Fname)
	if err != nil {
		return fmt.Errorf("os.Open(): %w", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)

	records, err := csvReader.ReadAll()

	if err != nil {
		return fmt.Errorf("csvReader.ReadAll(): %w", err)
	}

	if len(records) <= 1 {
		return errors.New("len(records): csv file must have at least one record to insert")
	}
	sqlgen := sqlgen.New(&records)
	tblName := util.ToTableName(cs.Fname)

	createTableQuery := sqlgen.CreateTblQuery(tblName)

	if _, err := cs.DB.Query(createTableQuery); err != nil {
		return err
	}

	insertQuery := sqlgen.InsertQuery(tblName)
	if _, err := cs.DB.Query(insertQuery); err != nil {
		return err
	}

	return nil
}
