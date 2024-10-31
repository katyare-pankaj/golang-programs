package gateway

import (
	"errors"
	"fmt"
)

// PaymentGateway represents a payment gateway interface.
type PaymentGateway interface {
	ProcessPayment(amount float64) error
}

// CreditCardGateway implements the PaymentGateway interface.
type CreditCardGateway struct {
	apiKey string
}

// NewCreditCardGateway creates a new CreditCardGateway instance.
func NewCreditCardGateway(apiKey string) *CreditCardGateway {
	return &CreditCardGateway{apiKey: apiKey}
}

// ProcessPayment processes a payment using the credit card gateway.
func (g *CreditCardGateway) ProcessPayment(amount float64) error {
	// Simulate payment processing logic
	fmt.Printf("Processing payment of $%.2f using credit card\n", amount)
	if amount > 1000 {
		return errors.New("payment amount exceeds limit")
	}
	return nil
}
