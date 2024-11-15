package util

import (
	"fmt"
	"time"
)

func Timer(name string) func() {
	start := time.Now()

	return func() {
		elapsed := time.Since(start)
		fmt.Printf("--[ %s took %s\n", name, elapsed)
	}
}
