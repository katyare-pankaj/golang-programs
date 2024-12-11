package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Server holds the HTTP server and a way to signal when to shut down
type Server struct {
	http.Server
	shutdown chan struct{}
	wg       sync.WaitGroup
}

// NewServer initializes a new Server instance
func NewServer(addr string) *Server {
	srv := &Server{
		shutdown: make(chan struct{}),
	}
	srv.Server.Addr = addr
	http.HandleFunc("/", srv.handleRequest)
	return srv
}

// handleRequest is a simple HTTP handler
func (srv *Server) handleRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

// Start begins listening for HTTP requests
func (srv *Server) Start() {
	srv.wg.Add(1)
	go func() {
		defer srv.wg.Done()
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("ListenAndServe(): %s\n", err)
		}
	}()
}

// Shutdown gracefully stops the server
func (srv *Server) Shutdown() {
	close(srv.shutdown)
	fmt.Println("Shutting down server gracefully...")
	if err := srv.Close(); err != nil {
		fmt.Printf("Error during server shutdown: %s\n", err)
	}
	srv.wg.Wait()
	fmt.Println("All goroutines finished. Server shut down.")
}

func main() {
	srv := NewServer(":8080")
	srv.Start()

	// Listen for termination signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received
	<-c
	srv.Shutdown()
}
