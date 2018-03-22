package main

import (
	"fmt"
	"unicode/utf8"
)

// START OMIT
func main() {
	s := "ab€"

	fmt.Println("string length:", len(s))
	fmt.Println("string slice:", s[:3])
	fmt.Printf("string bytes: %08b\n", []byte(s))

	fmt.Println("runes in string:", utf8.RuneCountInString(s))
}

// END OMIT
