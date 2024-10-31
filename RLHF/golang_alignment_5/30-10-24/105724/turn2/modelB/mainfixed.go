package main

import (
	"go-programs/RLHF/golang_alignment_5/30-10-24/105724/turn2/modelB/customer"
	"go-programs/RLHF/golang_alignment_5/30-10-24/105724/turn2/modelB/gateway"
	"log"
)

func main() {
	// Create a new CreditCardGateway instance with a dummy API key
	ccGateway := gateway.NewCreditCardGateway("dummy-api-key")

	cust := customer.NewCustomer("John Doe", ccGateway)

	// Process a payment
	amount := 150.0
	err := cust.ProcessPayment(amount)
	if err != nil {
		log.Fatalf("Failed to process payment: %v", err)
	}

}
