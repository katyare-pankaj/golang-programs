package main

import (
	"fmt"
)

func main() {
	// Total number of records and number of columns
	numRecords := 10
	numColumns := 10

	// Create a 2D slice to represent the dataset
	dataset := make([][]int, numRecords)
	for i := 0; i < numRecords; i++ {
		dataset[i] = make([]int, numColumns)
		for j := 0; j < numColumns; j++ {
			dataset[i][j] = (i * numColumns) + j
		}
	}

	// Partition the dataset by columns
	numPartitions := 3 // Number of partitions
	partitions := make([][][]int, numPartitions)
	for i := 0; i < numPartitions; i++ {
		startColumn := i * (numColumns / numPartitions)
		endColumn := startColumn + (numColumns / numPartitions) - 1
		for j := 0; j < numRecords; j++ {
			partitions[i] = append(partitions[i], dataset[j][startColumn:endColumn+1])
		}
	}
	// Process each partition independently
	for i, partition := range partitions {
		fmt.Printf("Partition %d:\n", i)
		for _, record := range partition {
			fmt.Println(record)
		}
		fmt.Println()
	}
}
