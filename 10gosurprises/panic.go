package main

import "fmt"

// Start OMIT
func main() {
	doSomething()
	fmt.Println("Made it through")
}

func doSomething() {
	doSomethingRisky()
}

func doSomethingRisky() {
	panic("FREAK OUT")
}

// End OMIT
