package main

import (
	"fmt"
	"time"
)

// Start OMIT
var Delay = 500 // HL

func main() {
	for {
		fmt.Printf("Sleeping for %d milliseconds\n", Delay)
		time.Sleep(Delay * time.Millisecond)
		fmt.Printf("I'm awake\n")
	}
}

// End OMIT
