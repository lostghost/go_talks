package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Partner   Partner
}
type Partner struct {
	FirstName string
	LastName  string
}

// Start OMIT
func main() {
	Bob := Person{
		FirstName: "Robert",
		LastName:  "Smith",
		Partner: Partner{
			FirstName: "Jane",
			LastName:  "Smith"}}
	j, _ := json.Marshal(Bob)
	fmt.Println(string(j))
}

// End OMIT
