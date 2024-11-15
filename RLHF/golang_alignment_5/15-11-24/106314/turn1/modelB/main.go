package main

import (
	"fmt"
	"sync"
	"time"
)

// MockDataSource provides a mocked data source for testing
type MockDataSource struct {
	transactions []string
	mu           sync.Mutex
}

// GetTransactions returns a slice of mock transaction data
func (m *MockDataSource) GetTransactions() []string {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.transactions
}
func processTransactionsWithThrottling(dataSource *MockDataSource, throttleCh <-chan time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case interval := <-throttleCh:
			transactions := dataSource.GetTransactions()
			// Simulate transaction processing and real-time recommendations here.
			fmt.Println("Processing transactions:", transactions)
			time.Sleep(interval)
		default:
			// No throttle interval available, check for exit
			return
		}
	}
}
func main() {
	mockData := &MockDataSource{
		transactions: []string{
			"Transaction 1",
			"Transaction 2", "Transaction 3", "Transaction 4",
		},
	}
	var wg sync.WaitGroup
	wg.Add(1)
	// Example throttle interval: 2 seconds
	throttleInterval := time.Second * 2
	throttleCh := make(chan time.Duration)

	go processTransactionsWithThrottling(mockData, throttleCh, &wg)
	go func() {
		for i := 0; i < 3; i++ {
			throttleCh <- throttleInterval
			time.Sleep(throttleInterval)
		}
		close(throttleCh)
		wg.Wait()
		fmt.Println("Processing completed.")
	}()
}
