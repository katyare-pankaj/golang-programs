package emailsender

import (
	"fmt"
	"net/smtp"
)

// Email represents an email to be sent.
type Email struct {
	From    string
	To      []string
	Subject string
	Body    string
}

// Sender is the interface for sending emails.
type Sender interface {
	Send(email Email) error
}

// SMTPSender sends emails using SMTP.
type SMTPSender struct {
	Host     string
	Port     int
	Username string
	Password string
}

// NewSMTPSender creates a new SMTPSender.
func NewSMTPSender(host string, port int, username, password string) *SMTPSender {
	return &SMTPSender{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}
}

// Send sends an email using SMTP.
func (s *SMTPSender) Send(email Email) error {
	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	to := email.To
	msg := []byte("To: " + email.To[0] + "\r\n" +
		"Subject: " + email.Subject + "\r\n" +
		"\r\n" +
		email.Body)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", s.Host, s.Port), auth, email.From, to, msg)
	return err
}

// TemplateRenderer is the interface for rendering email templates.
type TemplateRenderer interface {
	Render(templateName string, data interface{}) (string, error)
}

// SimpleTemplateRenderer is a simple implementation of TemplateRenderer.
type SimpleTemplateRenderer struct {
	Templates map[string]string
}

// NewSimpleTemplateRenderer creates a new SimpleTemplateRenderer.
func NewSimpleTemplateRenderer(templates map[string]string) *SimpleTemplateRenderer {
	return &SimpleTemplateRenderer{
		Templates: templates,
	}
}

// Render renders an email template using the provided data.
func (r *SimpleTemplateRenderer) Render(templateName string, data interface{}) (string, error) {
	template, ok := r.Templates[templateName]
	if !ok {
		return "", fmt.Errorf("template not found: %s", templateName)
	}
	return template, nil
}

// EmailSender is the main struct that provides the functionality to send emails.
type EmailSender struct {
	Sender           Sender
	TemplateRenderer TemplateRenderer
}

// NewEmailSender creates a new EmailSender.
func NewEmailSender(sender Sender, templateRenderer TemplateRenderer) *EmailSender {
	return &EmailSender{
		Sender:           sender,
		TemplateRenderer: templateRenderer,
	}
}

// SendEmail sends an email using the provided sender and template renderer.
func (e *EmailSender) SendEmail(email Email) error {
	renderedBody, err := e.TemplateRenderer.Render(email.Subject, nil)
	if err != nil {
		return err
	}
	email.Body = renderedBody
	return e.Sender.Send(email)
}
