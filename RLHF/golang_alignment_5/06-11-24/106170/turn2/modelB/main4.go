package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	// Sample data
	data := []string{"apple", "banana", "cherry", "damson", "elderberry"}

	// Number of partitions
	numPartitions := 3

	// Partition the data based on hash
	partitions := make([][]string, numPartitions)
	for _, value := range data {
		h := fnv.New32a()
		h.Write([]byte(value))
		partitionIndex := int(h.Sum32()) % numPartitions
		partitions[partitionIndex] = append(partitions[partitionIndex], value)
	}

	// Display the partitions
	for i, partition := range partitions {
		fmt.Printf("Partition %d: %v\n", i, partition)
	}
}
