package main

import (
	"fmt"
	"go-programs/RLHF/golang_random/28-11-24/389103/turn3/modelB/bankaccount"
)

func main() {
	// Create a new bank account with an initial balance of $100.
	account, err := bankaccount.NewBankAccount("Alice", 100.0)
	if err != nil {
		panic(err) // Handle error appropriately
	}

	// Perform various operations on the account
	err = account.Deposit(50.0)
	if err != nil {
		fmt.Println("Error depositing:", err)
	}

	err = account.Withdraw(30.0)
	if err != nil {
		fmt.Println("Error withdrawing:", err)
	}

	err = account.Withdraw(200.0) // Insufficient balance
	if err != nil {
		fmt.Println("Error withdrawing:", err)
	}

	// Display the account balance
	fmt.Println(account)
}
