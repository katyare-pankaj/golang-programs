package main

import (
	"fmt"
	"go-programs/RLHF/golang_alignment_5/21-11-24/357250/turn1/modelB/emailsender"
	"log"
)

type MockSender struct{}

// Send simulates sending an email by just printing the email details.
func (m *MockSender) Send(email emailsender.Email) error {
	fmt.Println("Simulating sending email...")
	fmt.Println("From:", email.From)
	fmt.Println("To:", email.To)
	fmt.Println("Subject:", email.Subject)
	fmt.Println("Body:", email.Body)
	return nil
}

func main() {
	// Create a MockSender instead of the real SMTPSender
	mockSender := &MockSender{}

	// Define some email templates (for demo purposes)
	templates := map[string]string{
		"welcome": "Welcome, {{.Name}}! We are glad to have you.",
	}

	// Create a simple template renderer
	templateRenderer := emailsender.NewSimpleTemplateRenderer(templates)

	// Create an EmailSender with the mock sender and template renderer
	emailSender := emailsender.NewEmailSender(mockSender, templateRenderer)

	// Prepare the email data
	email := emailsender.Email{
		From:    "your_email@gmail.com",
		To:      []string{"recipient@example.com"},
		Subject: "welcome", // Use the template name
		Body:    "",        // Body will be rendered from the template
	}

	// Simulate sending the email
	err := emailSender.SendEmail(email)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	} else {
		fmt.Println("Email simulation successful!")
	}
}
