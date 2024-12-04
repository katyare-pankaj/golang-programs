package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// ConnectionPool represents a pool of connections
type ConnectionPool struct {
	maxConnections int
	connections    []*connection
	mutex          sync.Mutex
}

// connection is a simple placeholder for a real connection
type connection struct {
	id int
}

func (cp *ConnectionPool) Acquire(ctx context.Context) (*connection, error) {
	cp.mutex.Lock()
	defer cp.mutex.Unlock()

	if len(cp.connections) == 0 && cp.maxConnections == 0 {
		return nil, fmt.Errorf("connection pool is empty and max connections are not defined")
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		if len(cp.connections) == 0 {
			cp.connections = make([]*connection, cp.maxConnections)
			for i := range cp.connections {
				cp.connections[i] = &connection{id: i}
			}
		}

		con := cp.connections[len(cp.connections)-1]
		cp.connections = cp.connections[:len(cp.connections)-1]
		return con, nil
	}
}

func (cp *ConnectionPool) Release(ctx context.Context, con *connection) {
	cp.mutex.Lock()
	defer cp.mutex.Unlock()

	select {
	case <-ctx.Done():
		log.Printf("Release called but context canceled")
	default:
		cp.connections = append(cp.connections, con)
	}
}

func doWork(ctx context.Context, pool *ConnectionPool, id int) {
	con, err := pool.Acquire(ctx)
	if err != nil {
		log.Printf("Worker %d: Error acquiring connection: %v\n", id, err)
		return
	}
	defer func() {
		pool.Release(ctx, con)
	}()

	for {
		select {
		case <-ctx.Done():
			log.Printf("Worker %d: Context canceled, releasing connection\n", id)
			return
		default:
			fmt.Printf("Worker %d: Using connection %d\n", id, con.id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	pool := &ConnectionPool{maxConnections: 3}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go doWork(ctx, pool, 1)
	go doWork(ctx, pool, 2)
	go doWork(ctx, pool, 3)

	time.Sleep(3 * time.Second)

	log.Println("Main: Canceling context...")
	cancel()

	time.Sleep(2 * time.Second)

	log.Println("Main: Done")
}
