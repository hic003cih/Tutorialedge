package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Go Sorting Tutorial")

	myInts := []int{1, 3, 2, 6, 3, 4}
	fmt.Println(myInts)

	// we can use the sort.Ints
	sort.Ints(myInts)
	//sort.Ints(myInts)
	fmt.Println(myInts)
}

/* package main

import "fmt"

func main() {
    fmt.Println("Go Sorting Tutorial")

    // our unsorted int array
    unsorted := []int{1,3,2,6,3,4}
	fmt.Println(unsorted)
} */
