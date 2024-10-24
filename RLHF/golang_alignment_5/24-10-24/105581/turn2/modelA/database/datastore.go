// datastore/datastore.go

package datastore

import (
	"sync"
)

// DataStore represents a distributed data store.
type DataStore struct {
	data   map[string]int
	mu     sync.Mutex
	notify chan struct{}
}

// NewDataStore creates a new DataStore.
func NewDataStore() *DataStore {
	return &DataStore{
		data:   make(map[string]int),
		notify: make(chan struct{}),
	}
}

// Get retrieves data from the store.
func (ds *DataStore) Get(key string) (int, bool) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	value, ok := ds.data[key]
	return value, ok
}

// Set stores data in the store and notifies listeners.
func (ds *DataStore) Set(key string, value int) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.data[key] = value
	close(ds.notify)
	ds.notify = make(chan struct{})
}

// Watch waits for data changes in the store.
func (ds *DataStore) Watch() <-chan struct{} {
	return ds.notify
}
