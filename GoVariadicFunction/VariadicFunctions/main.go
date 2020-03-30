package main

import (
	"fmt"
)

func myVariadicFunction(args ...string) {
	fmt.Println(args)
}

func main() {
	myVariadicFunction("hello", "world")

}
