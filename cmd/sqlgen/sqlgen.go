package sqlgen

import (
	"fmt"

	"github.com/madhab452/csvtosql/cmd/util"
)

type SqlGen struct {
	records *[][]string
}

// New creates an instance of SqlGen
func New(records *[][]string) *SqlGen {
	return &SqlGen{
		records: records,
	}
}

// InsertQuery return insert statement
func (sg *SqlGen) InsertQuery(tblName string) string {
	headers := (*sg.records)[0]
	columnList := ""
	for i, header := range headers {
		columnList += util.ToColumnName(header)
		if i != len(headers)-1 {
			columnList += ", "
		}
	}
	str := fmt.Sprintf("INSERT INTO %s (%s) VALUES ", tblName, columnList)

	data := (*sg.records)[1:]
	for i, row := range data {
		str += "("
		for j, col := range row {
			str += fmt.Sprintf("'%s'", col)
			if j != len(row)-1 {
				str += ", "
			}
		}
		str += ")"
		if i != len(data)-1 {
			str += ","
		}
	}
	str += ";"
	return str
}

// CreateTblQuery return create table statement
func (sq *SqlGen) CreateTblQuery(tblname string) string {
	str := ""
	headers := (*sq.records)[0]
	for i, header := range headers {
		str += fmt.Sprintf("\t%s VARCHAR(255)", util.ToColumnName(header))
		if i != len(headers)-1 {
			str += ",\n"
		}
	}

	return fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
    %s
	);`, tblname, str)
}
