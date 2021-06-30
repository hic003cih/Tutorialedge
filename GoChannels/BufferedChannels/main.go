//buffered Channels緩衝通道
//會因為只有一個通道,而產生block的問題
//給定大小的隊列就可以解決blocked的問題
//只有當隊列的大小都滿了才會blocked
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CalculateValue(c chan int) {
	value := rand.Intn(10)
	fmt.Println("Calculated Random Value: {}", value)
	//
	//time.Sleep(1000 * time.Millisecond)
	c <- value
	fmt.Println("This executes regardless as the send is now non-blocking")
}

func main() {
	fmt.Println("Go Channel Tutorial")

	valueChannel := make(chan int, 2)
	defer close(valueChannel)

	go CalculateValue(valueChannel)
	go CalculateValue(valueChannel)

	values := <-valueChannel
	//value2 := <-valueChannel
	fmt.Println(values)
	//fmt.Println(value2)
	//延遲讓main不會馬上結束
	//沒有延遲goroutine還沒執行完,整個main執行完就會結束程式了
	time.Sleep(1000 * time.Millisecond)
}
