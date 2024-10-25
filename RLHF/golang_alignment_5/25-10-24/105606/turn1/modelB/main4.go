package main

import (
	"fmt"
	"strings"

	"github.com/aymerick/raymond"
)

func main() {
	logMessage := "The server encountered an error: Internal Server Error (500)"
	// Simple NLP using template matching
	template := "The server encountered an error: {{error_message}} ({{error_code}})"
	rendered, _ := raymond.Render(template, map[string]interface{}{
		"error_message": strings.Split(logMessage, ": ")[1],
		"error_code":    strings.Split(logMessage, " (")[1][:3],
	})
	fmt.Println("Extracted Information:", rendered)
}
