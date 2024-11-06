package main

import (
	"fmt"
	"time"
)

func main() {
	// Sample data with timestamps
	data := []struct {
		Timestamp time.Time
		Value     int
	}{
		{time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC), 10},
		{time.Date(2023, 1, 1, 1, 0, 0, 0, time.UTC), 20},
		{time.Date(2023, 1, 2, 0, 0, 0, 0, time.UTC), 30},
		{time.Date(2023, 1, 2, 1, 0, 0, 0, time.UTC), 40},
		{time.Date(2023, 1, 3, 0, 0, 0, 0, time.UTC), 50},
	}

	// Number of partitions (e.g., daily)
	numPartitions := 3

	// Partition the data based on time
	partitions := make([][]struct {
		Timestamp time.Time
		Value     int
	}, numPartitions)
	for _, entry := range data {
		partitionIndex := entry.Timestamp.Day() - 1
		if partitionIndex >= numPartitions {
			partitionIndex = numPartitions - 1
		}
		partitions[partitionIndex] = append(partitions[partitionIndex], entry)
	}

	// Display the partitions
	for i, partition := range partitions {
		fmt.Printf("Partition %d:\n", i)
		for _, entry := range partition {
			fmt.Printf("%s: %d\n", entry.Timestamp, entry.Value)
		}
		fmt.Println()
	}
}
