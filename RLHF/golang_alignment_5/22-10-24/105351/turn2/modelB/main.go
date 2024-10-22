package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"text/template"

	"github.com/coreos/etcd/store"
	"github.com/gorilla/mux"
)

// Define a struct to represent employee data
type Employee struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

func main() {
	// Initialize router
	r := mux.NewRouter()
	// Route handler for the employee data entry page
	r.HandleFunc("/employee/data", employeeDataHandler).Methods("GET", "POST")

	// Start the server
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func employeeDataHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")

	switch r.Method {
	case "GET":
		// Handle GET request: Show the employee data entry form
		tmpl := template.Must(template.ParseFiles("employee_data_entry.html"))
		tmpl.Execute(w, nil)
	case "POST":
		// Handle POST request: Process the employee data entry
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Create an employee struct from the form data
		employee := Employee{
			FirstName: r.FormValue("first_name"),
			LastName:  r.FormValue("last_name"),
			Email:     r.FormValue("email"),
		}
		// Perform data validation and consistency checks
		if !isValidEmployeeData(employee) {
			// Handle invalid data and display error message to the user
			session.AddFlash("Invalid data provided. Please check the fields and try again.")
			session.Save(r, w)
			http.Redirect(w, r, "/employee/data", http.StatusSeeOther)
			return
		}
		// If data is valid, proceed with further processing (e.g., saving to the HRIS system)
		// For this example, we'll just display a success message.
		session.AddFlash("Employee data saved successfully!")
		session.Save(r, w)
		http.Redirect(w, r, "/employee/data", http.StatusSeeOther)
	}
}

func isValidEmployeeData(employee Employee) bool {
	// Perform data consistency checks here
	// For example, you can check for empty strings, valid email format, etc.
	if employee.FirstName == "" || employee.LastName == "" || employee.Email == "" {
		return false
	}
	// Simple email validation using regular expression
	const emailRegex = `^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`
	if !matchString(emailRegex, employee.Email) {
		return false
	}
	return true
}

func matchString(regex string, s string) bool {
	match, _ := regexp.MatchString(regex, s)
	return match
}
