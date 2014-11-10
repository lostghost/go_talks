package main

import "fmt"

type Partner struct {
	FirstName string
	LastName  string
}

// Start OMIT
type Person struct {
	FirstName string
	LastName  string
	Partner   Partner
}

func (p Person) String() string { // HL
	return p.FirstName + " " + p.LastName // HL
} // HL

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
