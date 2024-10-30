package main

import "fmt"

// BankAccount encapsulates the data and behaviors of a bank account
type BankAccount struct {
	balance float64
}

// Deposit funds into the account
func (a *BankAccount) Deposit(amount float64) {
	a.balance += amount
}

// Withdraw funds from the account
func (a *BankAccount) Withdraw(amount float64) {
	if amount <= a.balance {
		a.balance -= amount
	} else {
		fmt.Println("Insufficient balance")
	}
}

// GetBalance returns the current account balance
func (a *BankAccount) GetBalance() float64 {
	return a.balance
}

func main() {
	// Create a new bank account
	account := &BankAccount{}

	// Deposit funds
	account.Deposit(100.0)

	// Withdraw funds
	account.Withdraw(50.0)

	// Print the current balance
	balance := account.GetBalance()
	fmt.Println("Current Balance:", balance)
}
