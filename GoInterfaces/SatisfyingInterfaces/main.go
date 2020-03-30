package main

import "fmt"

type Employee interface {
	Language() string
	Age() int
	Random() (string, error)
}
type Engineer struct {
	Name string
}

func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}
func (e Engineer) Random() (string, error) {
	return "test", nil
}
func (e Engineer) Age() int {
	return 29
}
func main() {
	var programmers []Employee
	elliot := Engineer{Name: "Elliot"}

	programmers = append(programmers, elliot)
	fmt.Println(elliot)
	fmt.Println(elliot.Language())
	fmt.Println(elliot.Random())
	fmt.Println(elliot.Age())

	fmt.Println(programmers[0].Age())
}

/* package main

type Employee interface {
	Language() string
	Age() int
	Random() (string, error)
}

type Engineer struct {
	Name string
}

func (e *Engineer) Language() string {
	return e.Name + " programs in Go"
}

func main() {
	// This will throw an error
	var programmers []Employee
	elliot := Engineer{Name: "Elliot"}
	// Engineer does not implement the Employee interface
	// you'll need to implement Age() and Random()
	programmers = append(programmers, elliot)
}
*/
