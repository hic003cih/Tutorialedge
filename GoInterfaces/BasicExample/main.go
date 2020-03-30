package main

import "fmt"

func myFunc(a interface{}) {
	fmt.Println(a)
}
func main() {
	myAge := 25
	myFunc(myAge)
}

/* package main

import (
    "fmt"
)

func myFunc(a interface{}) {
    fmt.Println(a)
}

func main() {
    var my_age int
    my_age = 25

    myFunc(my_age)
} */
