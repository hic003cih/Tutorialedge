package main

import (
	"fmt"
	"time"
)

func myFunc() {
	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
}

// coolFunc takes in a function
// as a parameter
func coolFunc(a func()) {
	// it then immediately calls that functino
	a()
}

func main() {
	fmt.Printf("Type: %T\n", myFunc)
	// here we call our coolFunc function
	// passing in myFunc
	coolFunc(myFunc)
}
