package main

import (
	"fmt"
	"net/http"
	"strings"
)

// Service struct holds the path prefix and the associated HTTP handler.
type Service struct {
	Path    string
	Handler func(http.ResponseWriter, *http.Request)
}

// Gateway struct to manage services.
type Gateway struct {
	services []Service
}

// RegisterService allows adding new services to the gateway.
func (g *Gateway) RegisterService(path string, handler func(http.ResponseWriter, *http.Request)) {
	g.services = append(g.services, Service{Path: path, Handler: handler})
}

// ServeHTTP handles incoming requests and forwards them to the appropriate service.
func (g *Gateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, service := range g.services {
		if strings.HasPrefix(r.URL.Path, service.Path) {
			service.Handler(w, r)
			return
		}
	}
	http.Error(w, "Service not found", http.StatusNotFound)
}

func main() {
	// Create a new instance of Gateway
	gateway := &Gateway{}

	// Register services with their respective handlers
	gateway.RegisterService("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	gateway.RegisterService("/api/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Fetching list of users...")
	})

	gateway.RegisterService("/api/posts", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Fetching list of posts...")
	})

	// Start the HTTP server with the gateway handling requests
	fmt.Println("Gateway server running on port 8080")
	if err := http.ListenAndServe(":8080", gateway); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
