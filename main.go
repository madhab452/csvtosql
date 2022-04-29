package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/madhab452/csvtosql/cmd"
)

func main() {
	context := context.Background()
	l := log.Default()
	opts := &cmd.Option{
		DBURL: os.Getenv("DBURL"),
	}

	csvtosql, err := cmd.New(context, l, opts)

	if err != nil {
		fmt.Println("cmd.New(): %w", err)
	}

	if err := csvtosql.Do(); err != nil {
		fmt.Println("csvtosql.Do(): %w", err)
	}
}
