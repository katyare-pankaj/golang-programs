// main.go

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yourorg/reusable-code/emailtemplate"
)

func main() {
	http.HandleFunc("/signup", handleSignup)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Process signup form data and prepare email data
	emailData := map[string]string{
		"VolunteerName": "John Doe",
		"EventName":     "Charity Cleanup Event",
	}

	err := emailtemplate.SendEmail(w, "signup_confirmation.html", emailData)
	if err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Signup successful! Check your email for confirmation.")
}
