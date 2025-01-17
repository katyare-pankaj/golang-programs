package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"
)

// ResourceHandler is a struct that holds resources that need to be managed.
type ResourceHandler struct {
	file *os.File
	db   *sql.DB
	conn net.Conn
}

// NewResourceHandler creates and initializes a ResourceHandler with the needed resources.
func NewResourceHandler(filePath string, dbConnString string, networkAddr string) (*ResourceHandler, error) {
	handler := &ResourceHandler{}

	// Open a file and defer its closure
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	handler.file = file

	// Establish a database connection and defer its closure
	db, err := sql.Open("your-db-driver", dbConnString)
	if err != nil {
		handler.file.Close() // Ensure any allocated resource gets cleaned up in case of error
		return nil, fmt.Errorf("failed to open database: %v", err)
	}
	handler.db = db

	// Establish a network connection and defer its closure
	conn, err := net.Dial("tcp", networkAddr)
	if err != nil {
		handler.file.Close()
		handler.db.Close()
		return nil, fmt.Errorf("failed to establish network connection: %v", err)
	}
	handler.conn = conn

	return handler, nil
}

// PerformOperations performs some operations using the resources managed by ResourceHandler.
func (rh *ResourceHandler) PerformOperations() error {
	defer rh.cleanup()

	// Simulate some operations

	// If an error occurs, cleanup will still be called when returning early
	return nil
}

// cleanup ensures all resources are properly closed.
func (rh *ResourceHandler) cleanup() {
	if rh.conn != nil {
		fmt.Println("Closing network connection")
		rh.conn.Close()
	}

	if rh.db != nil {
		fmt.Println("Closing database connection")
		rh.db.Close()
	}

	if rh.file != nil {
		fmt.Println("Closing file")
		rh.file.Close()
	}
}
