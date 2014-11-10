package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Printf("Sleeping for %d milliseconds\n", 500)
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("I'm awake\n")
	}
}
