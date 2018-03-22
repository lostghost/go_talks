package main

import "fmt"

// START OMIT
func main() {
	v := []byte("€")

	fmt.Printf("Quoted: %+q\n", v)
	fmt.Printf("Bin: %08b\n", v)
}

// END OMIT
