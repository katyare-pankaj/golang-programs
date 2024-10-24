// processor/processor.go

package processor

import (
	"fmt"
	"sync"

	"github.com/example/datastore"
)

// Processor handles data processing and synchronization.
type Processor struct {
	datastore *datastore.DataStore
	mu        sync.Mutex
}

// NewProcessor creates a new Processor.
func NewProcessor(datastore *datastore.DataStore) *Processor {
	return &Processor{
		datastore: datastore,
	}
}

// Process starts data processing and synchronization.
func (p *Processor) Process() {
	for {
		select {
		case <-p.datastore.Watch():
			p.mu.Lock()
			// Perform processing here
			fmt.Println("Data updated in the datastore. Performing processing...")
			p.mu.Unlock()
		}
	}
}

// UpdateData updates data in the datastore
func (p *Processor) UpdateData(key string, value int) {
	p.datastore.Set(key, value)
}
