
package main  
import (  
    "fmt"
    "log"
    "net/http"
    "regexp"
    "strings"
)
const (  
    // Define a regex pattern for validating email addresses
    emailRegex = `^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`
)
//validateEmail returns true if the email is valid  
func validateEmail(email string) bool {  
    match, _ := regexp.MatchString(emailRegex, email)  
    return match
}
//sanitizeInput function sanitizes user input by removing any HTML tags or script elements.
func sanitizeInput(input string) string {
    // You can use various libraries like "github.com/microcosm-cc/bluemonday" for more robust HTML sanitization.
    return strings.ReplaceAll(strings.ReplaceAll(input, "<", "&lt;"), ">", "&gt;")
}
func handleCallback(w http.ResponseWriter, r *http.Request) {  
    if r.Method != "POST" {  
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)  
        return  
    }  
    defer r.Body.Close()  
    if err := r.ParseForm(); err != nil {  
        http.Error(w, "Invalid request body", http.StatusBadRequest)  
        return  
    }  
    // Validate and sanitize user input
    userName := sanitizeInput(r.FormValue("username"))
    email := sanitizeInput(r.FormValue("email"))
    if userName == "" || email == "" {
        http.Error(w, "Username or email cannot be empty", http.StatusBadRequest)
        return
    }
    if !validateEmail(email) {
        http.Error(w, "Invalid email address", http.StatusBadRequest)
        return
    }
    // Do something with the valid and sanitized user input
    fmt.Println("Callback received:")
    fmt.Println("Username:", userName)
    fmt.Println("Email:", email)
    w.WriteHeader(http.StatusOK)  
    fmt.Fprintf(w, "Success! Your callback has been processed.")  
}  
func main() {  
    fmt.Println("Secure Golang API Callback Server started...")  
    http.HandleFunc("/callback", handleCallback)  
    err := http.ListenAndServe(":8080", nil)  
    if err != nil {  
        log.Fatal("Error starting the server: ", err)  
    }  
} 
