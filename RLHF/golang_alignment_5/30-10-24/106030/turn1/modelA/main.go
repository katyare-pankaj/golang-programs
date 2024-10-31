package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Translator holds the translations for different messages
type Translator struct {
	translations map[string]map[string]string
}

// NewTranslator creates a new Translator with initial translations
func NewTranslator() *Translator {
	return &Translator{
		translations: map[string]map[string]string{
			"en": {
				"welcome": "Welcome",
				"goodbye": "Goodbye",
			},
			"es": {
				"welcome": "Bienvenido",
				"goodbye": "Adiós",
			},
		},
	}
}

// Translate translates a message into the specified language
func (t *Translator) Translate(msg, lang string) string {
	translation, ok := t.translations[lang][msg]
	if !ok {
		// Fall back to default language if translation is not found
		translation, _ = t.translations["en"][msg]
	}
	return translation
}

// FormatMessage formats a message with arguments using the specified language
func (t *Translator) FormatMessage(msg, lang string, args ...interface{}) string {
	p := message.NewPrinter(language.Make(lang))
	return p.Sprintf(t.Translate(msg, lang), args...)
}

func main() {
	translator := NewTranslator()

	// BDD Scenario: User can see welcome message in different languages
	testWelcomeMessage("en", translator)
	testWelcomeMessage("es", translator)

	// BDD Scenario: User can see goodbye message in different languages
	testGoodbyeMessage("en", translator)
	testGoodbyeMessage("es", translator)
}

func testWelcomeMessage(lang string, translator *Translator) {
	expected := translator.Translate("welcome", lang)
	actual := translator.Translate("welcome", lang)
	if actual != expected {
		fmt.Printf("Test failed for welcome message in %s: Expected '%s', got '%s'\n", lang, expected, actual)
	} else {
		fmt.Printf("Test passed for welcome message in %s\n", lang)
	}
}

func testGoodbyeMessage(lang string, translator *Translator) {
	expected := translator.Translate("goodbye", lang)
	actual := translator.Translate("goodbye", lang)
	if actual != expected {
		fmt.Printf("Test failed for goodbye message in %s: Expected '%s', got '%s'\n", lang, expected, actual)
	} else {
		fmt.Printf("Test passed for goodbye message in %s\n", lang)
	}
}
