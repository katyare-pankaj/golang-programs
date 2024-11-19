package main

import (
	"fmt"
)

type PaymentMethod interface {
	ProcessPayment(amount float64) error
}

type CreditCard struct {
	Number string
	CVV    string
	Expiry string
}

func (cc CreditCard) ProcessPayment(amount float64) error {
	// Simulate credit card payment processing
	fmt.Printf("Processing payment of $%.2f using Credit Card %s...\n", amount, cc.Number)
	return nil
}

type PayPal struct {
	Email string
}

func (pp PayPal) ProcessPayment(amount float64) error {
	// Simulate PayPal payment processing
	fmt.Printf("Processing payment of $%.2f using PayPal %s...\n", amount, pp.Email)
	return nil
}

type BankTransfer struct {
	AccountNumber string
	SortCode      string
}

func (bt BankTransfer) ProcessPayment(amount float64) error {
	// Simulate bank transfer payment processing
	fmt.Printf("Processing payment of $%.2f using Bank Transfer %s...\n", amount, bt.AccountNumber)
	return nil
}

func processPayment(pm PaymentMethod, amount float64) error {
	return pm.ProcessPayment(amount)
}

func main() {
	// Define payment methods
	creditCard := CreditCard{Number: "1234-5678-9012-3456", CVV: "123", Expiry: "01/24"}
	payPal := PayPal{Email: "example@example.com"}
	bankTransfer := BankTransfer{AccountNumber: "78901234", SortCode: "56-78-90"}

	// Process payments using different methods
	if err := processPayment(creditCard, 100.00); err != nil {
		fmt.Println("Error processing payment:", err)
	}
	if err := processPayment(payPal, 50.00); err != nil {
		fmt.Println("Error processing payment:", err)
	}
	if err := processPayment(bankTransfer, 75.00); err != nil {
		fmt.Println("Error processing payment:", err)
	}
}
