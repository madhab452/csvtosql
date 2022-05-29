package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	_ "github.com/lib/pq"

	"github.com/madhab452/csvtosql/cmd/qb"
	"github.com/madhab452/csvtosql/cmd/reader"
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

	reader := reader.NewReader(f)
	headers, chunks, err := reader.ReadChunks(10000)
	if err != nil {
		return fmt.Errorf("reader.ReadChunks(): %w", err)
	}
	if headers == nil {
		return fmt.Errorf("empty file")
	}

	if len(chunks[0]) <= 1 {
		return fmt.Errorf("atleast one record is required")
	}

	tblname := sqlutil.ToTableName(cs.Fname)

	createTableQuery := qb.CreateTbl(func(qb *qb.CreateTblBuilder) {
		qb.Table(tblname)
		for _, col := range headers {
			qb.AddCol(sqlutil.ToColumnName(col))
		}
	}).ToSql()

	if _, err := cs.DB.Query(createTableQuery); err != nil {
		return fmt.Errorf("cs.DB.Query(createTableQuery): %w", err)
	}

	for i, chunk := range chunks {
		insertQuery := qb.Insert(func(qb *qb.InsertBuilder) {
			qb.Table(tblname)
			for _, col := range headers {
				qb.AddCol(sqlutil.ToColumnName(col))
			}
			for _, row := range chunk { //remove header
				qb.AddRow(row)
			}
		}).ToSql()

		rows, err := cs.DB.Query(insertQuery)
		if err != nil {
			return fmt.Errorf("cs.DB.Query(insertQuery): %w", err)
		}
		Close(rows, cs.Log)
		fmt.Printf("%v records inserted \n", (i+1)*len(chunk))
	}

	return nil
}

func Close(r io.Closer, log *log.Logger) {
	err := r.Close()
	if err != nil {
		log.Println(err)
	}
}
