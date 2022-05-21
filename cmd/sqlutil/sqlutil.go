package sqlutil

import (
	"regexp"
	"strconv"
	"strings"
)

// Characters not in the given list
var re = regexp.MustCompile(`[^a-zA-Z0-9\_]`)

func remove(str string) string {
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

// ToColumnName returns a nice name for table column.
func ToColumnName(header string) string {
	if header == "" {
		ucounter = ucounter + 1
		return "unknown_" + strconv.Itoa(ucounter)

	}
	mutheader := re.ReplaceAllString(header, "_")
	if len(mutheader) == 1 {
		return mutheader
	}

	return remove(mutheader)
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

	if str[len(str)-4:] == "_csv" { // remove extention
		str = str[0 : len(str)-4]
	}

	str = strings.ToLower(str)
	str = remove(str)

	return str
}
