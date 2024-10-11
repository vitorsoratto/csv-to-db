package util

import (
	"fmt"
	"time"
)

func Timer(name string) func() {
	start := time.Now()
	elapsed := time.Since(start)

	return func() {
		fmt.Printf("--[ %s took %s\n", name, elapsed)
	}
}
