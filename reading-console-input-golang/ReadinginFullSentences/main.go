package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("simple shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("->")
		text, _ := reader.ReadString('\n')
		//text = strings.Replace(text, "\n", "", -1)

		//windows have to use this
		text = strings.Replace(text, "\r\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}

	}
}
