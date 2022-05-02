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

// ToTableName returns a nice name for table.
// example given ./csvs/BTC-USD-2.csv returns btc_usd_2
// :future-improvements using regex.
func ToTableName(fname string) string {
	var res []byte
	for i, _ := range fname {
		ch := fname[i]
		if ch == '/' { // reset if its a path, not a filename
			res = []byte{}
		} else if ch == '.' {
			break
		} else {
			if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9' || ch == '_' {
				res = append(res, ch)
			} else {
				res = append(res, '_')
			}
		}
	}

	return strings.ToLower(string(res))
}

func toSQLFriendlyName(str string) string {
	str = strings.ToLower(str)
	str = strings.TrimSpace(str)
	str = strings.ReplaceAll(str, " ", "_")
	str = re.ReplaceAllString(str, "")

	return str
}
