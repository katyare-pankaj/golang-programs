// datastore_test.go

package datastore

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDataStore(t *testing.T) {
	ds := NewDataStore()

	// Test Set and Get
	ds.Set("key", 1)
	value, ok := ds.Get("key")
	assert.True(t, ok)
	assert.Equal(t, 1, value)

	// Test Watch
	done := make(chan struct{})
	go func() {
		ds.Watch()
		close(done)
	}()
	ds.Set("anotherKey", 2)
	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("Watch did not notify within the timeout")
	}
}
