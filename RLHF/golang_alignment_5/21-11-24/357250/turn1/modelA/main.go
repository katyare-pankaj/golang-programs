package main

import (
	"fmt"
	"go-programs/RLHF/golang_alignment_5/21-11-24/357250/turn1/modelA/email"
)

// Main function to run the code
func main() {
	// Load the email configuration
	config, err := email.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	// Create a Transport with the loaded configuration
	transport := email.NewTransport(config)

	// Create a Sender using the Transport
	sender := email.NewSender(transport)

	// Create an Email to send
	email := &email.Email{
		From:    "sender@example.com",
		To:      []string{"recipient@example.com"},
		Subject: "Test Email",
		Body:    "This is a test email.",
		IsHTML:  false,
	}

	// Send the email
	err = sender.Send(email)
	if err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Println("Email sent successfully!")
	}
}
