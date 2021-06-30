package main

import "fmt"

func main() {
	// we make our anonymous function concurrent using `go`
	go func() {
		fmt.Println("Executing my Concurrent anonymous function")
	}()
	// we have to once again block until our anonymous goroutine
	// has finished or our main() function will complete without
	// printing anything
	fmt.Scanln()
}
