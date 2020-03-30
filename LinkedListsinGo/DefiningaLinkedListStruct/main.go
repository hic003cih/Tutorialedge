package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("Go Linked Lists Tutorial")
	mylist := list.New()
	mylist.PushBack(1)
	mylist.PushFront(4)
	mylist.InsertAfter(4, 5)

	for element := mylist.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}

/*
func main() {
    fmt.Println("Go Linked Lists Tutorial")
    mylist := list.New()
    mylist.PushBack(1)
    mylist.PushFront(2)
    // we now have a linked list with '1' at the back of the list
	// and '2' at the front of the list.

	for element := mylist.Front(); element != nil; element = element.Next() {
    // do something with element.Value
    fmt.Println(element.Value)
    }
} */
