package main

import (
	"crypto/md5"
	"fmt"
)

// START OMIT
func main() {
	v := md5.Sum([]byte("The quick brown fox jumps over the lazy dog"))

	fmt.Printf("%08b\n", v)
	fmt.Printf("%d\n", v)
	fmt.Printf("%s\n", v)
}

// END OMIT
