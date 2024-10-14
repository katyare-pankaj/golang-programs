package main

import (
	"net/http"
	"sync"
	"time"
)

// ConnectionPool represents a pool of HTTP connections
type ConnectionPool struct {
	mu      sync.Mutex
	idle    []*http.Transport
	active  []*http.Transport
	maxIdle int
}

// NewConnectionPool creates a new ConnectionPool
func NewConnectionPool(maxIdle int) *ConnectionPool {
	return &ConnectionPool{maxIdle: maxIdle}
}

// Get retrieves an idle connection from the pool or creates a new one if necessary
func (p *ConnectionPool) Get() *http.Transport {
	p.mu.Lock()
	defer p.mu.Unlock()

	for len(p.idle) > 0 {
		transport := p.idle[0]
		p.idle = p.idle[1:]
		p.active = append(p.active, transport)
		return transport
	}

	// Create a new connection if the pool is empty
	transport := &http.Transport{}
	p.active = append(p.active, transport)
	return transport
}

// Put returns an active connection to the idle pool
func (p *ConnectionPool) Put(transport *http.Transport) {
	p.mu.Lock()
	defer p.mu.Unlock()

	// Remove the transport from the active list
	for i, t := range p.active {
		if t == transport {
			p.active = append(p.active[:i], p.active[i+1:]...)
			break
		}
	}

	// Add the transport to the idle list if the limit is not reached
	if len(p.idle) < p.maxIdle {
		p.idle = append(p.idle, transport)
	}
}

func main() {
	pool := NewConnectionPool(10) // Create a connection pool with a maximum of 10 idle connections
	// Use the connection pool in your API gateway logic
	// For demonstration, we'll just get and put connections repeatedly
	for i := 0; i < 20; i++ {
		transport := pool.Get()
		// Use the transport for your API request
		// ...

		// Return the connection to the pool after the request is handled
		pool.Put(transport)
		time.Sleep(time.Second) // Simulate some work
	}
}
