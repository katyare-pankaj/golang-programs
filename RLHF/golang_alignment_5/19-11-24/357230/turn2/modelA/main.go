package main

import (
	"fmt"
)

// Define the Payment interface
type Payment interface {
	Charge(amount float64) error
	String() string
}

// Implement the CreditCard struct that conforms to the Payment interface
type CreditCard struct {
	number string
	expiry string
	cvv    string
}

func (cc CreditCard) Charge(amount float64) error {
	fmt.Printf("Charging CreditCard %s for $%.2f...\n", cc.number, amount)
	return nil
}

func (cc CreditCard) String() string {
	return fmt.Sprintf("CreditCard {number: %s, expiry: %s, cvv: %s}", cc.number, cc.expiry, cc.cvv)
}

// Implement the PayPal struct that conforms to the Payment interface
type PayPal struct {
	email string
}

func (pp PayPal) Charge(amount float64) error {
	fmt.Printf("Charging PayPal %s for $%.2f...\n", pp.email, amount)
	return nil
}

func (pp PayPal) String() string {
	return fmt.Sprintf("PayPal {email: %s}", pp.email)
}

// Implement the BankTransfer struct that conforms to the Payment interface
type BankTransfer struct {
	accountNumber string
	routingNumber string
}

func (bt BankTransfer) Charge(amount float64) error {
	fmt.Printf("Processing BankTransfer %s for $%.2f...\n", bt.accountNumber, amount)
	return nil
}

func (bt BankTransfer) String() string {
	return fmt.Sprintf("BankTransfer {accountNumber: %s, routingNumber: %s}", bt.accountNumber, bt.routingNumber)
}

// Function to process payments for a list of payment methods
func processPayments(payments []Payment) {
	for _, payment := range payments {
		fmt.Println("Payment Method:", payment)
		if err := payment.Charge(100.0); err != nil {
			fmt.Println("Error charging payment:", err)
		} else {
			fmt.Println("Payment successful.")
		}
	}
}

func main() {
	// Create instances of different payment methods
	creditCard := CreditCard{number: "4111111111111111", expiry: "12/25", cvv: "123"}
	paypal := PayPal{email: "user@example.com"}
	bankTransfer := BankTransfer{accountNumber: "123456789", routingNumber: "098765432"}

	// Create a slice of Payment to hold different payment methods
	paymentMethods := []Payment{creditCard, paypal, bankTransfer}

	// Process payments
	processPayments(paymentMethods)
}
