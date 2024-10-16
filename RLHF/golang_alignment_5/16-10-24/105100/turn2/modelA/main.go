package main

import (
	"fmt"
	"sync"
)

// DistributedLog is a fault-tolerant log structure
type DistributedLog struct {
	mu          sync.Mutex
	replicas    []*Replica
	commitIndex int
}

// Replica represents a single replica of the distributed log
type Replica struct {
	mu        sync.Mutex
	log       []string
	committed bool
}

// NewDistributedLog creates a new distributed log with the specified number of replicas
func NewDistributedLog(numReplicas int) *DistributedLog {
	dl := &DistributedLog{}
	for i := 0; i < numReplicas; i++ {
		dl.replicas = append(dl.replicas, &Replica{})
	}
	return dl
}

// AppendEntries appends entries to the log of all replicas
func (dl *DistributedLog) AppendEntries(entries []string) {
	dl.mu.Lock()
	defer dl.mu.Unlock()

	for _, replica := range dl.replicas {
		replica.mu.Lock()
		defer replica.mu.Unlock()
		replica.log = append(replica.log, entries...)
	}
}

// Commit commits the entries at the specified index to all replicas
func (dl *DistributedLog) Commit(index int) {
	dl.mu.Lock()
	defer dl.mu.Unlock()

	if index <= dl.commitIndex {
		return
	}

	for _, replica := range dl.replicas {
		replica.mu.Lock()
		defer replica.mu.Unlock()
		if len(replica.log) >= index {
			replica.committed = true
		}
	}
	dl.commitIndex = index
}

// GetCommittedLog returns the committed log entries
func (dl *DistributedLog) GetCommittedLog() []string {
	dl.mu.Lock()
	defer dl.mu.Unlock()

	var committedLog []string
	for i := 0; i <= dl.commitIndex; i++ {
		committedLog = append(committedLog, dl.replicas[0].log[i])
	}
	return committedLog
}

func main() {
	// Create a distributed log with 3 replicas
	log := NewDistributedLog(3)

	// Simulate appending entries to the log
	entries := []string{"Entry 1", "Entry 2", "Entry 3", "Entry 4", "Entry 5"}
	log.AppendEntries(entries)

	// Simulate committing entries at index 3
	log.Commit(3)

	// Get the committed log entries
	committedLog := log.GetCommittedLog()
	fmt.Println("Committed Log:")
	for _, entry := range committedLog {
		fmt.Println(entry)
	}
}
