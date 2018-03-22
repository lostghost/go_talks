package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// START OMIT
func main() {
	v := []byte{Bx("1100001"), Bx("1100010"), Bx("1100011")}

	fmt.Printf("%s", v)
}

// END OMIT

func Bx(in string) byte {
	v, err := strconv.ParseUint(regexp.MustCompile("[^01]").ReplaceAllString(in, ""), 2, 8)
	if err != nil {
		panic(err)
	}
	return byte(v)
}
