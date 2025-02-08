package main

import (
	"fmt"
	"time"
)

func main() {
	const rows = 10000
	const cols = 1000
	var matrix [rows][cols]int
	start := time.Now()
	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = i * j
		}
	}
	elapsed := time.Since(start)
	fmt.Printf("Time taken to calculate matrix: %s\n", elapsed)
}
