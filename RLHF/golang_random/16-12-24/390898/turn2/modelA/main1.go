package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GatewayHandler(w http.ResponseWriter, r *http.Request) {
	// This example assumes traffic will be directed to either the old legacy API or the newer service
	if r.URL.Path == "/user" {
		response, err := http.Get("http://localhost:8081/user?id=" + r.URL.Query().Get("id"))
		if err != nil {
			http.Error(w, "Failed to get user", http.StatusBadGateway)
			return
		}
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)

		w.WriteHeader(response.StatusCode)
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
		return
	}
	// Here you could direct to the legacy system
}

func main() {
	http.HandleFunc("/", GatewayHandler)
	fmt.Println("API Gateway is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
