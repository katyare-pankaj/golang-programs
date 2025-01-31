package main

import (
	"fmt"
	"net/http"
	"strings"
)

type service struct {
	path    string
	handler func(http.ResponseWriter, *http.Request)
}
type gateway struct {
	services []service
}

func (g *gateway) RegisterService(path string, handler func(http.ResponseWriter, *http.Request)) {
	g.services = append(g.services, service{path, handler})
}
func (g *gateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, s := range g.services {
		if strings.HasPrefix(r.URL.Path, s.path) {
			s.handler(w, r)
			return
		}
	}
	http.Error(w, "Not found", http.StatusNotFound)
}
func main() {
	g := gateway{}
	// Define and register service handlers
	g.RegisterService("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	g.RegisterService("/api/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "List of users")
	})
	g.RegisterService("/api/posts", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "List of posts")
	})
	// Start the gateway server
	fmt.Println("Gateway server running on port 8080")
	err := http.ListenAndServe(":8080", &g)
	if err != nil {
		panic(err)
	}
}
