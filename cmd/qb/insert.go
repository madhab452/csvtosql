package qb

import (
	"fmt"
	"strings"
)

type insert struct {
	tblname string
	cols    []string
	rows    [][]string
}

// InsertBuilder builds insert query.
type InsertBuilder struct {
	insert
}

// Table name of the table.
func (cib *InsertBuilder) Table(tblname string) *InsertBuilder {
	cib.insert.tblname = tblname

	return cib
}

// AddCols add cols.
func (cib *InsertBuilder) AddCol(colname string) *InsertBuilder {
	cib.insert.cols = append(cib.insert.cols, colname)

	return cib
}

// AddRow adds row.
func (cib *InsertBuilder) AddRow(row []string) *InsertBuilder {
	cib.insert.rows = append(cib.insert.rows, row)

	return cib
}

// ToSql converts the QueryBuilder to sql query.
func (cib *InsertBuilder) ToSql() string {
	var rowspg []string

	for _, row := range cib.rows {
		rowspg = append(rowspg, strings.Join(row, "','"))
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES ('%s');", cib.tblname, strings.Join(cib.cols, ","), strings.Join(rowspg, "'),('"))
}

type insertBuild func(*InsertBuilder)

// Insert return insertBuilder.
func Insert(action insertBuild) *InsertBuilder {
	builder := &InsertBuilder{
		insert: insert{},
	}
	action(builder)

	return builder
}
