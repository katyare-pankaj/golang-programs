package main

import (
	"fmt"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	// Set the French locale
	fr := language.French
	p := message.NewPrinter(fr)

	// Format the current date in French
	now := time.Now()
	dateFormat := "02 January 2006"
	dateStr := p.Sprintf(dateFormat, now.Day(), now.Month(), now.Year())

	fmt.Println("Date in French:", dateStr)
}
