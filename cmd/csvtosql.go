package cmd

import (
	"context"
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	_ "github.com/lib/pq"

	"github.com/madhab452/csvtosql/cmd/qb"
	"github.com/madhab452/csvtosql/cmd/sqlutil"
)

var filePattern = regexp.MustCompile(".csv$")

// CsvToSql ...
type CsvToSql struct {
	Fname string
	DB    *sql.DB
	Log   *log.Logger
}

// Option holds configuration option for command.
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
		Log:   log,
		DB:    db,
	}, nil
}

// Exec validates csv data, prepare sql query and runs query against specified db
func (cs *CsvToSql) Exec() error {
	f, err := os.Open(cs.Fname)
	if err != nil {
		return fmt.Errorf("os.Open(): %w", err)
	}
	defer Close(f, cs.Log)

	csvReader := csv.NewReader(f)

	records, err := csvReader.ReadAll()
	if err != nil { // errors like wrong number if fields in csv
		return fmt.Errorf("csvReader.ReadAll(): %w", err)
	}
	if len(records) <= 1 {
		return fmt.Errorf("atleast one record is required")
	}

	tblname := sqlutil.ToTableName(cs.Fname)

	createTableQuery := qb.CreateTbl(func(qb *qb.CreateTblBuilder) {
		qb.Table(tblname)
		for _, col := range records[0] {
			qb.AddCol(sqlutil.ToColumnName(col))
		}
	}).ToSql()

	insertQuery := qb.Insert(func(qb *qb.InsertBuilder) {
		qb.Table(tblname)

		for _, col := range records[0] {
			qb.AddCol(sqlutil.ToColumnName(col))
		}
		for _, row := range records[1:] {
			qb.AddRow(row)
		}
	}).ToSql()

	if _, err := cs.DB.Query(createTableQuery); err != nil {
		return fmt.Errorf("cs.DB.Query(createTableQuery): %w", err)
	}

	if _, err := cs.DB.Query(insertQuery); err != nil {
		return fmt.Errorf("cs.DB.Query(insertQuery): %w", err)
	}

	return nil
}

func Close(r io.Closer, log *log.Logger) {
	err := r.Close()
	if err != nil {
		log.Println(err)
	}
}
