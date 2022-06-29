package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/madhab452/csvtosql/cmd"
)

var log *logrus.Entry

func main() {
	log = logrus.NewEntry(logrus.New())

	start := time.Now()
	ctx := context.Background()

	csvtosql, err := cmd.New(ctx, log)

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
