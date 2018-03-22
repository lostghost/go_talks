package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

// START OMIT
func main() {
	s := "abc"
	b := []byte{Bx("1100001"), Bx("1100010"), Bx("1100011")}

	fmt.Println(reflect.DeepEqual([]byte(s), b))
	fmt.Println(reflect.DeepEqual(s, string(b)))
}

// END OMIT

func Bx(in string) byte {
	v, err := strconv.ParseUint(regexp.MustCompile("[^01]").ReplaceAllString(in, ""), 2, 8)
	if err != nil {
		panic(err)
	}
	return byte(v)
}
