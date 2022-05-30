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
func (ib *InsertBuilder) Table(tblname string) *InsertBuilder {
	ib.insert.tblname = tblname

	return ib
}

// AddCols add cols.
func (ib *InsertBuilder) AddCol(colname string) *InsertBuilder {
	ib.insert.cols = append(ib.insert.cols, colname)

	return ib
}

// AddRow adds row.
func (ib *InsertBuilder) AddRow(row []string) *InsertBuilder {
	ib.insert.rows = append(ib.insert.rows, row)

	return ib
}

// ToSql converts the QueryBuilder to sql query.
func (ib *InsertBuilder) ToSql() string {
	var pgrows []string
	for i, row := range ib.rows {
		for j, v := range row {
			ib.rows[i][j] = fmt.Sprintf("'%s'", v)
		}
		pgrow := strings.Join(ib.rows[i], ",")
		pgrowwrapped := fmt.Sprintf("(%s)", pgrow)
		pgrows = append(pgrows, pgrowwrapped)
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES %s;", ib.tblname, strings.Join(ib.cols, ", "), strings.Join(pgrows, ", "))
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
