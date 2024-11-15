package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Order represents an e-commerce order.
type Order struct {
	ID        string  `json:"id"`
	Status    string  `json:"status"`
	Timestamp string  `json:"timestamp"`
	Amount    float64 `json:"amount"`
}

func main() {
	http.HandleFunc("/orders", handleOrder)
	http.ListenAndServe(":8080", nil)
}

func handleOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idempotencyKey := r.Header.Get("Idempotency-Key")

	if idempotencyKey == "" {
		http.Error(w, "Idempotency-Key header is required", http.StatusBadRequest)
		return
	}

	if hasAlreadyProcessed(idempotencyKey) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Order already processed successfully.",
		})
		return
	}

	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Simulate order processing
	newOrder, err := processOrder(idempotencyKey, order)
	if err != nil {
		http.Error(w, "Order processing failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)
}

func processOrder(idempotencyKey string, order Order) (*Order, error) {
	time.Sleep(2 * time.Second)

	newOrder := &Order{
		ID:        idempotencyKey,
		Status:    "pending",
		Timestamp: time.Now().Format(time.RFC3339),
		Amount:    order.Amount,
	}

	storeOrder(newOrder)
	return newOrder, nil
}

func storeOrder(order *Order) {
	fmt.Println("Storing order:", order)
	// In a real application, you would store the order in a database.
}

var processedOrders = make(map[string]bool)

func hasAlreadyProcessed(idempotencyKey string) bool {
	_, ok := processedOrders[idempotencyKey]
	return ok
}
