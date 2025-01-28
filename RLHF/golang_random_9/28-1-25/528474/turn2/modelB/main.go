package main  
import (  
    "fmt"
    "net"
    "sync"
    "time"
)
// ConnectionPool manages a pool of reusable connections.
type ConnectionPool struct {
    mu      sync.Mutex
    active  map[net.Conn]bool  
    idle     chan net.Conn
    capacity int
    closed   bool
}
// NewConnectionPool creates a new pool with the specified capacity.
func NewConnectionPool(capacity int) *ConnectionPool {
    p := &ConnectionPool{
        active:   make(map[net.Conn]bool),
        idle:     make(chan net.Conn),
        capacity: capacity,
    }
    return p
}
// Get retrieves a connection from the pool, waiting if necessary.
func (p *ConnectionPool) Get() (net.Conn, error) {
    p.mu.Lock()
    defer p.mu.Unlock()
    if p.closed {
        return nil, fmt.Errorf("pool is closed")
    }
    for len(p.idle) == 0 && len(p.active) < p.capacity {
        // Create a new connection if there are no idle ones and we are below capacity
        conn, err := createConnection()
        if err != nil {
            return nil, err
        }
        p.active[conn] = true
        go p.releaseConn(conn)
    }
    select {
    case conn := <-p.idle:
        return conn, nil
    default:
        // If there are no idle connections, return an error (pool is full)
        return nil, fmt.Errorf("pool is full")
    }
}

// Put returns a connection to the pool.
func (p *ConnectionPool) Put(conn net.Conn) error {
    p.mu.Lock()
    defer p.mu.Unlock()
    if _, ok := p.active[conn]; !ok {
        return fmt.Errorf("connection not in pool")
    }
    delete(p.active, conn)
    select {
    case p.idle <- conn:
    default:
        // If the idle channel is full, close the connection
        return conn.Close()
    }
    return nil
}
// Close closes the pool, disconnecting all active connections.
func (p *ConnectionPool) Close() error {
    p.mu.Lock()
    defer p.mu.Unlock()
    if p.closed {
        return nil
    }
    p.closed = true
    close(p.idle)
    for conn := range p.active {
        if err := conn.Close(); err != nil {
            return err
        }
    }
    return nil
}
func createConnection() (net.Conn, error) {
    // Simulate creating a new network connection
    return &dummyConn{id: time.Now().Unix()}, nil
}

// dummyConn is a placeholder for actual network connections.
type dummyConn struct {
    id int64
}

func (d *dummyConn) Read(b []byte) (n int, err error) { return 0, nil }
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

// Simulate connection use by a worker.
func worker(pool *ConnectionPool, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 3; i++ {
        conn, err := pool.Get()
        if err != nil {
            fmt.Println("Error getting connection:", err)
            return