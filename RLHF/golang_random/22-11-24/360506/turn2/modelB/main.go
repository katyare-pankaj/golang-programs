package emailsender

import (
	"bytes"
	"fmt"
	"html/template"
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
		"MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n" +
		email.Body)
	err := smtp.SendMail(fmt.Sprintf("%s:%d", s.Host, s.Port), auth, email.From, to, msg)
	return err
}

// TemplateSender sends emails using SMTP and renders email content using templates.
type TemplateSender struct {
	Sender      Sender
	TemplateDir string // Directory containing email templates
}

// NewTemplateSender creates a new TemplateSender.
func NewTemplateSender(sender Sender, templateDir string) *TemplateSender {
	return &TemplateSender{
		Sender:      sender,
		TemplateDir: templateDir,
	}
}

// SendTemplateEmail sends an email using SMTP and rendering the email content using the provided template name and data.
func (t *TemplateSender) SendTemplateEmail(email Email, templateName string, data interface{}) error {
	// Parse the template
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/%s.html", t.TemplateDir, templateName))
	if err != nil {
		return err
	}

	// Render the template to a buffer
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return err
	}

	// Set the rendered template as the email body
	email.Body = buf.String()

	// Send the email using the underlying Sender
	return t.Sender.Send(email)
}
