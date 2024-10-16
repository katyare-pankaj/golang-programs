package main

import (
	"fmt"
)

// Define an interface for 'Account' with a method 'CalculateInterest()'
type Account interface {
	CalculateInterest() float64
}

// 'SavingsAccount' struct implements the 'Account' interface
type SavingsAccount struct {
	balance      float64
	interestRate float64
}

func (s SavingsAccount) CalculateInterest() float64 {
	// Simple interest calculation for Savings
	return s.balance * s.interestRate * 0.01
}

// 'CurrentAccount' struct implements the 'Account' interface
type CurrentAccount struct {
	balance        float64
	overdraftLimit float64
}

func (c CurrentAccount) CalculateInterest() float64 {
	// No interest for Current Accounts
	return 0
}

// 'BusinessAccount' struct implements the 'Account' interface
type BusinessAccount struct {
	balance      float64
	interestRate float64
	minBalance   float64
}

func (b BusinessAccount) CalculateInterest() float64 {
	// Interest calculation for Business Accounts with min balance check
	if b.balance > b.minBalance {
		return b.balance * b.interestRate * 0.01
	}
	return 0
}

func totalInterestEarned(accounts []Account) float64 {
	var totalInterest float64
	for _, a := range accounts {
		totalInterest += a.CalculateInterest()
	}
	return totalInterest
}

func main() {
	// Create instances of different account types
	savings := SavingsAccount{balance: 1000, interestRate: 5}
	current := CurrentAccount{balance: 200, overdraftLimit: 500}
	business := BusinessAccount{balance: 5000, interestRate: 3, minBalance: 1000}

	accounts := []Account{savings, current, business}

	//Calculate total interest using interface method
	fmt.Println("Total Interest Earned:", totalInterestEarned(accounts))
}
