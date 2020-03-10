package main

import (
	"fmt"
)

type Employee struct {
	Name string
}

//這邊必須使用*Employee(指針)來鎖定傳入的參數和外面定義的是同一個,
//沒有使用指針的話,會和function一樣,複製一份新的來用
//func (pointer *Pointer) myMethod()
func (e *Employee) UpdateName(newName string) {
	e.Name = newName
}

func (e *Employee) PrintName() {
	fmt.Println(e.Name)
}

func main() {
	var employee Employee
	employee.Name = "Elliot"
	employee.UpdateName("Forbsey")
	employee.PrintName()
}
