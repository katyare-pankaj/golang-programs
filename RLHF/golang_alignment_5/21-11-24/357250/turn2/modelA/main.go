package main

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"text/template"
)

// EmailConfig holds SMTP configuration.
type EmailConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

// Email holds email details.
type Email struct {
	From     string
	To       []string
	Subject  string
	Template string // Template file name
	Data     interface{}
}

// SendEmail sends an email using SMTP after rendering the template.
func SendEmail(config EmailConfig, email Email) error {
	// Render the template with the provided data
	tpl, err := template.ParseFiles(email.Template)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}
	var renderedBody bytes.Buffer
	if err := tpl.Execute(&renderedBody, email.Data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// Create the email message
	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)
	to := email.To
	msg := []byte("To: " + email.To[0] + "\r\n" +
		"Subject: " + email.Subject + "\r\n" +
		"\r\n" +
		renderedBody.String())

	// Send the email
	err = smtp.SendMail(fmt.Sprintf("%s:%d", config.Host, config.Port), auth, email.From, to, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}

func main() {
	// Example SMTP configuration
	config := EmailConfig{
		Host:     "smtp.gmail.com",
		Port:     587,
		Username: "your-username@example.com",
		Password: "your-password",
	}

	// Example email with template
	email := Email{
		From:     "sender@example.com",
		To:       []string{"recipient@example.com"},
		Subject:  "Welcome Email",
		Template: "welcome.tmpl",
		Data:     map[string]string{"Name": "John Doe"},
	}

	// Send the email
	if err := SendEmail(config, email); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Email sent successfully!")
}
