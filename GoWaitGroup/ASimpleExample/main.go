//WaitGroups
//每次goroutines還沒執行完.main就跳出了
//雖然可以用time.sleep解決
//但用WaitGroups是更好的選擇
package main

import (
	"fmt"
	"sync"
)

//sync.WaitGroup必須傳址(sync.WaitGroup)
//裡面的Done才會對到正確的sync.WaitGroup
func myFunc(waitgroup *sync.WaitGroup) {
	fmt.Println("Inside my goroutine")
	//使用done來通知WaitGroup,goroutine已經結束
	waitgroup.Done()
}
func main() {
	fmt.Println("Hello World")
	//先實例化一個 sync.WaitGroup
	var waitgroup sync.WaitGroup

	//使用add來建立你想要wait的goroutines數量
	//必須在你執行goroutines之前建立數量
	waitgroup.Add(1)

	//必須要傳址
	go myFunc(&waitgroup)

	//阻止main執行,直到前面add的goroutines數量全部執行完成
	waitgroup.Wait()

	fmt.Println("Finished Execution")

}
