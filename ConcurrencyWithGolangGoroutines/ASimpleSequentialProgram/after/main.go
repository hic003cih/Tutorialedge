package main

import (
	"fmt"
	"time"
)

// notice we've not changed anything in this function
// when compared to our previous sequential program
func compute(value int) {
	for i := 0; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("Goroutine Tutorial")

	// notice how we've added the 'go' keyword
	// in front of both our compute function calls

	/* //如果前面有先放一個有沒go 開頭的func
	//會先執行完上面那個才會執行有go的
	compute(10)
	go compute(15)
	//如果前面有先放一個有go 開頭的func
	//會執行go func以後,馬上又在執行沒有go的,幾乎同時執行
	go compute(10)
	compute(15) */

	go compute(10)
	go compute(15)

	var input string

	//main func在 goroutines 執行之前就回執行完了
	//放一個等待鍵盤輸入,讓goroutines有時間可以執行完
	fmt.Scanln(&input)
}
