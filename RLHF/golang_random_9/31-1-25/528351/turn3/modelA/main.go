package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

// Service struct holds the path prefix, the handler, and optional authentication credentials.
type Service struct {
	Path          string
	Handler       func(http.ResponseWriter, *http.Request)
	RequireAuth   bool
	Username      string
	Password      string
	DynamicParams []string
}

// Gateway struct to manage services.
type Gateway struct {
	services []Service
}

// RegisterService allows adding new services to the gateway.
func (g *Gateway) RegisterService(service Service) {
	g.services = append(g.services, service)
}

// ParsePath extracts dynamic parameters from the route.
func (s *Service) ParsePath(requestPath string) map[string]string {
	pathParts := strings.Split(requestPath, "/")
	serviceParts := strings.Split(s.Path, "/")

	params := make(map[string]string)
	for index, part := range serviceParts {
		if strings.HasPrefix(part, ":") {
			paramName := part[1:] // remove ':' prefix
			params[paramName] = pathParts[index]
		}
	}
	return params
}

// BasicAuth performs basic HTTP authentication.
func (s *Service) BasicAuth(w http.ResponseWriter, r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}

	// Extract base64 payload from the Authorization header
	payload := strings.TrimPrefix(authHeader, "Basic ")
	decoded, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}

	payloadParts := strings.SplitN(string(decoded), ":", 2)
	if len(payloadParts) != 2 {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}

	username, password := payloadParts[0], payloadParts[1]
	if username != s.Username || password != s.Password {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return false
	}

	return true
}

// ServeHTTP handles incoming requests and forwards them to the appropriate service.
func (g *Gateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, service := range g.services {
		if matches, err := pathMatches(service.Path, r.URL.Path); err == nil && matches {
			// Perform basic authentication if required
			if service.RequireAuth && !service.BasicAuth(w, r) {
				return
			}

			params := service.ParsePath(r.URL.Path)
			fmt.Fprintf(w, "Dynamic Params: %v\n", params)

			service.Handler(w, r)
			return
		}
	}
	http.Error(w, "Service not found", http.StatusNotFound)
}

func pathMatches(servicePath, requestPath string) (bool, error) {
	// Compare paths, considering dynamic segments
	serviceParts := strings.Split(servicePath, "/")
	requestParts := strings.Split(requestPath, "/")
	if len(serviceParts) != len(requestParts) {
		return false, nil
	}
	for i := range serviceParts {
		if !strings.HasPrefix(serviceParts[i], ":") && serviceParts[i] != requestParts[i] {
			return false, nil
		}
	}
	return true, nil
}

func main() {
	// Create a new instance of Gateway
	gateway := &Gateway{}

	// Register services with their respective handlers and authentication
	gateway.RegisterService(Service{
		Path:    "/api/users/:userId",
		Handler: userHandler,
	})

	gateway.RegisterService(Service{
		Path:        "/api/posts",
		Handler:     postHandler,
		RequireAuth: true,
		Username:    "admin",
		Password:    "secret",
	})

	// Start the HTTP server with the gateway handling requests
	fmt.Println("Gateway server running on port 8080")
	if err := http.ListenAndServe(":8080", gateway); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handling User Service\n")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Handling Post Service\n")
}
