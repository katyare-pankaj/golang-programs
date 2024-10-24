// keyvaluestore/keyvaluestore.go

package keyvaluestore

import (
	"sync"
)

// KeyValueStore represents a distributed key-value store.
type KeyValueStore struct {
	data map[string]string
	mu   sync.Mutex
}

// NewKeyValueStore creates a new KeyValueStore.
func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		data: make(map[string]string),
	}
}

// Get retrieves the value associated with the given key from the store.
func (k *KeyValueStore) Get(key string) (string, bool) {
	k.mu.Lock()
	defer k.mu.Unlock()
	value, ok := k.data[key]
	return value, ok
}

// Set updates the value associated with the given key in the store.
func (k *KeyValueStore) Set(key string, value string) {
	k.mu.Lock()
	defer k.mu.Unlock()
	k.data[key] = value
}

// Delete removes the key-value pair from the store.
func (k *KeyValueStore) Delete(key string) {
	k.mu.Lock()
	defer k.mu.Unlock()
	delete(k.data, key)
}
