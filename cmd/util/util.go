package util

import (
	"regexp"
	"strings"
)

var Re = regexp.MustCompile(`[^\w]`)

func ToColumnName(header string) string {
	return toSqlFriendlyName(header)
}

func ToTableName(fname string) string {
	//TODO: table name doesn't looks good. Please use a different sanitizer.
	return toSqlFriendlyName(fname)
}

func toSqlFriendlyName(str string) string {
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, " ", "_")

	str = Re.ReplaceAllString(str, "")
	return str
}
