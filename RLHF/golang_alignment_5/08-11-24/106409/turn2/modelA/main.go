package main

import (
	"crypto/rand"
	"fmt"
	"sync"
)

// SecureDataStore represents a secure store for sensitive data
type SecureDataStore struct {
	data map[string][]byte
	mu   sync.RWMutex
}

// NewSecureDataStore creates a new SecureDataStore
func NewSecureDataStore() *SecureDataStore {
	return &SecureDataStore{
		data: make(map[string][]byte),
	}
}

// StoreData securely stores sensitive data in the store
func (s *SecureDataStore) StoreData(key string, data []byte) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = data
}

// RetrieveData securely retrieves sensitive data from the store
func (s *SecureDataStore) RetrieveData(key string) ([]byte, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	data, ok := s.data[key]
	return data, ok
}

// DeleteData securely deletes sensitive data from the store
func (s *SecureDataStore) DeleteData(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}

func generateRandomData(size int) []byte {
	data := make([]byte, size)
	_, err := rand.Read(data)
	if err != nil {
		panic(err)
	}
	return data
}

func handleSensitiveDataConcurrently(store *SecureDataStore, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate concurrent access to the store
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("data-%d", i)
		data := generateRandomData(32)

		// Store data concurrently
		go store.StoreData(key, data)

		// Retrieve data concurrently
		go func(k string) {
			storedData, ok := store.RetrieveData(k)
			if !ok {
				fmt.Printf("Data for key %s not found\n", k)
				return
			}
			if len(storedData) != len(data) || !compareData(storedData, data) {
				fmt.Printf("Data integrity error for key %s\n", k)
			}
		}(key)

		// Delete data concurrently
		go store.DeleteData(key)
	}
}

func compareData(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	store := NewSecureDataStore()
	var wg sync.WaitGroup

	wg.Add(1)
	go handleSensitiveDataConcurrently(store, &wg)

	wg.Wait()
	fmt.Println("Concurrent data handling completed successfully")
}
