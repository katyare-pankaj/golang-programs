package main

import (
	"fmt"
	"strings"
)

// Function represents a function in the codebase
type Function struct {
	name    string
	content string
}

// SimilarityMatrix represents a matrix of similarities between functions
type SimilarityMatrix struct {
	functions []*Function
	matrix    [][]int
}

// NewSimilarityMatrix creates a new similarity matrix
func NewSimilarityMatrix(functions []*Function) *SimilarityMatrix {
	matrix := make([][]int, len(functions))
	for i := 0; i < len(functions); i++ {
		matrix[i] = make([]int, len(functions))
	}
	return &SimilarityMatrix{functions: functions, matrix: matrix}
}

// CalculateSimilarities calculates the similarity between all functions in the matrix
func (sm *SimilarityMatrix) CalculateSimilarities() {
	for i := 0; i < len(sm.functions); i++ {
		for j := i + 1; j < len(sm.functions); j++ {
			similarity := calculateSimilarity(sm.functions[i].content, sm.functions[j].content)
			sm.matrix[i][j] = similarity
			sm.matrix[j][i] = similarity
		}
	}
}

// calculateSimilarity calculates the similarity between two functions using LCS
func calculateSimilarity(func1, func2 string) int {
	tokens1 := strings.Fields(func1)
	tokens2 := strings.Fields(func2)

	dp := make([][]int, len(tokens1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(tokens2)+1)
	}

	for i := 1; i <= len(tokens1); i++ {
		for j := 1; j <= len(tokens2); j++ {
			if tokens1[i-1] == tokens2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(tokens1)][len(tokens2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// FindSimilarFunctions finds functions with similarity greater than or equal to the threshold
func (sm *SimilarityMatrix) FindSimilarFunctions(threshold int) [][]string {
	similarFunctions := [][]string{}
	for i := 0; i < len(sm.functions); i++ {
		for j := i + 1; j < len(sm.functions); j++ {
			if sm.matrix[i][j] >= threshold {
				similarFunctions = append(similarFunctions, []string{sm.functions[i].name, sm.functions[j].name})
			}
		}
	}
	return similarFunctions
}

func main() {
	functions := []*Function{
		{name: "add", content: "func add(a, b int) int { return a + b }"},
		{name: "subtract", content: "func subtract(a, b int) int { return a - b }"},
		{name: "multiply", content: "func multiply(a, b int) int { return a * b }"},
		{name: "addOne", content: "func addOne(num int) int { return add(num, 1) }"},
		{name: "double", content: "func double(num int) int { return multiply(num, 2) }"},
		{name: "isEven", content: "func isEven(num int) bool { return num % 2 == 0 }"},
		{name: "findMax", content: "func findMax(a, b int) int { if a > b { return a } else { return b } }"},
		{name: "anotherMax", content: "func anotherMax(a, b int) int { if a > b { return a } else { return b } }"},
	}

	similarityMatrix := NewSimilarityMatrix(functions)
	similarityMatrix.CalculateSimilarities()

	similarFunctions := similarityMatrix.FindSimilarFunctions(3) // Threshold for similarity

	fmt.Println("Similar Functions:")
	for _, pair := range similarFunctions {
		fmt.Println(pair[0], "and", pair[1], "are similar.")
	}
}
