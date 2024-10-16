package main

import (
	"fmt"
)

// Account interface defines the common behavior that all accounts should follow
type Account interface {
	CalculateInterest() float64
	Deposit(amount float64)
	Withdraw(amount float64)
	Balance() float64
}

// SavingsAccount represents a savings account with a fixed interest rate
type SavingsAccount struct {
	balance      float64
	interestRate float64
}

// CalculateInterest implements the Account interface for SavingsAccount
func (s *SavingsAccount) CalculateInterest() float64 {
	return s.balance * s.interestRate
}

func (s *SavingsAccount) Deposit(amount float64) {
	s.balance += amount
}

func (s *SavingsAccount) Withdraw(amount float64) {
	s.balance -= amount
}

func (s *SavingsAccount) Balance() float64 {
	return s.balance
}

// CurrentAccount represents a current account with no interest
type CurrentAccount struct {
	balance float64
}

// CalculateInterest implements the Account interface for CurrentAccount
func (c *CurrentAccount) CalculateInterest() float64 {
	return 0
}

func (c *CurrentAccount) Deposit(amount float64) {
	c.balance += amount
}

func (c *CurrentAccount) Withdraw(amount float64) {
	c.balance -= amount
}

func (c *CurrentAccount) Balance() float64 {
	return c.balance
}

// BusinessAccount represents a business account with a varying interest rate
type BusinessAccount struct {
	balance      float64
	interestRate float64
}

// CalculateInterest implements the Account interface for BusinessAccount
func (b *BusinessAccount) CalculateInterest() float64 {
	return b.balance * b.interestRate
}

func (b *BusinessAccount) Deposit(amount float64) {
	b.balance += amount
}

func (b *BusinessAccount) Withdraw(amount float64) {
	b.balance -= amount
}

func (b *BusinessAccount) Balance() float64 {
	return b.balance
}

// calculateInterestForAccounts takes a slice of Accounts and calculates the total interest earned
func calculateInterestForAccounts(accounts []Account) float64 {
	totalInterest := 0.0
	for _, account := range accounts {
		totalInterest += account.CalculateInterest()
	}
	return totalInterest
}

func main() {
	// Create instances of different account types
	savings := &SavingsAccount{balance: 1000, interestRate: 0.05}
	current := &CurrentAccount{balance: 500}
	business := &BusinessAccount{balance: 2000, interestRate: 0.03}

	// Store accounts in a slice of Accounts using interface values
	accounts := []Account{savings, current, business}

	// Calculate interest for all accounts using the interface method
	totalInterest := calculateInterestForAccounts(accounts)

	fmt.Println("Total interest earned:", totalInterest)
	// Output: Total interest earned: 150
}
