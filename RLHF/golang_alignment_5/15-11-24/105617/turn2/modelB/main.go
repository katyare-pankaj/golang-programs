package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// Transaction represent a transaction or operation for demonstrative using in smart city context
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
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Read order data from the request body
	var order Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Invalid order data.",
		})
		return
	}

	idempotencyKey := r.Header.Get("Idempotency-Key")

	if hasAlreadyProcessedOrder(idempotencyKey) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Order has already been processed.",
		})
		return
	}

	// Simulate processing the order
	newOrder, err := processOrder(idempotencyKey, order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Order processing failed.",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)
}

func processOrder(idempotencyKey string, order Order) (*Order, error) {
	// Simulate order processing
	time.Sleep(1 * time.Second)

	newOrder := &Order{
		ID:        idempotencyKey,
		Status:    "pending",
		Timestamp: time.Now().Format(time.RFC3339),
		Amount:    order.Amount,
	}

	// Log or store the order elsewhere (e.g., database)

	return newOrder, nil
}

var orderProcessedKeys = make(map[string]bool)

func hasAlreadyProcessedOrder(idempotencyKey string) bool {
	_, ok := orderProcessedKeys[idempotencyKey]
	return ok
}
