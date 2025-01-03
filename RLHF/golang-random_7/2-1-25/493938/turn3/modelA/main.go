package main

import (
	"fmt"
	"log"
)

// ErrorMessageTemplate represents an error message template for a specific language
type ErrorMessageTemplate struct {
	ID       int
	Language string
	Template string
}

// ErrorMessageTemplates is a map of error message templates
var ErrorMessageTemplates = map[int]map[string]ErrorMessageTemplate{
	1: {
		"en": {ID: 1, Language: "en", Template: "Invalid ID: %v"},
		"es": {ID: 1, Language: "es", Template: "ID inválido: %v"},
	},
	2: {
		"en": {ID: 2, Language: "en", Template: "Unauthorized action: %v"},
		"es": {ID: 2, Language: "es", Template: "Acción no autorizada: %v"},
	},
}

// GetErrorMessage retrieves an error message template based on error code and language
func GetErrorMessage(code int, language string) string {
	template, ok := ErrorMessageTemplates[code][language]
	if !ok {
		// Default to English if language not found
		enTemplate, okEn := ErrorMessageTemplates[code]["en"]
		if !okEn {
			// Default to an unknown error message
			return "Unknown error"
		}
		return enTemplate.Template
	}
	return template.Template
}

// FormatError formats an error message using a template and additional context
func FormatError(code int, language string, context interface{}) string {
	template := GetErrorMessage(code, language)
	return fmt.Sprintf(template, context)
}

func main() {
	// Example user context
	userLanguage := "es"

	// Example error scenarios
	errorID1 := 1
	errorID2 := 2
	errorContext1 := "user123"
	errorContext2 := "action delete"

	// Format and print errors
	errorMessage1 := FormatError(errorID1, userLanguage, errorContext1)
	errorMessage2 := FormatError(errorID2, userLanguage, errorContext2)

	log.Printf("Error Message 1: %v\n", errorMessage1)
	log.Printf("Error Message 2: %v\n", errorMessage2)
}
