package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	// ConsentCookieName is the name of the cookie used to store consent status
	ConsentCookieName = "consent"
	// ConsentDuration is the duration for which the consent cookie will remain valid
	ConsentDuration = 365 * 24 * time.Hour
)

func getConsent(w http.ResponseWriter, r *http.Request) bool {
	// Check if the consent cookie exists
	cookie, err := r.Cookie(ConsentCookieName)
	if err != nil {
		return false // Consent cookie not found, user has not given consent
	}

	// Parse the consent value from the cookie
	consent := cookie.Value == "true"
	return consent
}
func setConsent(w http.ResponseWriter, r *http.Request, consent bool) {
	// Create a new consent cookie with the given consent status
	cookie := http.Cookie{
		Name:     ConsentCookieName,
		Value:    fmt.Sprintf("%t", consent),
		Expires:  time.Now().Add(ConsentDuration),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	// Set the cookie in the response header
	http.SetCookie(w, &cookie)
}
func handleConsentForm(w http.ResponseWriter, r *http.Request) {
	// Handle the form submission where users consent to data processing
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error parsing form: %v", err)
		return
	}

	// Check if the user has given consent by checking the consent checkbox
	consent := r.FormValue("consent") == "true"

	// Set the consent cookie based on the user's decision
	setConsent(w, r, consent)

	// Redirect to the home page
	http.Redirect(w, r, "/", http.StatusFound)
}
func handleHome(w http.ResponseWriter, r *http.Request) {
	// Check if the user has given consent
	consent := getConsent(w, r)

	if consent {
		// User has given consent, show the sensitive data or handle sensitive operations
		fmt.Fprintf(w, "Welcome! You have given consent to process your data.")
	} else {
		// User has not given consent, show a consent form
		fmt.Fprintf(w, `
            <html>
                <head>
                    <title>Consent Form</title>
                </head>
                <body>
                    <form action="/consent" method="post">
                        <p>Please give your consent to process your data:</p>
                        <input type="checkbox" name="consent" value="true"> I consent<br>
                        <input type="submit" value="Submit">
                    </form>
                </body>
            </html>
        `)
	}
}
func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/consent", handleConsentForm)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
