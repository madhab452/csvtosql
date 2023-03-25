// Package sqlutil has some helper function for sql.
package sqlutil

import (
	"regexp"
	"strconv"
	"strings"
)

// Characters not in the given list
var re = regexp.MustCompile(`[^a-zA-Z0-9\_]`)

func rmeoveNextUnderscore(str string) string {
	var sb strings.Builder
	// replace multiple _ with single _. given
	sb.WriteByte(str[0])
	for i := 1; i < len(str); i++ {
		if str[i] == '_' && str[i-1] == '_' {
			continue
		} else {
			sb.WriteByte(str[i])
		}
	}
	return sb.String()
}

var ucounter int

func removeSpecialChars(s string) string {
	return strings.ReplaceAll(s, "'", "")
}

// ToColumnName returns a nice name for table column.
func ToColumnName(h string) string {
	_, after, ok := strings.Cut(h, "=>")
	if ok {
		return removeSpecialChars(after)
	}

	if h == "" {
		ucounter++
		return "unknown_" + strconv.Itoa(ucounter)
	}
	mutheader := re.ReplaceAllString(h, "_")
	if len(mutheader) == 1 {
		return mutheader
	}

	return rmeoveNextUnderscore(mutheader)
}

// ToTableName returns a nice name for table.
// example given ./csvs/BTC-USD-2.csv returns btc_usd_2.
func ToTableName(fname string) string {
	var res []byte
	for i := range fname {
		ch := fname[i]
		if ch == '/' { // reset if its a path, not a filename
			res = []byte{}
		} else {
			if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9' || ch == '_' {
				res = append(res, ch)
			} else {
				res = append(res, '_')
			}
		}
	}
	str := string(res)

	if str[len(str)-4:] == "_csv" { // remove extension
		str = str[0 : len(str)-4]
	}

	str = strings.ToLower(str)
	str = rmeoveNextUnderscore(str)

	return str
}
