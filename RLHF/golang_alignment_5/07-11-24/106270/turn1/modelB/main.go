package main

import (
	"fmt"
)

// Abstracted Account interface
type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64)
	Balance() float64
}

// Concrete struct implementing the Account interface
type SavingsAccount struct {
	balance float64
}

func (a *SavingsAccount) Deposit(amount float64) {
	a.balance += amount
}

func (a *SavingsAccount) Withdraw(amount float64) {
	if amount > a.balance {
		fmt.Println("Insufficient balance")
		return
	}
	a.balance -= amount
}

func (a *SavingsAccount) Balance() float64 {
	return a.balance
}

// Function to perform account transactions
func performTransactions(account Account, deposits []float64, withdrawals []float64) {
	for _, deposit := range deposits {
		account.Deposit(deposit)
	}
	for _, withdrawal := range withdrawals {
		account.Withdraw(withdrawal)
	}
}

func main() {
	savingsAccount := &SavingsAccount{}

	// Perform transactions using the abstracted account interface
	performTransactions(savingsAccount, []float64{100, 50}, []float64{30, 20})

	// Check balance after transactions
	fmt.Println("Balance:", savingsAccount.Balance()) // Output: Balance: 100
}
