// payments/customer/module.go

package customer

import (
	"fmt"
	"go-programs/RLHF/golang_alignment_5/30-10-24/105724/turn2/modelB/gateway"
)

// Customer represents a customer with a payment method.
type Customer struct {
	Name          string
	PaymentMethod gateway.PaymentGateway
}

// NewCustomer creates a new customer instance.
func NewCustomer(name string, paymentMethod gateway.PaymentGateway) *Customer {
	return &Customer{Name: name, PaymentMethod: paymentMethod}
}

// ProcessPayment processes a payment for the customer.
func (c *Customer) ProcessPayment(amount float64) error {
	fmt.Printf("Processing payment for %s\n", c.Name)
	return c.PaymentMethod.ProcessPayment(amount)
}
