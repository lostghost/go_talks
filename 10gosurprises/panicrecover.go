package main

import "fmt"

// Start OMIT
func main() {
	doSomething()
	fmt.Println("Made it through")
}

func doSomething() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered", r)
		}
	}()
	doSomethingRisky()
}

func doSomethingRisky() {
	// Oh nose
	panic("FREAK OUT")
}

// End OMIT
