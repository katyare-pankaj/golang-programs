package main

import (
	"fmt"
)

// ErrorTemplate represents a template for an error message in a specific language
type ErrorTemplate struct {
	Language string
	Template string
}

// ErrorCustomizer holds the error templates for different languages
type ErrorCustomizer struct {
	templates map[ErrorCode]map[string]string
}

// NewErrorCustomizer creates a new ErrorCustomizer with empty templates
func NewErrorCustomizer() *ErrorCustomizer {
	return &ErrorCustomizer{
		templates: make(map[ErrorCode]map[string]string),
	}
}

// ErrorCode is an enum for different types of errors
type ErrorCode int

const (
	InvalidInputError ErrorCode = iota
	InternalServerError
)

// AddTemplate adds a new error template for a specific error code and language
func (ec *ErrorCustomizer) AddTemplate(code ErrorCode, template ErrorTemplate) {
	if _, ok := ec.templates[code]; !ok {
		ec.templates[code] = make(map[string]string)
	}
	ec.templates[code][template.Language] = template.Template
}

// GetMessage retrieves the error message for a specific error code and language, using a default language if the requested one is not found
func (ec *ErrorCustomizer) GetMessage(code ErrorCode, language string) string {
	if templates, ok := ec.templates[code]; ok {
		if msg, ok := templates[language]; ok {
			return msg
		}
		// Fall back to the default language if the requested one is not found
		for lang, msg := range templates {
			return msg
		}
	}
	return "Unknown error" // Default error message if no template is found
}

// FormatError formats an error message using a custom template for the specified language
func (ec *ErrorCustomizer) FormatError(code ErrorCode, language string, args ...interface{}) string {
	template := ec.GetMessage(code, language)
	return fmt.Sprintf(template, args...)
}

func main() {
	// Create an error customizer
	ec := NewErrorCustomizer()

	// Add error templates for different languages
	ec.AddTemplate(InvalidInputError, ErrorTemplate{Language: "en", Template: "Invalid input: %s"})
	ec.AddTemplate(InvalidInputError, ErrorTemplate{Language: "es", Template: "Entrada inválida: %s"})
	ec.AddTemplate(InternalServerError, ErrorTemplate{Language: "en", Template: "Internal server error"})
	ec.AddTemplate(InternalServerError, ErrorTemplate{Language: "es", Template: "Error interno del servidor"})

	// Simulate errors and format them with custom messages in different languages
	invalidInputError := ec.FormatError(InvalidInputError, "es", "Invalid email format")
	internalServerError := ec.FormatError(InternalServerError, "en")

	fmt.Println("Invalid Input Error (es):", invalidInputError)
	fmt.Println("Internal Server Error (en):", internalServerError)

	// Handle an error without a specific template for the requested language
	unknownError := ec.FormatError(100, "fr", "This is an unknown error")
	fmt.Println("Unknown Error (fr):", unknownError)
}
