package main

import "fmt"

func main() {
	go func() {
		fmt.Println("Executing my Concurrent anonymous function")
	}()

	fmt.Scanln()
}
