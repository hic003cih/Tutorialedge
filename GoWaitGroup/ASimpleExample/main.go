package main

import (
	"fmt"
	"sync"
)

func myFunc(waitgroup *sync.WaitGroup) {
	fmt.Println("Inside my goroutine")
	waitgroup.Done()
}
func main() {
	var waitgroup sync.WaitGroup

	waitgroup.Add(1)

	go myFunc(&waitgroup)

	waitgroup.Wait()

	fmt.Println("Finished Execution")

}
