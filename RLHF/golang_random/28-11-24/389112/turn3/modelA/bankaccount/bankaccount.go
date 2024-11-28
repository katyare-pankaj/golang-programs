package bankaccount

import (
	"errors"
	"fmt"
	"time"
)

// BankAccount represents a bank account with a balance, name, and transaction history.
type BankAccount struct {
	name         string
	balance      float64
	transactions []Transaction
}

// NewBankAccount creates a new BankAccount with a given name and initial balance.
func NewBankAccount(name string, initialBalance float64) (*BankAccount, error) {
	if initialBalance < 0 {
		return nil, errors.New("initial balance cannot be negative")
	}
	account := &BankAccount{name: name, balance: initialBalance}
	account.transactions = append(account.transactions, Transaction{
		Date:        time.Now(),
		Type:        "Opening",
		Amount:      initialBalance,
		Description: fmt.Sprintf("Initial deposit for %s", name),
	})
	return account, nil
}

// Deposit deposits an amount into the account and records the transaction.
func (a *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	a.balance += amount
	a.transactions = append(a.transactions, Transaction{
		Date:        time.Now(),
		Type:        "Deposit",
		Amount:      amount,
		Description: fmt.Sprintf("Deposit to %s", a.name),
	})
	return nil
}

// Withdraw withdraws an amount from the account and records the transaction.
func (a *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if a.balance < amount {
		return errors.New("insufficient balance")
	}
	a.balance -= amount
	a.transactions = append(a.transactions, Transaction{
		Date:        time.Now(),
		Type:        "Withdrawal",
		Amount:      amount,
		Description: fmt.Sprintf("Withdrawal from %s", a.name),
	})
	return nil
}

// Balance returns the current balance of the account.
func (a *BankAccount) Balance() float64 {
	return a.balance
}

// Name returns the name of the account holder.
func (a *BankAccount) Name() string {
	return a.name
}

// Transaction represents a transaction made on the bank account.
type Transaction struct {
	Date        time.Time
	Type        string
	Amount      float64
	Description string
}

// TransactionHistory returns the transaction history of the account.
func (a *BankAccount) TransactionHistory() []Transaction {
	return a.transactions
}
