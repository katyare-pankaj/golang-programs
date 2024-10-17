package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendEmail(t *testing.T) {
	sender := "sender@example.com"
	recipient := "recipient@example.com"
	subject := "Test Email"
	body := "This is a test email."

	err := SendEmail(sender, recipient, subject, body)
	assert.NoError(t, err, "failed to send email")
}
