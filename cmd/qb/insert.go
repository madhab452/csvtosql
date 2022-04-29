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

type InsertBuilder struct {
	insert
}

// Table name of the table
func (ctb *InsertBuilder) Table(tblname string) *InsertBuilder {
	ctb.insert.tblname = tblname
	return ctb
}

// AddCols add cols
func (ctb *InsertBuilder) AddCol(colname string) *InsertBuilder {
	ctb.insert.cols = append(ctb.cols, colname)
	return ctb
}

// AddRow adds row
func (ctb *InsertBuilder) AddRow(row []string) *InsertBuilder {
	ctb.insert.rows = append(ctb.rows, row)
	return ctb
}

// ToSql converts the QueryBuilder to sql query
func (ct *InsertBuilder) ToSql() string {
	var rowspg []string

	for _, row := range ct.rows {
		rowspg = append(rowspg, strings.Join(row, "','"))
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES ('%s');", ct.tblname, strings.Join(ct.cols, ","), strings.Join(rowspg, "'),('"))
}

type insertBuild func(*InsertBuilder)

func Insert(action insertBuild) *InsertBuilder {
	builder := &InsertBuilder{}
	action(builder)
	return builder
}
