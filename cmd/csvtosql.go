package cmd

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/madhab452/csvtosql/cmd/qb"
	"github.com/madhab452/csvtosql/cmd/reader"
	"github.com/madhab452/csvtosql/cmd/sqlutil"
)

var filePattern = regexp.MustCompile(".csv$")

// Cts means csv to sql
type Cts struct {
	Fpath string
	DB    *sql.DB
	Log   *logrus.Entry
}

// New Creates an instance of CsvToSql
func New(ctx context.Context, log *logrus.Entry) (*Cts, error) {
	fpath := flag.String("f", "", "Csv file to import. (Required)")
	database := flag.String("db", "postgres", "Database {postgres|mysql|elastic}.")
	dbURL := flag.String("dburl", "", "Db connection url")

	help := flag.Bool("help", false, "")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *fpath == "" {
		fmt.Println("missing required arguments -f")
		fmt.Println("Usase: ")
		flag.PrintDefaults()
		os.Exit(2)
	}

	if *database != "" {
		if ok := map[string]bool{"postgres": true, "mysql": true, "elastic": true}[*database]; !ok {
			fmt.Println("invalid database")
			fmt.Println("Usase: ")
			flag.PrintDefaults()
			os.Exit(2)
		}
	}

	if *dbURL == "" {
		fmt.Printf("missing db connection url.")
		fmt.Println("Usase: ")
		flag.PrintDefaults()
		os.Exit(2)
	}

	if !filePattern.MatchString(*fpath) {
		return nil, fmt.Errorf("invalid filename, expected csv file")
	}

	db, err := sql.Open("postgres", *dbURL)
	if err != nil {
		return nil, fmt.Errorf("sql.Open(): %w", err)
	}

	return &Cts{
		Fpath: *fpath,
		Log:   log,
		DB:    db,
	}, nil
}

// Exec validates csv data, prepare sql query and runs query against specified db
func (cs *Cts) Exec() error {
	f, err := os.Open(cs.Fpath)
	if err != nil {
		return fmt.Errorf("os.Open(): %w", err)
	}
	defer close(f, cs.Log)

	r := reader.NewReader(f)
	headers, chunks, err := r.ReadChunks(10000)
	if err != nil {
		return fmt.Errorf("reader.ReadChunks(): %w", err)
	}
	if headers == nil {
		return fmt.Errorf("empty file")
	}

	if len(chunks[0]) <= 1 {
		return fmt.Errorf("atleast one record is required")
	}

	tblname := sqlutil.ToTableName(cs.Fpath)

	createTableQuery := qb.CreateTbl(func(qb *qb.CreateTblBuilder) {
		qb.Table(tblname)
		for _, col := range headers {
			qb.AddCol(sqlutil.ToColumnName(col))
		}
	}).ToSQL()

	//nolint: gocritic
	if _, err := cs.DB.Query(createTableQuery); err != nil {
		return fmt.Errorf("cs.DB.Query(createTableQuery): %w", err)
	}

	for i, chunk := range chunks {
		insertQuery := qb.Insert(func(qb *qb.InsertBuilder) {
			qb.Table(tblname)
			for _, col := range headers {
				qb.AddCol(sqlutil.ToColumnName(col))
			}
			for _, row := range chunk {
				qb.AddRow(row)
			}
		}).ToSQL()

		rows, err := cs.DB.Query(insertQuery)
		if err != nil {
			return fmt.Errorf("cs.DB.Query(insertQuery): %w", err)
		}
		close(rows, cs.Log)
		fmt.Printf("%v records inserted \n", (i+1)*len(chunk))
	}

	return nil
}

// close close anything that implements io.Closer
func close(r io.Closer, log *logrus.Entry) {
	err := r.Close()
	if err != nil {
		log.WithError(err).Printf("r.Close()")
	}
}
