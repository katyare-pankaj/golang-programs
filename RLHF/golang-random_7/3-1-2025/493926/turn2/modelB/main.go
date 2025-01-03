package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	// Mock database package
	"github.com/golang/mock/gomock"
)

// Mock database interface
type MockDB interface {
	GetData() (string, error)
	SetData(string) error
}

// Mock database implementation
type MockDBImpl struct {
	ctrl   *gomock.Controller
	mu     sync.Mutex // Mutex to protect data
	data   string
	errors error
}

func (m *MockDBImpl) GetData() (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	time.Sleep(time.Second) // Simulate some I/O delay
	return m.data, m.errors
}

func (m *MockDBImpl) SetData(data string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data = data
	return nil
}

func NewMockDB(ctrl *gomock.Controller) MockDB {
	return &MockDBImpl{ctrl: ctrl}
}

func main() {
	ctrl := gomock.NewController(GinkgoT())
	defer ctrl.Finish()

	db := NewMockDB(ctrl)

	var wg sync.WaitGroup
	numOperations := 10

	// Start reading data from the database concurrently
	wg.Add(numOperations)
	for i := 0; i < numOperations; i++ {
		go func() {
			defer wg.Done()
			data, err := db.GetData()
			if err != nil {
				log.Printf("Error reading data: %v", err)
				return
			}
			fmt.Println("Read data:", data)
		}()
	}

	// Start writing data to the database concurrently
	wg.Add(numOperations)
	for i := 0; i < numOperations; i++ {
		go func() {
			defer wg.Done()
			// Simulate writing data to the database
			data := fmt.Sprintf("Sample Data %d", i)
			err := db.SetData(data)
			if err != nil {
				log.Printf("Error writing data: %v", err)
				return
			}
			fmt.Println("Wrote data to database:", data)
		}()
	}

	wg.Wait()
	fmt.Println("All tasks completed.")
}
