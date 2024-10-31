package main

import (
	"fmt"

	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

// SynchronizedData holds the data that needs to be synchronized
type SynchronizedData struct {
	WelcomeMessage string
	GoodbyeMessage string
}

var syncedData SynchronizedData

func loadTranslations(language string) error {
	cat, err := catalog.FromFile("translations/" + language + ".txt")
	if err != nil {
		return err
	}

	p := message.NewPrinter(language)
	syncedData.WelcomeMessage = p.Sprintf(cat.Message("welcome"))
	syncedData.GoodbyeMessage = p.Sprintf(cat.Message("goodbye"))

	return nil
}

func main() {
	// Choose the language you want to use (e.g., "en", "es", "fr")
	language := "en"

	// Load the translations for the chosen language
	if err := loadTranslations(language); err != nil {
		fmt.Println("Error loading translations:", err)
		return
	}

	// Now you can access the synchronized data in the desired language
	fmt.Println(syncedData.WelcomeMessage)
	fmt.Println(syncedData.GoodbyeMessage)
}
