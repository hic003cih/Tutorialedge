package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title  string `json:"title"`
	Author Author
}

type Author struct {
	Sales     int  `json:"book_sales"`
	Age       int  `json:"age"`
	Developer bool `json:"is_developer"`
}

type SensorReading struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Time     string `json:"time"`
}

func main() {
	author := Author{Sales: 3, Age: 25, Developer: true}
	book := Book{Title: "Learning Concurrency in Python", Author: author}

	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))

	byteArray2, err2 := json.MarshalIndent(book, "", "  ")
	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(string(byteArray2))

	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`

	var reading SensorReading
	err = json.Unmarshal([]byte(jsonString), &reading)
	fmt.Printf("%+v\n", reading)

	var reading2 map[string]interface{}
	err2 = json.Unmarshal([]byte(jsonString), &reading2)
	fmt.Printf("%+v\n", reading2)

}

/* type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Sales     int  `json:"book_sales"`
	Age       int  `json:"age"`
	Developer bool `json:"is_developer"`
}

type SensorReading struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
	Time     string `json:"time"`
}

func main() {

	author := Author{Sales: 3, Age: 25, Developer: true}
	book := Book{Title: "Learning Concurrency in Python", Author: author}

	byteArray, err := json.Marshal(book)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(byteArray))

	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`

	var reading SensorReading
	err = json.Unmarshal([]byte(jsonString), &reading)
	fmt.Printf("%+v\n", reading)
} */
