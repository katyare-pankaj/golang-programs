package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Translator holds translations for different messages
type Translator struct {
	translations map[string]map[string]string
}

// NewTranslator creates a new Translator
func NewTranslator() *Translator {
	return &Translator{
		translations: map[string]map[string]string{
			"en": {
				"date_format":     "January 2, 2006",
				"currency_format": "$%.2f",
			},
			"es": {
				"date_format":     "02 de enero de 2006",
				"currency_format": "€%.2f",
			},
		},
	}
}

// Translate translates a message
func (t *Translator) Translate(msg, lang string) string {
	translation, ok := t.translations[lang][msg]
	if !ok {
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

	// BDD Scenario: Data synchronization for different languages and formats
	testDataSynchronization("en", translator)
	testDataSynchronization("es", translator)
}

func testDataSynchronization(lang string, translator *Translator) {
	// Synchronizing date
	date := "2023-07-31"
	expectedDateFormat := translator.Translate("date_format", lang)
	actualDate := translator.FormatMessage("date_format", lang, date)
	fmt.Printf("Testing date format in %s: Expected '%s', Got '%s'\n", lang, expectedDateFormat, actualDate)

	// Synchronizing currency
	amount := 123.45
	expectedCurrencyFormat := translator.Translate("currency_format", lang)
	actualCurrency := translator.FormatMessage("currency_format", lang, amount)
	fmt.Printf("Testing currency format in %s: Expected '%s', Got '%s'\n", lang, expectedCurrencyFormat, actualCurrency)
}
