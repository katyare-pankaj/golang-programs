package main

import (
	"fmt"
	"sync"
	"time"
)

type Bank struct {
	balance int
	mutex   sync.Mutex
}

func (b *Bank) Deposit(amount int) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.balance += amount
}

func (b *Bank) Withdraw(amount int) bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if b.balance >= amount {
		b.balance -= amount
		return true
	}
	return false
}

func (b *Bank) Balance() int {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.balance
}

func transferMoney(from *Bank, to *Bank, amount int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Use a loop to ensure transfer succeeds eventually
	for {
		if from.Withdraw(amount) {
			to.Deposit(amount)
			fmt.Printf("Transferring %d from %s to %s. New balances: %s = %d, %s = %d\n",
				amount, from.balance, to.balance)
			return
		}
		// Simulate some time for other goroutines to make progress
		time.Sleep(time.Millisecond * 10)
	}
}

func main() {
	var wg sync.WaitGroup

	bank1 := &Bank{balance: 100}
	bank2 := &Bank{balance: 50}

	wg.Add(2)
	go transferMoney(bank1, bank2, 20, &wg)
	go transferMoney(bank2, bank1, 30, &wg)

	wg.Wait()

	fmt.Println("Final balances:")
	fmt.Printf("Bank1: %d\n", bank1.Balance())
	fmt.Printf("Bank2: %d\n", bank2.Balance())
}
