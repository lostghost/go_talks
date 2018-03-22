package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// START OMIT
func main() {
	minVal := Bx("00000000")
	maxVal := Bx("11111111")
	myVal := Bx("01100001")

	fmt.Println("Min:", minVal)
	fmt.Println("Max:", maxVal)
	fmt.Println("Mine:", myVal)
}

// END OMIT

func Bx(in string) byte {
	v, err := strconv.ParseUint(regexp.MustCompile("[^01]").ReplaceAllString(in, ""), 2, 8)
	if err != nil {
		panic(err)
	}
	return byte(v)
}
