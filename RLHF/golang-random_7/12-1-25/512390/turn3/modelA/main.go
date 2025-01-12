package main

import (
	"fmt"
)

func generateFormattedIDs(format string, start int, end int, batchSize int) {
	for i := start; i <= end; i += batchSize {
		batch := make([]string, 0)
		for j := 0; j < batchSize && i+j <= end; j++ {
			batch = append(batch, fmt.Sprintf(format, i+j))
		}
		fmt.Println("Batch", (i/batchSize)+1, ":", batch)
	}
}

func main() {
	format := "ID%05d"
	start := 1
	end := 100
	batchSize := 10

	generateFormattedIDs(format, start, end, batchSize)
}
