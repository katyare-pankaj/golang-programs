package main

import (
	"fmt"
	"time"
)

// Recursive function to calculate the length of the Longest Common Subsequence (LCS) of two strings
func lcsLengthRecursive(text1 string, text2 string, m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if text1[m-1] == text2[n-1] {
		return 1 + lcsLengthRecursive(text1, text2, m-1, n-1)
	}
	return max(lcsLengthRecursive(text1, text2, m-1, n), lcsLengthRecursive(text1, text2, m, n-1))
}

// Dynamic Programming approach to optimize the LCS length calculation
func lcsLengthDP(text1 string, text2 string, m int, n int) int {
	// Create a 2D array to store pre-computed LCS lengths
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// Initialize the base cases
	for i := 0; i <= m; i++ {
		dp[i][0] = 0
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = 0
	}

	// Calculate LCS lengths for each character of both strings
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	text1 := "AGGTAB"
	text2 := "GXTXAYB"
	m := len(text1)
	n := len(text2)

	fmt.Println("Recursive Approach:")
	recursiveTime := measureTime(func() {
		lcsLength := lcsLengthRecursive(text1, text2, m, n)
		fmt.Println("Length of LCS is:", lcsLength)
	})
	fmt.Println("Recursive time:", recursiveTime)

	fmt.Println("\nDynamic Programming Approach:")
	dpTime := measureTime(func() {
		lcsLength := lcsLengthDP(text1, text2, m, n)
		fmt.Println("Length of LCS is:", lcsLength)
	})
	fmt.Println("Dynamic Programming time:", dpTime)
}

func measureTime(block func()) time.Duration {
	start := time.Now()
	block()
	return time.Since(start)
}
