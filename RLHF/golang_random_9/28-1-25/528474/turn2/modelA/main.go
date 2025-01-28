package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// ConnectionPool handles a pool of net.Conn objects.
type ConnectionPool struct {
	mu       sync.Mutex
	cond     *sync.Cond
	conns    []net.Conn
	capacity int
}

// NewConnectionPool initializes a ConnectionPool with a given capacity.
func NewConnectionPool(capacity int) *ConnectionPool {
	p := &ConnectionPool{
		capacity: capacity,
		conns:    make([]net.Conn, 0, capacity),
	}
	p.cond = sync.NewCond(&p.mu)
	return p
}

// Get retrieves a connection from the pool or waits if none are available.
func (p *ConnectionPool) Get() net.Conn {
	p.mu.Lock()
	defer p.mu.Unlock()

	for len(p.conns) == 0 {
		p.cond.Wait() // Wait for a connection to be available.
	}

	// Pop a connection from the slice (LIFO strategy)
	conn := p.conns[len(p.conns)-1]
	p.conns = p.conns[:len(p.conns)-1]

	return conn
}

// Put adds a connection back to the pool, closing it if the pool is full.
func (p *ConnectionPool) Put(conn net.Conn) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.conns) < p.capacity {
		p.conns = append(p.conns, conn)
		p.cond.Signal() // Notify one waiting goroutine that a connection is available.
	} else {
		conn.Close() // Close the connection if the pool is full.
	}
}

// createConnection simulates establishing a new network connection (for demonstration).
func createConnection() net.Conn {
	return &dummyConn{id: time.Now().UnixNano()}
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

func main() {
	pool := NewConnectionPool(3)
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			conn := createConnection()
			fmt.Printf("Created connection %d\n", conn.(*dummyConn).id)

			pool.Put(conn)
			reusedConn := pool.Get()
			fmt.Printf("Goroutine %d using connection %d\n", id, reusedConn.(*dummyConn).id)
			pool.Put(reusedConn)
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines finished")
}
