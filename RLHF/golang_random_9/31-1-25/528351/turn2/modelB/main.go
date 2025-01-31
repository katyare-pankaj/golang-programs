package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

type service struct {
	path    string
	handler http.Handler
	active  bool
	mutex   sync.Mutex
}

func (s *service) SetActive(active bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.active = active
}
func (s *service) IsActive() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.active
}

type gateway struct {
	services []service
	logger   *log.Logger
}

func newGateway(logger *log.Logger) *gateway {
	return &gateway{logger: logger}
}
func (g *gateway) RegisterService(path string, handler http.Handler) {
	g.services = append(g.services, service{path, handler, true, sync.Mutex{}})
}
func (g *gateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.logger.Printf("%s %s", r.Method, r.URL.Path)
	for _, s := range g.services {
		if strings.HasPrefix(r.URL.Path, s.path) && s.IsActive() {
			s.handler.ServeHTTP(w, r)
			return
		}
	}
	http.Error(w, "Not found", http.StatusNotFound)
}
func main() {
	logger := log.New(os.Stdout, "gateway: ", log.Ldate|log.Ltime|log.Lshortfile)
	g := newGateway(logger)
	// Define and register service handlers
	g.RegisterService("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}))
	g.RegisterService("/api/users", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "List of users")
	}))
	g.RegisterService("/api/posts", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "List of posts")
	}))
	// Start the gateway server
	fmt.Println("Gateway server running on port 8080")
	err := http.ListenAndServe(":8080", g)
	if err != nil {
		panic(err)
	}
}
