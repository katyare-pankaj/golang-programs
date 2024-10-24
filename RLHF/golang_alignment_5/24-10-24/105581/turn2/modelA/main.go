// main.go

package main

import (
	"time"

	"github.com/example/datastore"
	"github.com/example/processor"
)

func main() {
	// Create a distributed data store
	dataStore := datastore.NewDataStore()

	// Create a processor that will handle data processing
	proc := processor.NewProcessor(dataStore)

	// Start the processor in a goroutine
	go proc.Process()

	// Simulate updates to the data store
	dataStore.Set("user1", 100)
	time.Sleep(1 * time.Second)
	dataStore.Set("user2", 200)
}
