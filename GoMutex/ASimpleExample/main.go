//Mutex用來防止同時存取同一筆資料時,造成的資料丟失

package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func init() {
	balance = 1000
}

func deposit(value int, wg *sync.WaitGroup) {
	//先將func鎖住

	//Lock不能使用兩次,要不然會造成Deadlock
	mutex.Lock()
	fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
	balance += value
	//執行結束以後將func解鎖
	//使用了Lock最後一定要用Unlock
	//要不然會造成Deadlock
	mutex.Unlock()
	wg.Done()
	fmt.Printf("New Balance %d\n", balance)
}
func withdraw(value int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("withdrawing %d to account with balance: %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	wg.Done()
	fmt.Printf("New Balance %d\n", balance)
}
func main() {
	fmt.Println("Go Mutex Example")

	var wg sync.WaitGroup
	wg.Add(2)
	go deposit(500, &wg)
	go withdraw(100, &wg)

	wg.Wait()

	fmt.Printf("New Balance %d\n", balance)
}
