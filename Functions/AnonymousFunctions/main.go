package main

import (
	"fmt"
)

func addOne() func() int {
	var x int
	// we define and return an
	// anonymous function which in turn
	// returns an integer value
	return func() int {
		// this anonymous function
		// has access to the x variable
		// defined in the parent function
		x++
		return x
	}
}

func main() {

	myFunc := addOne()
	x := myFunc()
	y := addOne()

	//執行addOne()內的
	fmt.Println(&myFunc)
	fmt.Println(myFunc()) // 2
	fmt.Println(myFunc()) // 3
	fmt.Println(myFunc()) // 4
	fmt.Println(myFunc()) // 5

	fmt.Println(&y)
	fmt.Println("y1:", y()) // 1
	fmt.Println("y2:", y()) // 2
	fmt.Println("y3:", y()) // 3
	fmt.Println("y4:", y()) // 4

	fmt.Println(&x)
	fmt.Println("x1:", x) //1
	fmt.Println("x2:", x) //1
	fmt.Println("x3:", x) //1
	fmt.Println("x4:", x) //1
}
