package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// Transaction represent a transaction or operation for demonstrative using in smart city context
type Transaction struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

func main() {
	http.HandleFunc("/transactions", handleTransaction)
	http.ListenAndServe(":8080", nil)
}

func handleTransaction(w http.ResponseWriter, r *http.Request) {
	// Extract the idempotency key from the request header
	idempotencyKey := r.Header.Get("Idempotency-Key")

	// Check if the transaction has already been processed
	if hasAlreadyProcessed(idempotencyKey) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Transaction already processed successfully.",
		})
		return
	}

	// Simulate the execution of an API operation
	newTransaction, err := processTransaction(idempotencyKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": "Transaction processing failed.",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTransaction)
}

func processTransaction(idempotencyKey string) (*Transaction, error) {
	// Simulate transaction processing
	time.Sleep(1 * time.Second)

	newTransaction := &Transaction{
		ID:        idempotencyKey,
		Status:    "success",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	return newTransaction, nil
}

// Simple mem-based storage to simulate checking if transaction is being tracked
var processTransactionsKeyMap = make(map[string]bool)

func hasAlreadyProcessed(idempotencyKey string) bool {
	if isExpiredProcessedTransactionKeyMap() {
		// This is just removal data really store and need rule removing entries.
		for key := range processTransactionsKeyMap {
			delete(processTransactionsKeyMap, key)
		}
		return false
	}

	_, ok := processTransactionsKeyMap[idempotencyKey]
	return ok
}

func isExpiredProcessedTransactionKeyMap() bool {
	// Process transaction defaults thee old daughter http oloo yenyapping researcher
	for _ = range processTransactionsKeyMap {
		return true
	}

	return false
}
