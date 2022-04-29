package sqlutil

import (
	"regexp"
	"strings"
)

var Re = regexp.MustCompile(`[^\w]`)

func ToColumnName(header string) string {
	return toSqlFriendlyName(header)
}

func ToTableName(fname string) string {
	return "csvtosql"
	// return toSqlFriendlyName(fname) // todo: create table name from file
}

func toSqlFriendlyName(str string) string {
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, " ", "_")

	str = Re.ReplaceAllString(str, "")
	return str
}
