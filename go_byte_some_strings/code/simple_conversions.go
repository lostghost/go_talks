package main

import (
	"fmt"
	"os"
)

// START OMIT
func main() {
	strVal := "Hello world"
	byteVal := []byte(strVal)
	strFromBytes := string(byteVal)

	fmt.Println(strVal)
	fmt.Println(strFromBytes)
	os.Stdout.Write(byteVal)
}

// END OMIT
