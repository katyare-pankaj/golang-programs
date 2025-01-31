package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Service struct {
	Path     string
	Handler  http.Handler
	Username string
	Password string
}

type Gateway struct {
	services []Service
}

func (g *Gateway) RegisterService(path string, handler http.Handler, username, password string) {
	g.services = append(g.services, Service{Path: path, Handler: handler, Username: username, Password: password})
}

func (g *Gateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, service := range g.services {
		if strings.HasPrefix(r.URL.Path, service.Path) {
			// Check if basic authentication is required for this service
			if service.Username != "" && service.Password != "" {
				// Perform basic authentication
				username, password, ok := r.BasicAuth()
				if !ok || username != service.Username || password != service.Password {
					w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
					http.Error(w, "Unauthorized", http.StatusUnauthorized)
					return
				}
			}
			service.Handler.ServeHTTP(w, r)
			return
		}
	}
	http.Error(w, "Service not found", http.StatusNotFound)
}

func main() {
	gateway := &Gateway{}

	gateway.RegisterService("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}), "", "") // No authentication required for this service

	gateway.RegisterService("/api/users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Fetching list of users...")
	}), "user", "password") // Authentication required for this service

	gateway.RegisterService("/api/posts", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Fetching list of posts...")
	}), "", "") // No authentication required for this service

	fmt.Println("Gateway server running on port 8080")
	if err := http.ListenAndServe(":8080", gateway); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
