package main

import (
	"fmt"
	"math/rand"
)

//把channel傳進來
func CalculateValue(values chan int) {
	//隨機產生數字
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)

	//把數字送到values通到
	//發送一個值到一個信道
	values <- value
}

func main() {
	fmt.Println("Go Channel Tutorial")

	//建立新的channel
	//讓後面的CalculateValue可以使用
	//創建myChannel，這是一個類型的頻道 int
	values := make(chan int)

	//延遲關閉通道,在main func 結束後再關閉
	defer close(values)

	go CalculateValue(values)

	//接收values通道的值
	//這裡會進行block,直到從通道接收到值才會繼續執行完main func
	//從通道接收值
	value := <-values
	fmt.Println(value)
}
