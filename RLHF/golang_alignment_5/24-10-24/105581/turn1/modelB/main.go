// datasource/datasource.go

package main

import (
	"sync"
)

// DataSource represents a real-time data source.
type DataSource interface {
	Subscribe(callback func(data interface{}))
	Unsubscribe(callback func(data interface{}))
	GetData() interface{}
}

// InMemoryDataSource implements the DataSource interface using an in-memory map.
type InMemoryDataSource struct {
	data        map[string]interface{}
	subscribers []func(data interface{})
	mu          sync.Mutex
}

// NewInMemoryDataSource creates a new InMemoryDataSource.
func NewInMemoryDataSource() *InMemoryDataSource {
	return &InMemoryDataSource{
		data:        make(map[string]interface{}),
		subscribers: []func(data interface{}){},
	}
}

// Subscribe implements the DataSource interface.
func (d *InMemoryDataSource) Subscribe(callback func(data interface{})) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.subscribers = append(d.subscribers, callback)
}

// Unsubscribe implements the DataSource interface.
func (d *InMemoryDataSource) Unsubscribe(callback func(data interface{})) {
	d.mu.Lock()
	defer d.mu.Unlock()
	for i, subscriber := range d.subscribers {
		if subscriber == callback {
			d.subscribers = append(d.subscribers[:i], d.subscribers[i+1:]...)
			return
		}
	}
}

// GetData implements the DataSource interface.
func (d *InMemoryDataSource) GetData() interface{} {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.data
}

// UpdateData updates the data in the data source and notifies all subscribers.
func (d *InMemoryDataSource) UpdateData(key string, value interface{}) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.data[key] = value
	for _, subscriber := range d.subscribers {
		subscriber(d.data)
	}
}
