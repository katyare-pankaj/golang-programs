package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Initialize data
func initData() []int {
	data := make([]int, 100)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(data); i++ {
		data[i] = rand.Intn(100)
	}
	return data
}

// First Sprint: Run the initial risk assessment model
func sprint1(data []int) int {
	// Simple summation-based risk assessment model
	risk := 0
	for _, value := range data {
		risk += value
	}
	return risk
}

// Sprint 2: Define and run a new data mining technique (e.g., mean)
func sprint2(data []int) int {
	// Calculate the mean as a new data mining technique
	total := 0
	for _, value := range data {
		total += value
	}
	mean := total / len(data)
	return mean
}

// Sprint 3: Iterate on the previous sprint's result.
func sprint3(data []int, previousRisk int) int {
	// Apply a threshold to the previous sprint's result to identify high-risk situations.
	highRiskThreshold := 50
	var newRisk int
	if previousRisk > highRiskThreshold {
		newRisk = 100
	} else {
		newRisk = 0
	}
	return newRisk
}
func main() {
	data := initData()
	// Sprint 1: Run the initial risk assessment model
	riskSprint1 := sprint1(data)
	fmt.Println("Sprint 1 Risk:", riskSprint1)

	// Sprint 2: Define and run a new data mining technique
	riskSprint2 := sprint2(data)
	fmt.Println("Sprint 2 Risk:", riskSprint2)

	// Sprint 3: Iterate on the previous sprint's result
	riskSprint3 := sprint3(data, riskSprint2)
	fmt.Println("Sprint 3 Risk:", riskSprint3)
}
