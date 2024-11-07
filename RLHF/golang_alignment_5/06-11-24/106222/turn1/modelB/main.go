package main

import (
	"fmt"

	"github.com/gonum/graph/algorithm/components"
	"github.com/gonum/graph/simple"
)

// Sample functions with some code duplication
func func1() {
	fmt.Println("Code snippet 1")
	fmt.Println("Similar code 1")
}

func func2() {
	fmt.Println("Code snippet 2")
	fmt.Println("Similar code 1")
}

func func3() {
	fmt.Println("Different code")
}

func func4() {
	fmt.Println("Similar code 1")
	fmt.Println("Code snippet 4")
}

// Calculate similarity between two functions (Example: Hamming distance)
func calculateSimilarity(code1, code2 []string) float64 {
	// Hamming distance for simplicity
	maxLen := max(len(code1), len(code2))
	diff := 0
	for i := 0; i < maxLen; i++ {
		if i >= len(code1) || i >= len(code2) || code1[i] != code2[i] {
			diff++
		}
	}
	return 1.0 - float64(diff)/float64(maxLen)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	// Create the code graph
	g := simple.NewDirectedGraph()
	functions := []string{"func1", "func2", "func3", "func4"}
	// Define similarity threshold
	similarityThreshold := 0.7
	for _, f1 := range functions {
		for _, f2 := range functions {
			if f1 != f2 {
				// get function code strings from some source
				// For simplicity, we'll use hardcoded code.
				code1 := []string{"Line1", "Similar", "Line3"}
				code2 := []string{"LineA", "Similar", "LineC"}
				// Calculate similarity
				similarity := calculateSimilarity(code1, code2)

				if similarity >= similarityThreshold {
					g.SetEdge(g.NewEdge(f1, f2))
				}
			}
		}
	}
	// Find strongly connected components (SCCs)
	sccs := components.StronglyConnected(g)
	for _, scc := range sccs {
		fmt.Println("Strongly Connected Component:")
		for _, f := range scc {
			fmt.Printf(" -> %s\n", f.ID())
		
