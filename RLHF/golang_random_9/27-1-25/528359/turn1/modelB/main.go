package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	// Max connections allowed in the pool
	maxConnections = 100
	// Buffer size for reading/writing data
	bufferSize = 1024
)

// Connection struct representing a connection to a client
type Connection struct {
	conn       net.Conn
	lastActive time.Time
}

// connectionPool slice holding all the active connections
var connectionPool []*Connection
var poolMutex sync.Mutex

func handleClient(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, bufferSize)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		// Handle the client request
		data := buf[:n]
		// Simple echo for demonstration
		conn.Write(data)

		// Update the connection's last active time
		poolMutex.Lock()
		for _, c := range connectionPool {
			if c.conn == conn {
				c.lastActive = time.Now()
				break
			}
		}
		poolMutex.Unlock()
	}
}

func manageConnections() {
	for {
		// Periodically check for inactive connections and close them
		time.Sleep(5 * time.Second)
		poolMutex.Lock()
		for i, c := range connectionPool {
			if time.Since(c.lastActive) > 30*time.Second {
				fmt.Println("Closing inactive connection:", c.conn.RemoteAddr())
				c.conn.Close()
				connectionPool = append(connectionPool[:i], connectionPool[i+1:]...)
			}
		}
		poolMutex.Unlock()
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	// Start managing connections in the background
	go manageConnections()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}
		// Limit the number of connections in the pool
		poolMutex.Lock()
		if len(connectionPool) >= maxConnections {
			fmt.Println("Connection pool full, rejecting connection from:", conn.RemoteAddr())
			conn.Close()
			poolMutex.Unlock()
			continue
		}
		// Add the new connection to the pool
		newConn := &Connection{conn: conn, lastActive: time.Now()}
		connectionPool = append(connectionPool, newConn)
		poolMutex.Unlock()

		fmt.Println("New connection accepted from:", conn.RemoteAddr())
		go handleClient(conn)
	}
}
