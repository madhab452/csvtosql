package sqlutil

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`[^\w]`)

// ToColumnName returns a nice name for table column.
func ToColumnName(header string) string {
	return toSQLFriendlyName(header)
}

// ToTableName returns a nice name for table
func ToTableName(fname string) string {
	// return toSqlFriendlyName(fname) // todo: create table name from file
	return "csvtosql"
}

func toSQLFriendlyName(str string) string {
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, " ", "_")
	str = re.ReplaceAllString(str, "")

	return str
}
