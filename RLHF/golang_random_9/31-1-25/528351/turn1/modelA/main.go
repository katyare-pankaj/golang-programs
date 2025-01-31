package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Service struct {
	Name string
	URL  string
}

type APIGateway struct {
	Services []Service
}

func NewAPIGateway() *APIGateway {
	return &APIGateway{Services: []Service{}}
}

func (gateway *APIGateway) RegisterService(name, url string) {
	service := Service{Name: name, URL: url}
	gateway.Services = append(gateway.Services, service)
}

func (gateway *APIGateway) RouteRequest(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) < 2 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	serviceName := pathParts[1]
	servicePath := strings.Join(pathParts[2:], "/")

	service, err := gateway.findService(serviceName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	proxyURL := fmt.Sprintf("%s/%s", service.URL, servicePath)
	resp, err := http.Get(proxyURL)
	if err != nil {
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func (gateway *APIGateway) findService(name string) (Service, error) {
	for _, service := range gateway.Services {
		if service.Name == name {
			return service, nil
		}
	}
	return Service{}, fmt.Errorf("service not found: %s", name)
}

func main() {
	gateway := NewAPIGateway()

	// Register services
	gateway.RegisterService("service1", "http://localhost:8081")
	gateway.RegisterService("service2", "http://localhost:8082")

	http.HandleFunc("/", gateway.RouteRequest)

	fmt.Println("API Gateway is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
