// processor_test.go

package processor

import (
	"sync"
	"testing"
	"time"

	"github.com/example/datastore"
)

func TestProcessor(t *testing.T) {
	ds := datastore.NewDataStore()
	p := NewProcessor(ds)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		p.Process()
		wg.Done()
	}()

	// Test data processing
	ds.Set("user1", 100)
	time.Sleep(time.Second) // Allow for processing

	wg.Wait()
}
