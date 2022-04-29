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
	}

	csvtosql, err := cmd.New(context, l, opts)

	if err != nil {
		log.Println("cmd.New():", err)
	}

	if err := csvtosql.Do(); err != nil {
		log.Println("csvtosql.Do():", err)
	}

	fmt.Println("time taken:", time.Since(start))
}
