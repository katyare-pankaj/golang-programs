package main

import (
	"fmt"
	"time"
)

// Metrics is a struct to store metrics data
type Metrics struct {
	Complexity    float64
	ExecutionTime time.Duration
}

// knapsackProblem is a dynamic programming solution for the 0/1 knapsack problem.
// It takes an array of items, each with a weight and a profit, and a knapsack capacity,
// and returns the maximum profit that can be achieved.
func knapsackProblem(weights []int, profits []int, W int) int {
	//Create a 2D array to store intermediate results
	dp := make([][]int, len(weights)+1)
	for i := 0; i < len(weights)+1; i++ {
		dp[i] = make([]int, W+1)
	}

	// Initialize the base case: If the knapsack capacity is 0 or there are no items, return 0
	for i := 0; i < len(weights)+1; i++ {
		dp[i][0] = 0
	}
	for j := 0; j < W+1; j++ {
		dp[0][j] = 0
	}

	// Populate the 2D array using dynamic programming
	for i := 1; i < len(weights)+1; i++ {
		for j := 1; j < W+1; j++ {
			if weights[i-1] <= j {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weights[i-1]]+profits[i-1])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	// Return the maximum profit from the last cell in the 2D array
	return dp[len(weights)][W]
}

// max function returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// calculateMetrics calculates the code complexity and execution time for the knapsackProblem function
func calculateMetrics(weights []int, profits []int, W int) Metrics {
	// start timer
	start := time.Now()

	// calculate knapsack problem
	maxProfit := knapsackProblem(weights, profits, W)

	// stop timer
	end := time.Now()
	executionTime := end.Sub(start)

	// calculate code complexity
	complexity := metrics.CalculateComplexity(knapsackProblem)

	return Metrics{
		Complexity:    complexity,
		ExecutionTime: executionTime,
	}
}

func main() {
	// Sample data for the knapsack problem
	weights := []int{2, 3, 4, 5}
	profits := []int{3, 7, 2, 9}
	W := 6

	// Calculate metrics
	metricsData := calculateMetrics(weights, profits, W)

	// Print the results
	fmt.Println("Maximum profit:", knapsackProblem(weights, profits, W))
	fmt.Println("Code Complexity:", metricsData.Complexity)
	fmt.Println("Execution Time:", metricsData.ExecutionTime)
}
