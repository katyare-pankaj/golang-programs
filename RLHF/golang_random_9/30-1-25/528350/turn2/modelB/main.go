package main  
import (  
    "fmt"
    "regexp"
    "crypto/rand"
    "encoding/hex"
    "strings"
)

// GenerateRandomString generates a random string of given length
func GenerateRandomString(n int) string {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if err != nil {
        fmt.Println("Error generating random string:", err)
        return ""
    }
    return hex.EncodeToString(b)
}

// AnonymizeEmail anonymizes a given email address
func AnonymizeEmail(email string) string {
    atIndex := strings.Index(email, "@")
    if atIndex == -1 {
        // Return the original email if it doesn't contain '@'
        return email
    }

    // Generate a random string for the local-part
    randomLocalPart := GenerateRandomString(8)

    // Concatenate the random local-part with the domain
    return fmt.Sprintf("%s@%s", randomLocalPart, email[atIndex+1:])
}

// AnonymizePhoneNumber anonymizes a given phone number
func AnonymizePhoneNumber(phoneNumber string) string {
    // Remove any non-numeric characters from the phone number
    phoneNumber = regexp.MustCompile(`[^0-9]`).ReplaceAllString(phoneNumber, "")

    // Check if the phone number has at least 10 digits
    if len(phoneNumber) < 10 {
        // Return the original phone number if it doesn't have at least 10 digits
        return phoneNumber
    }

    // Generate a random string for the first 3 digits
    randomFirstThree := GenerateRandomString(3)

    // Generate a random string for the last 4 digits
    randomLastFour := GenerateRandomString(4)

    // Replace the first 3 and last 4 digits with random strings
    anonymizedPhoneNumber := regexp.MustCompile(`^([0-9]{3})([0-9]{3})([0-9]{4})$`).ReplaceAllString(phoneNumber, "$1"+randomFirstThree+"$3"+randomLastFour)

    return anonymizedPhoneNumber
}

// AnonymizeData takes a slice of strings representing sensitive user data and returns a new slice with data anonymized
func AnonymizeData(data []string) []string {
    emailPattern := regexp.MustCompile(`[\w._%+-]+@[\w.-]+\.[a-zA-Z]{2,}`)
    phonePattern := regexp.MustCompile(`\+?[0-9]{1,4}[ -]?$?[0-9]{1,3}?$?[ -]?[0-9]{3,4}[ -]?[0-9]{3,4}`)
    
    anonymizedData := make([]string, len(data))
    copy(anonymizedData, data)

    for i, value := range anonymizedData {
        anonymizedValue := emailPattern.ReplaceAllStringFunc(value, AnonymizeEmail)
        anonymizedValue = phonePattern.ReplaceAllStringFunc(anonymizedValue, AnonymizePhoneNumber)
        anonymizedData[i] = anonymizedValue
    }

    return anonymizedData
}

func main() {  
    data := []string{
		"John Doe john.doe@example.com 123-456-7890",
		"Jane Doe jane.doe@example.org +1 (555) 123-4567",
		"Alice alice@example.net 555-456-1234",
		"Bob bob@example.com +61 (4) 1234-5678",
		"Charlie charlie@example.com 01234 567 890",
	}

    anonymizedData := AnonymizeData(data)
	
    fmt.Println("Original Data:")