package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/madhab452/csvtosql/cmd"
)

func main() {
	start := time.Now()
	context := context.Background()
	l := log.Default()
	opts := &cmd.Option{
		DBURL: os.Getenv("DBURL"),
		Fpath: "",
	}

	csvtosql, err := cmd.New(context, l, opts)

	if err != nil {
		fmt.Printf("cmd.New(): %v \n", err)
		return
	}

	if err := csvtosql.Exec(); err != nil {
		fmt.Printf("csvtosql.Exec(): %v", err)
		return
	}

	log.Println("time taken:", time.Since(start))
}
