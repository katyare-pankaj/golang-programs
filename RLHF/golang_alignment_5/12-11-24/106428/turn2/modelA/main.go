package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Set DNS security headers
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' https://*.googleapis.com; style-src 'self' https://*.googleapis.com; img-src 'self' data: https://*.googleapis.com;")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
	w.Header().Set("Feature-Policy", "geolocation 'none'; microphone 'none'; camera 'none'; magnetometer 'none'; gyroscope 'none'; speaker 'none'; fullscreen 'self'; payment 'none'")
	w.Header().Set("Permissions-Policy", "interest-cohort=()")

	fmt.Fprintf(w, "Welcome to our nonprofit application!")
}
