package main

import (
	"fmt"
	"go-programs/RLHF/golang_random/28-11-24/389103/turn2/modelA/bank"
)

func main() {
	// Create two bank accounts
	account1 := bank.NewAccount(1000) // Initialize with a balance of 1000
	account2 := bank.NewAccount(500)  // Initialize with a balance of 500

	fmt.Printf("Initial Balance of Account 1: %.2f\n", account1.Balance())
	fmt.Printf("Initial Balance of Account 2: %.2f\n", account2.Balance())

	// Deposit money into account1
	account1.Deposit(200)
	fmt.Printf("Balance of Account 1 after deposit: %.2f\n", account1.Balance())

	// Withdraw money from account2
	account2.Withdraw(100)
	fmt.Printf("Balance of Account 2 after withdrawal: %.2f\n", account2.Balance())

	// Transfer money from account1 to account2
	account1.Transfer(300, account2)
	fmt.Printf("Balance of Account 1 after transfer: %.2f\n", account1.Balance())
	fmt.Printf("Balance of Account 2 after transfer: %.2f\n", account2.Balance())

	// Attempt to withdraw an amount greater than the available balance in account1
	account1.Withdraw(1500)
	fmt.Printf("Balance of Account 1 after failed withdrawal attempt: %.2f\n", account1.Balance())
}
