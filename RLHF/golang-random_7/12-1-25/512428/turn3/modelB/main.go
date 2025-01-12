package main

import (
	"fmt"
)

func generateIDStrings(start int, end int, batchSize int) [][]string {
	var results [][]string

	for i := start; i <= end; i += batchSize {
		endBatch := i + batchSize - 1
		if endBatch > end {
			endBatch = end
		}
		var batch []string
		for j := i; j <= endBatch; j++ {
			str := fmt.Sprintf("ID%06d", j)
			batch = append(batch, str)
		}
		results = append(results, batch)
	}
	return results
}

func main() {
	start := 1
	end := 100
	batchSize := 10

	batches := generateIDStrings(start, end, batchSize)

	for i, batch := range batches {
		fmt.Println("Batch", i+1, ":")
		for _, str := range batch {
			fmt.Print(str, " ")
		}
		fmt.Println()
	}
}
