package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// database cache ,don't use this impliments in your applicaion and use real database like varnish
var employeeData = map[int]string{
	1: "Alice",
	2: "Bob",
}

// Route prefixes for better organization and code maintainability
const (
	userRoute     = "/users"
	APIversion    = "/api/v1"
	employeeStore = APIversion + userRoute + "/{id}"
)

// csrfRequestHandler sets 'X-CSRF-Token' header in the response
func csrfRequestHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Add your token manipulation logic here,such as generating and storing the token
		csrfToken := "CSRF_TOKEN" // Replace this with your actual CSRF token Generation method
		http.SetCookie(w, &http.Cookie{
			Name:     "CSRF_TOKEN",
			Value:    csrfToken,
			Secure:   false, // or https when siutil using https
			HttpOnly: true,
			Path:     "/",
		})
		w.Header().Set("X-CSRF-Token", csrfToken)
		next.ServeHTTP(w, r)
	})
}

// apiRequestHandler Check for https and validate JWT token. for production Environments always restrict all api access with HTTPs
func apiRequestHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isHTTPS := r.URL.Scheme == "https"
		if !isHTTPS {
			http.Error(w, "HTTPS is required", http.StatusForbidden)
			return
		}

		// JWT validation LOGIC
		// Replace this with actual JWT validation from the request headers
		authToken := "Bearer INVALID_TOKEN" // replace  with valid header
		if authToken == "" {
			http.Error(w, "Invalid JWT token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// validID Request validator   to specify valid $id ranges,identifier formats for into RESTo API paths
func validID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, _ := strconv.Atoi(vars["id"])

		// Simulate validation of user's permission to read specific employees
		if id < 1 || id > len(employeeData) {
			http.Error(w, "Invalid user ID", http.StatusNotFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	employeeName := employeeData[id]

	// control access bbased on authentication and authorization
	//..
	fmt.Fprintf(w, "Employee Name: %s\n", employeeName)
}

func main() {
	r := mux.NewRouter()

	// Middleware is attached here to secure the routes
	r.Use(csrfRequestHandler)

	// Handle backend API routes
	r.Path(employeeStore).Methods("GET").Handler(apiRequestHandler(validID(http.HandlerFunc(getEmployee))))

	//Static Files
	r.PathPrefix("/assets").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir("./assets/"))))

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", r)
}
