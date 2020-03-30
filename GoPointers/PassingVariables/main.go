package main

import "fmt"

func YearsUntilRetirement(age int) {
	fmt.Println(100 - age)
}

func main() {
	var age *int
	age = new(int)
	*age = 26

	YearsUntilRetirement(*age)
}
