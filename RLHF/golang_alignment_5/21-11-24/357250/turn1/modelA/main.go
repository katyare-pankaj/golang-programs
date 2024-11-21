package email

import "fmt"

// Email represents an email message.
type Email struct {
	From    string
	To      []string
	Subject string
	Body    string
	IsHTML  bool
}

// Sender interface defines the contract for sending an email.
type Sender interface {
	Send(email *Email) error
}

// Config contains the email configuration.
type Config struct {
	SMTPHost     string
	SMTPPort     int
	SMTPUsername string
	SMTPPassword string
	TLS          bool
}

// Transport is a struct to handle SMTP transport.
type Transport struct {
	config *Config
}

// NewTransport creates a new Transport instance.
func NewTransport(config *Config) *Transport {
	return &Transport{config: config}
}

// SendEmail sends an email using the configured SMTP transport.
func (t *Transport) SendEmail(email *Email) error {
	// Simplified logic for demonstration purposes.
	// In practice, you would use the net/smtp package here.
	return nil
}

// NewSender creates a new Sender instance.
func NewSender(transport Sender) *SenderImpl {
	return &SenderImpl{transport: transport}
}

// SenderImpl implements the Sender interface.
type SenderImpl struct {
	transport Sender
}

// Send sends an email using the configured transport.
func (s *SenderImpl) Send(email *Email) error {
	// Optionally, validate the email before sending
	if email == nil || email.From == "" || email.To == nil || len(email.To) == 0 {
		return fmt.Errorf("invalid email message")
	}

	// Send the email using the transport
	return s.transport.SendEmail(email)
}

// LoadConfig loads email configuration from a file or environment variables.
func LoadConfig() (*Config, error) {
	// Placeholder for actual configuration loading logic.
	// This could include parsing a yaml file, reading env vars, etc.
	config := &Config{
		SMTPHost:     "smtp.example.com",
		SMTPPort:     587,
		SMTPUsername: "user@example.com",
		SMTPPassword: "password",
		TLS:          true,
	}
	return config, nil
}
