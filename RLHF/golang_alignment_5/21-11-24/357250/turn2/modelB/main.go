
package main  
import (  
    "fmt"
    "log"
    "os"
    "emailsender"
)

func main() {
    // Replace with your SMTP server details
    smtpHost := "smtp.example.com"
    smtpPort := 587
    smtpUsername := "your_smtp_username"
    smtpPassword := "your_smtp_password"

    // Create the SMTP sender
    sender := emailsender.NewSMTPSender(smtpHost, smtpPort, smtpUsername, smtpPassword)