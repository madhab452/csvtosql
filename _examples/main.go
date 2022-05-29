// Write records into BTC-USD-LARGE.csv with sample data
package main

import (
	"bufio"
	"os"
)

const FILE = "./BTC-USD-LARGE.csv"
const TOTAL = 1000000 * 1 // 1 million

func main() {
	if _, err := os.Stat(FILE); err == nil {
		if err := os.Remove(FILE); err != nil {
			panic(err)
		}
	}

	f, err := os.OpenFile(FILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	w.WriteString("Date,Open,High,Low,Close,Adj Close,Volume\n")

	for i := 0; i < TOTAL; i++ {
		w.WriteString("2014-09-17,465.864014,468.174011,452.421997,457.334015,457.334015,21056800\n")
	}
	w.Flush()
}
