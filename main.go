package main

import (
	"fmt"

	"github.com/vitorsoratto/csv-to-db/cli"
)

func main() {
	if err := cli.CreateCommand().Execute(); err != nil {
		fmt.Println(err)
	}
}
