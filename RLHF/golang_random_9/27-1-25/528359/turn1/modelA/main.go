package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// ConnectionPool manages a pool of reusable connections.
type ConnectionPool struct {
	mu       sync.Mutex
	cond     *sync.Cond
	conns    []net.Conn
	capacity int
}

// NewConnectionPool creates a new pool with the specified capacity.
func NewConnectionPool(capacity int) *ConnectionPool {
	p := &ConnectionPool{
		capacity: capacity,
	}
	p.cond = sync.NewCond(&p.mu)
	return p
}

// Get retrieves a connection from the pool, waiting if necessary.
func (p *ConnectionPool) Get() net.Conn {
	p.mu.Lock()
	defer p.mu.Unlock()

	for len(p.conns) == 0 {
		// Wait for a connection to be put back into the pool.
		p.cond.Wait()
	}

	// Take a connection from the slice (LIFO strategy)
	conn := p.conns[len(p.conns)-1]
	p.conns = p.conns[:len(p.conns)-1]

	return conn
}

// Put returns a connection to the pool.
func (p *ConnectionPool) Put(conn net.Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.conns) < p.capacity {
		p.conns = append(p.conns, conn)
		p.cond.Signal()
	} else {
		conn.Close() // Can't add more, close the connection
	}
}

// Simulation functions, not meant for actual network handling.
func createConnection() net.Conn {
	// Simulate creating a new network connection
	return &dummyConn{id: time.Now().Unix()}
}

type dummyConn struct {
	id int64
}

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
	pool := NewConnectionPool(2)
	for i := 0; i < 5; i++ {
		go func(id int) {
			conn := createConnection()
			fmt.Printf("Created connection %d\n", conn.(*dummyConn).id)

			pool.Put(conn)
			reusedConn := pool.Get()
			fmt.Printf("Using connection %d\n", reusedConn.(*dummyConn).id)
			reusedConn.Close() // Always put back the connection
		}(i)
	}
	time.Sleep(5 * time.Second) // Wait to observe pool behavior
}
