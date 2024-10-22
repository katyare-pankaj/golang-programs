package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dustin/go-humanize"
)

func main() {
	// Initialize random number generator
	rand.Seed(time.Now().UnixNano())

	// Simulate a banking transaction
	balance := 1000.00
	transactionAmount := 50.00

	// Diversify the transaction process by adding a random delay
	delay := time.Duration(rand.Intn(500)) * time.Millisecond
	time.Sleep(delay)

	// Perform the transaction
	if balance >= transactionAmount {
		balance -= transactionAmount
		fmt.Println("Transaction successful! New balance:", humanize.Commaf(balance))
	} else {
		fmt.Println("Insufficient balance. Transaction failed.")
	}
}
