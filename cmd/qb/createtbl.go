// Package qb provides fluent api to create sql queries
package qb

import (
	"fmt"
	"strings"
)

type createTable struct {
	tblname string
	cols    []string
}

// CreateTblBuilder CreateTblBuilder
type CreateTblBuilder struct {
	createTable
}

// Table name of the table.
func (ctb *CreateTblBuilder) Table(tblname string) *CreateTblBuilder {
	ctb.createTable.tblname = tblname

	return ctb
}

// AddCol add column.
func (ctb *CreateTblBuilder) AddCol(col string) *CreateTblBuilder {
	ctb.createTable.cols = append(ctb.createTable.cols, col)
	return ctb
}

// ToSQL returns sql query.
func (ctb *CreateTblBuilder) ToSQL() string {
	var colspg []string

	for _, col := range ctb.cols {
		colspg = append(colspg, fmt.Sprintf("%s VARCHAR(255)", col))
	}

	sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s);`, ctb.tblname, strings.Join(colspg, ", "))

	return sql
}

type createTableBuild func(*CreateTblBuilder)

// CreateTbl return builder to create sql query.
func CreateTbl(action createTableBuild) *CreateTblBuilder {
	builder := &CreateTblBuilder{
		createTable: createTable{},
	}
	action(builder)

	return builder
}
