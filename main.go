package main

import (
	"context"
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
		Fname: "",
	}

	csvtosql, err := cmd.New(context, l, opts)

	if err != nil {
		log.Println("cmd.New():", err)
	}

	if err := csvtosql.Exec(); err != nil {
		log.Println("csvtosql.Exec():", err)
	}

	log.Println("time taken:", time.Since(start))
}
