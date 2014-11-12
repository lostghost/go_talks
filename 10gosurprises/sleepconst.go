package main

import (
	"fmt"
	"time"
)

// Start OMIT
const Delay = 500 // HL

func main() {
	for {
		fmt.Printf("Sleeping for %d milliseconds\n", Delay) // HL
		time.Sleep(Delay * time.Millisecond)                // HL
		fmt.Printf("I'm awake\n")
	}
}

// End OMIT
