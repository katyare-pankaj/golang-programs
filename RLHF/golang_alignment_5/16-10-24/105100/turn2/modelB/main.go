package main

import (
	"math/rand"
	"sync"
	"time"
)

// Replica represents a single replica of the data store.
type Replica struct {
	data     map[string]int64
	updating bool
	mu       sync.Mutex
}

// NewReplica creates a new Replica.
func NewReplica() *Replica {
	return &Replica{
		data:     make(map[string]int64),
		updating: false,
	}
}

// AsyncUpdate updates the replica's data asynchronously after a delay.
func (r *Replica) AsyncUpdate(key string, value int64) {
	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // Random delay
		r.mu.Lock()
		defer r.mu.Unlock()
		r.data[key] = value
		r.updating = false
	}()
}

// BestEffortUpdate attempts to update the replica's data best-effort.
func (r *Replica) BestEffortUpdate(key string, value int64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[key] = value
}

// ReadRepair performs read repair by checking if any replicas have newer data.
func (r *Replica) ReadRepair(otherReplicas ...*Replica) {
	for _, replica := range otherReplicas {
		replica.mu.Lock()
		defer replica.mu.Unlock()
		for key, value := range replica.data {
			if r.updating || r.data[key] < value {
				r.data[key] = value
				r.updating = false
			}
		}
	}
}

// GetData retrieves the data from the replica.
func (r *Replica) GetData(key string) int64 {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.data[key]
}

// DistributedDataStore is a fault-tolerant data store that leverages best-effort replication and read repair.
