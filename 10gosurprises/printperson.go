package main

import "fmt"

// Start OMIT
type Person struct {
	FirstName string
	LastName  string
	Partner   Partner
}
type Partner struct {
	FirstName string
	LastName  string
}

func main() {
	Bob := Person{
		FirstName: "Robert",
		LastName:  "Smith",
		Partner: Partner{
			FirstName: "Jane",
			LastName:  "Smith"}}
	fmt.Println(Bob)
}

// End OMIT
