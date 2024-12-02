package bank

// Account struct represents a bank account
type Account struct {
	balance float64
}

// NewAccount creates a new bank account with the given balance
func NewAccount(balance float64) *Account {
	if balance < 0 {
		balance = 0
	}
	return &Account{balance: balance}
}

// Deposit adds a specified amount to the account balance
func (a *Account) Deposit(amount float64) {
	if amount > 0 {
		a.balance += amount
	}
}

// Withdraw removes a specified amount from the account balance if it's available
func (a *Account) Withdraw(amount float64) {
	if amount > 0 && a.balance >= amount {
		a.balance -= amount
	}
}

// Balance returns the current balance of the account
func (a *Account) Balance() float64 {
	return a.balance
}

// Transfer transfers a specified amount from the current account to the target account if it's available
func (a *Account) Transfer(amount float64, target *Account) {
	a.Withdraw(amount)
	target.Deposit(amount)
}
