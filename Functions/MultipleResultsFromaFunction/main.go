package main

import (
	"fmt"
)

func myfunction(firstName string, lastName string) (string, error) {
	return firstName + " " + lastName, nil
}

func main() {
	fmt.Println("Hello World")

	// we can assign the results to multiple variables
	// by defining their names in a comma separated list
	// like so:
	fullName, err := myfunction("Elliot", "Forbes")
	if err != nil {
		fmt.Println("Handle Error Case")
	}
	fmt.Println(fullName)
}
