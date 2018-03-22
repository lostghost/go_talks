package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// START OMIT
func main() {
	euro := []byte{Bx("11100010"), Bx("10000010"), Bx("10101100")}
	rofl := []byte{Bx("11110000"), Bx("10011111"), Bx("10100100"), Bx("10100011")}
	hole := []byte{Bx("11110000"), Bx("10011111"), Bx("10010101"), Bx("10110011")}

	fmt.Printf("%s\n", euro)
	fmt.Printf("%s\n", rofl)
	fmt.Printf("%s\n", hole)
}

// END OMIT

func Bx(in string) byte {
	v, err := strconv.ParseUint(regexp.MustCompile("[^01]").ReplaceAllString(in, ""), 2, 8)
	if err != nil {
		panic(err)
	}
	return byte(v)
}
