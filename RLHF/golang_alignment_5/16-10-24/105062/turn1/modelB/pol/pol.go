package main

import (
	"fmt"
)

type Account struct {
	balance float64
}

func (a *Account) CalculateInterest() float64 {
	// Basic interest calculation for all accounts
	return a.balance * 0.05
}

func processAccounts(accounts []Account) float64 {
	totalInterest := 0.0
	for _, account := range accounts {
		totalInterest += account.CalculateInterest()
	}
	return totalInterest
}

func main() {
	savingsAccount := Account{balance: 1000.0}
	currentAccount := Account{balance: 2000.0}
	accounts := []Account{savingsAccount, currentAccount}

	totalInterest := processAccounts(accounts)
	fmt.Println("Total Interest:", totalInterest)
}
