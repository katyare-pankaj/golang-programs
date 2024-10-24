// replicator/replicator.go

package replicator

import (
	"sync"

	"github.com/example/keyvaluestore"
)

// Replicator handles the replication of data between multiple nodes in a distributed system.
type Replicator struct {
	localStore   *keyvaluestore.KeyValueStore
	remoteStores []*keyvaluestore.KeyValueStore
	mu           sync.Mutex
}

// NewReplicator creates a new Replicator.
func NewReplicator(localStore *keyvaluestore.KeyValueStore, remoteStores ...*keyvaluestore.KeyValueStore) *Replicator {
	return &Replicator{
		localStore:   localStore,
		remoteStores: remoteStores,
	}
}

// Replicate updates the value of a key in all remote stores with the value from the local store.
func (r *Replicator) Replicate(key string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	localValue, ok := r.localStore.Get(key)
	if !ok {
		return
	}

	for _, remoteStore := range r.remoteStores {
		remoteStore.Set(key, localValue)
	}
}

// UpdateLocalStore updates the value of a key in the local store.
func (r *Replicator) UpdateLocalStore(key string, value string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.localStore.Set(key, value)
	r.Replicate(key)
}
