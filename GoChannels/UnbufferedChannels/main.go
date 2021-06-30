//Unbuffered Channels無緩衝通道
//會因為只有一個通道,而產生block的問題
//當一個goroutine 傳資料給到channel
//goroutine 就會blocked直到接收到channel送過來的資料
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(c chan int) {
	//goroutine blocked
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	time.Sleep(1000 * time.Millisecond)
	//接收到channel的資料,解開blocked
	c <- value
	fmt.Println("Only Executes after another goroutine performs a receive on the channel")
}

func main() {
	fmt.Println("Go Channel Tutorial")

	valueChannel := make(chan int)
	defer close(valueChannel)

	go CalculateValue(valueChannel)
	go CalculateValue(valueChannel)

	//執行了兩個GO func
	//但最後卻只印出一個
	//因為第二個go func 再執行的時候被第一個go func搶了channel,所以被 blocked
	//在第一個go func釋放channel之前,就已經結束func了
	values := <-valueChannel

	//在放一個通道就可以解決channel被搶的問題
	//values2 := <-valueChannel

	fmt.Println(values)

	//mt.Println(values2)
}
