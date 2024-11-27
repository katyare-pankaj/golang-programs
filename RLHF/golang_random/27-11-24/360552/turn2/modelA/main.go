package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	// Simple example JWT secret key
	secretKey = []byte("your-secret-key-here")

	// Backend URLs
	backends = []string{"http://localhost:8081", "http://localhost:8082"}
)

// jwtMiddleware validates the incoming JWT in the Authorization header.
func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("Authorization")
		if !strings.HasPrefix(tokenStr, "Bearer ") {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		token, err := jwt.Parse(tokenStr[7:], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHS256); !ok {
				return nil, fmt.Errorf("Unsupported signing method: %v", token.Method)
			}
			return secretKey, nil
		})

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Token expired", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// logMiddleware logs the request before passing it to the next middleware.
func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Printf("Request: %s %s\n", r.Method, r.URL.Path)
		defer func() {
			duration := time.Now().Sub(startTime)
			status := w.Header().Get("Status")
			if status == "" {
				status = http.StatusText(w.WriteHeader(http.StatusInternalServerError))
			}
			log.Printf("Response: %s %s %s\n", status, duration.String(), r.RemoteAddr)
		}()
		next.ServeHTTP(w, r)
	})
}

// roundRobin selects a backend URL using round-robin logic.
func roundRobin() string {
	index := (time.Now().UnixNano() % int64(len(backends))) % int64(len(backends))
	return backends[index]
}

// proxy handles incoming requests and proxies them to the specified backend URL.
func proxy(backend string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/") {
			newPath := fmt.Sprintf("%s%s", backend, r.URL.Path)
			r.URL.Path = newPath
			r.URL.Scheme = "http"
			r.URL.Host = backend[strings.IndexByte(backend, ':')+1:]
		} else {
			r.URL.Path = fmt.Sprintf("%s/%s", backend, r.URL.Path)
		}
		r.URL.Host = backend
		client := &http.Client{}
		req, err := http.NewRequestWithContext(r.Context(), r.Method, r.URL.String(), r.Body)
		if err != nil {
			log.Fatalf("Error creating request: %v", err)
		}

		for k, v := range r.Header {
			req.Header[k] = v
		}
		req.Header.Set("Host", r.Host)

		res, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
			return
		}
		defer res.Body.Close()

		for k, v := range res.Header {
			w.Header()[k] = v
		}
		w.WriteHeader(res.StatusCode)
		_, err = http.Ioutil.Copy(w, res.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
		}
	}
}

func main() {
	log.Println("Starting API Gateway...")

	backend := roundRobin()
	http.HandleFunc("/api/v1/", logMiddleware(jwtMiddleware(proxy(backend))))

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
