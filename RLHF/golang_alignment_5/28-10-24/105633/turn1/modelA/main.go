package main

import (
	"fmt"
	"math"
)

// Calculate the total cost of distributing renewable energy among users using dynamic programming
func distributeRenewableEnergy(users []int, energySources []int) int {
	numUsers := len(users)
	numSources := len(energySources)

	// dp table to store intermediate results
	dp := make([][]int, numUsers+1)
	for i := 0; i < numUsers+1; i++ {
		dp[i] = make([]int, numSources+1)
	}

	// Initialize the base cases
	for i := 0; i <= numUsers; i++ {
		dp[i][0] = math.MaxInt32
	}
	for j := 0; j <= numSources; j++ {
		dp[0][j] = 0
	}

	// Calculate the dynamic programming table
	for i := 1; i <= numUsers; i++ {
		for j := 1; j <= numSources; j++ {
			currentUserEnergy := users[i-1]
			currentSourceEnergy := energySources[j-1]

			if currentUserEnergy <= currentSourceEnergy {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]+currentUserEnergy)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	// Return the minimum cost
	return dp[numUsers][numSources]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	users := []int{2, 3, 4, 3}
	energySources := []int{3, 5, 2, 4}

	fmt.Println("Users:", users)
	fmt.Println("Energy Sources: ", energySources)
	minCost := distributeRenewableEnergy(users, energySources)
	fmt.Println("Minimum Cost to distribute energy:", minCost)
}
