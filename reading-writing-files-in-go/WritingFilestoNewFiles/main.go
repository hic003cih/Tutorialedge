package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	mydata := []byte("All the data I wish to write to a file")

	err := ioutil.WriteFile("myfile.data", mydata, 0777)

	if err != nil {
		fmt.Println(err)
	}

	data, err := ioutil.ReadFile("myfile.data")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
