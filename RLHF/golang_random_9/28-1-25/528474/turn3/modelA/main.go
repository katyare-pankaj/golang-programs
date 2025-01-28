package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// ConnectionPool manages a fixed number of reusable connections.
type ConnectionPool struct {
	mu       sync.Mutex
	cond     *sync.Cond
	conns    []net.Conn
	capacity int
	active   int
}

// NewConnectionPool initializes a ConnectionPool with a given capacity.
func NewConnectionPool(capacity int) *ConnectionPool {
	p := &ConnectionPool{
		capacity: capacity,
		conns:    make([]net.Conn, 0, capacity),
		active:   0,
	}
	p.cond = sync.NewCond(&p.mu)
	return p
}

// Get retrieves a connection from the pool, or waits for one to become available.
func (p *ConnectionPool) Get() (net.Conn, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for len(p.conns) == 0 && p.active >= p.capacity {
		p.cond.Wait() // Wait for a connection to be available.
	}

	// If there are available connections, reuse one
	if len(p.conns) > 0 {
		conn := p.conns[len(p.conns)-1]
		p.conns = p.conns[:len(p.conns)-1]
		return conn, nil
	}

	// Otherwise, create a new connection if under capacity
	p.active++
	return createConnection(), nil
}

// Put returns a connection to the pool, closing it if the pool is full or it is broken.
func (p *ConnectionPool) Put(conn net.Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if !isConnBroken(conn) && len(p.conns) < p.capacity {
		p.conns = append(p.conns, conn)
		p.cond.Signal() // Notify one waiting goroutine that a connection is available.
	} else {
		conn.Close() // Close broken connections or if pool is full.
		p.active--
	}
}

// isConnBroken checks if the connection is broken (dummy implementation).
func isConnBroken(conn net.Conn) bool {
	// Placeholder: Implement a proper check or monitoring logic
	return false
}

// createConnection simulates establishing a new network connection (for demonstration).
func createConnection() net.Conn {
	return &dummyConn{id: time.Now().UnixNano()}
}

type dummyConn struct {
	id int64
}

// Implement net.Conn interface for dummyConn struct
func (d *dummyConn) Read(b []byte) (n int, err error)  { return 0, nil }
func (d *dummyConn) Write(b []byte) (n int, err error) { return 0, nil }
func (d *dummyConn) Close() error {
	fmt.Printf("Closing connection %d\n", d.id)
	return nil
}
func (d *dummyConn) LocalAddr() net.Addr                { return nil }
func (d *dummyConn) RemoteAddr() net.Addr               { return nil }
func (d *dummyConn) SetDeadline(t time.Time) error      { return nil }
func (d *dummyConn) SetReadDeadline(t time.Time) error  { return nil }
func (d *dummyConn) SetWriteDeadline(t time.Time) error { return nil }

// Example usage
func main() {
	pool := NewConnectionPool(3)
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			conn, err := pool.Get()
			if err != nil {
				fmt.Printf("Goroutine %d failed to get a connection\n", id)
				return
			}
			fmt.Printf("Goroutine %d using connection %d\n", id, conn.(*dummyConn).id)

			// Simulate some work by sleeping
			time.Sleep(time.Second)

			pool.Put(conn) // Return the connection to the pool
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines finished")
}
