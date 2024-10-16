package main

import (
	"fmt"
	"time"
)

// TCAAnalyzer represents the main class for TCA analysis
type TCAAnalyzer struct {
	transactionRecords []Transaction
	// Other state variables as needed
}

// Transaction represents a single transaction record
type Transaction struct {
	timestamp time.Time
	cost      float64
}

func main() {
	// Create a new TCA analyzer
	analyzer := NewTCAAnalyzer()

	// Process new transactions in an incremental manner
	analyzer.ProcessTransaction(Transaction{time.Now(), 10.0})
	analyzer.ProcessTransaction(Transaction{time.Now(), 20.0})
	// ... and so on

	// Generate reports and analysis
	fmt.Println("Total Transaction Cost:", analyzer.CalculateTotalCost())
	fmt.Println("Average Transaction Cost:", analyzer.CalculateAverageCost())
}

// NewTCAAnalyzer creates a new instance of TCAAnalyzer
func NewTCAAnalyzer() *TCAAnalyzer {
	return &TCAAnalyzer{}
}

// ProcessTransaction adds a new transaction to the analyzer
func (a *TCAAnalyzer) ProcessTransaction(transaction Transaction) {
	a.transactionRecords = append(a.transactionRecords, transaction)
}

// CalculateTotalCost sums up the costs of all transactions
func (a *TCAAnalyzer) CalculateTotalCost() float64 {
	totalCost := 0.0
	for _, transaction := range a.transactionRecords {
		totalCost += transaction.cost
	}
	return totalCost
}

// CalculateAverageCost calculates the average cost of transactions
func (a *TCAAnalyzer) CalculateAverageCost() float64 {
	totalCost := a.CalculateTotalCost()
	if len(a.transactionRecords) == 0 {
		return 0.0
	}
	return totalCost / float64(len(a.transactionRecords))
}
