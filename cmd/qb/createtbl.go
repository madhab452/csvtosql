package qb

import (
	"fmt"
	"strings"
)

type createTable struct {
	tblname string
	cols    []string
}

type CreateTblBuilder struct {
	createTable
}

// Table name of the table
func (ctb *CreateTblBuilder) Table(tblname string) *CreateTblBuilder {
	ctb.createTable.tblname = tblname
	return ctb
}
func (ctb *CreateTblBuilder) AddCol(col string) *CreateTblBuilder {
	ctb.createTable.cols = append(ctb.cols, col)
	return ctb
}

func (ct *CreateTblBuilder) ToSql() string {
	var colspg []string

	for _, col := range ct.cols {
		colspg = append(colspg, fmt.Sprintf("\t%s VARCHAR(255)", col))
	}

	return fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (%s);`, ct.tblname, strings.Join(colspg, ","))
}

type createTableBuild func(*CreateTblBuilder)

func CreateTbl(action createTableBuild) *CreateTblBuilder {
	builder := &CreateTblBuilder{}
	action(builder)
	return builder
}
