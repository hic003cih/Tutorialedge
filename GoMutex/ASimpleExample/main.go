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
	mutex.Lock()
	fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
	balance += value
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
