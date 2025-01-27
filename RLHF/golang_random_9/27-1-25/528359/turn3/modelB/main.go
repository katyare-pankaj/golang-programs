package main

import (
	"errors"
	"fmt"
	"net"
	"sync"
)

// ConnectionPool manages a pool of reusable connections.
type ConnectionPool struct {
	mu       sync.Mutex
	cond     *sync.Cond
	conns    []net.Conn
	capacity int
	closed   bool
}

// NewConnectionPool creates a new pool with the specified capacity.
func NewConnectionPool(capacity int) *ConnectionPool {
	p := &ConnectionPool{
		capacity: capacity,
		conns:    make([]net.Conn, 0, capacity),
	}
	p.cond = sync.NewCond(&p.mu)
	return p
}

// Get retrieves a connection from the pool, waiting if necessary.
func (p *ConnectionPool) Get() (net.Conn, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for len(p.conns) == 0 && !p.closed {
		p.cond.Wait() // Wait for a connection to be available or pool to close.
	}

	if p.closed {
		return nil, errors.New("pool is closed")
	}

	// Take a connection from the slice (LIFO strategy)
	conn := p.conns[len(p.conns)-1]
	p.conns = p.conns[:len(p.conns)-1]

	return conn, nil
}

// Put returns a connection to the pool.
func (p *ConnectionPool) Put(conn net.Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.conns) < p.capacity && !p.closed {
		p.conns = append(p.conns, conn)
		p.cond.Signal() // Notify one waiting goroutine that a connection is available.
	} else {
		conn.Close() // Can't add more or pool is closed, close the connection
	}
}

// Close closes the pool and cleans up all connections.
func (p *ConnectionPool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.closed = true

	// Close all remaining connections in the pool
	for _, conn := range p.conns {
		conn.Close()
	}
	p.conns = nil

	p.cond.Broadcast() // Notify all waiting goroutines that the pool is closed
}

// example usage:
func main() {
	pool := NewConnectionPool(5)
	wg := sync.WaitGroup{}

	// Simulate a task that requires connections
	task := func(id int) {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			conn, err := pool.Get()
			if err != nil {
				fmt.Printf("Goroutine %d: %v\n", id, err)
				return
			}
			fmt.Printf("Goroutine %d using connection %v\n", id, conn)

			pool.Put(conn)
		}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go task(i)
	}

	// After all tasks are done, close the pool
	fmt.Println("Closing pool")
	pool.Close()

	wg.Wait()
}
