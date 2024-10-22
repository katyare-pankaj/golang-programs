package main

import (
	"fmt"
	"sync"
	"time"
)

// CRMData represents a single entry in the CRM system.
type CRMData struct {
	ID        int
	Name      string
	Amount    float64
	CreatedAt time.Time
}

// AggregatedData represents the aggregated results for a particular partition.
type AggregatedData struct {
	TotalAmount float64
	Count       int
}

// partitionData divides the CRM data into multiple partitions based on a specified column (in this case, 'CreatedAt' year).
func partitionData(crmData []CRMData) map[int][]CRMData {
	partitions := make(map[int][]CRMData)

	for _, data := range crmData {
		year := data.CreatedAt.Year()
		partitions[year] = append(partitions[year], data)
	}
	return partitions
}

// aggregateData aggregates the data for a single partition.
func aggregateData(partition []CRMData) AggregatedData {
	var aggData AggregatedData
	for _, data := range partition {
		aggData.TotalAmount += data.Amount
		aggData.Count++
	}
	return aggData
}

// concurrentAggregateData uses parallel processing to aggregate data from multiple partitions concurrently.
func concurrentAggregateData(crmData []CRMData) map[int]AggregatedData {
	partitions := partitionData(crmData)
	results := make(map[int]AggregatedData)

	var wg sync.WaitGroup
	for year, partition := range partitions {
		wg.Add(1)
		go func(year int, partition []CRMData) {
			defer wg.Done()
			results[year] = aggregateData(partition)
		}(year, partition)
	}

	wg.Wait()
	return results
}

func main() {
	// Generate sample CRM data for demonstration purposes
	var crmData []CRMData
	for i := 0; i < 10000; i++ {
		crmData = append(crmData, CRMData{
			ID:        i,
			Name:      fmt.Sprintf("Customer %d", i),
			Amount:    float64(i) * 10.0,
			CreatedAt: time.Now().AddDate(0, 0, -i),
		})
	}

	startTime := time.Now()
	// Aggregate data concurrently
	aggregatedResults := concurrentAggregateData(crmData)
	duration := time.Since(startTime)

	fmt.Println("Concurrent Aggregation Results:")
	for year, result := range aggregatedResults {
		fmt.Printf("Year: %d, Total Amount: %.2f, Count: %d\n", year, result.TotalAmount, result.Count)
	}
	fmt.Printf("Aggregation completed in: %s\n", duration)
}
