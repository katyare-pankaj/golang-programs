package bankaccount

import "errors"

// BankAccount represents a bank account with a balance and a name.
type BankAccount struct {
	balance float64
	name    string
}

// NewBankAccount creates a new BankAccount with a given name and initial balance.
func NewBankAccount(name string, initialBalance float64) (*BankAccount, error) {
	if initialBalance < 0 {
		return nil, errors.New("initial balance cannot be negative")
	}
	return &BankAccount{name: name, balance: initialBalance}, nil
}

// Deposit deposits an amount into the account.
func (a *BankAccount) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be positive")
	}
	a.balance += amount
	return nil
}

// Withdraw withdraws an amount from the account.
func (a *BankAccount) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be positive")
	}
	if a.balance < amount {
		return errors.New("insufficient balance")
	}
	a.balance -= amount
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
