package main

import (
	"fmt"
	"net/smtp"
)

func SendEmail(sender, recipient, subject, body string) error {
	msg := []byte("To: " + recipient + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	err := smtp.SendMail("smtp.example.com:587",
		smtp.PlainAuth("", sender, "your_password", "smtp.example.com"),
		sender, []string{recipient}, msg)

	if err != nil {
		return fmt.Errorf("error sending email: %w", err)
	}

	return nil
}

func main() {
	sender := "sender@example.com"
	recipient := "recipient@example.com"
	subject := "Hello"
	body := "Hello, this is a test email."

	err := SendEmail(sender, recipient, subject, body)
	if err != nil {
		fmt.Println("error sending email:", err)
	}
}
