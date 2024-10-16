package main

import (
	"fmt"
)

// Define an interface for interest calculation
type InterestCalculator interface {
	CalculateInterest() float64
}

type Account struct {
	balance float64
}

func (a *Account) CalculateInterest() float64 {
	// Basic interest calculation for all accounts
	return a.balance * 0.05
}

// NewAccount creates a new account with the specified balance
func NewAccount(balance float64) *Account {
	return &Account{balance: balance}
}

type SavingsAccount struct {
	*Account
}

func (sa *SavingsAccount) CalculateInterest() float64 {
	// Interest calculation for Savings accounts (higher rate)
	return sa.balance * 0.08
}

type CurrentAccount struct {
	*Account
}

func (ca *CurrentAccount) CalculateInterest() float64 {
	// Interest calculation for Current accounts (lower rate)
	return ca.balance * 0.03
}

func processAccounts(accounts []InterestCalculator) float64 {
	totalInterest := 0.0
	for _, account := range accounts {
		totalInterest += account.CalculateInterest()
	}
	return totalInterest
}

func main() {
	savingsInterestCalculator := &SavingsAccount{Account: &Account{balance: 1000.0}}
	currentInterestCalculator := &CurrentAccount{Account: &Account{balance: 2000.0}}
	accounts := []InterestCalculator{savingsInterestCalculator, currentInterestCalculator}

	totalInterest := processAccounts(accounts)
	fmt.Println("Total Interest:", totalInterest)
}
