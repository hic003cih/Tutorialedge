package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner()
}

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
