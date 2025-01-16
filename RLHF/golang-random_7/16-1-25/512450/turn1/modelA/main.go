package main

import (
	"fmt"
	"strings"
)

// DSL functions
func html(tag string, args ...interface{}) string {
	return fmt.Sprintf("<%s>%s</%s>", tag, strings.Join(map[string]string{
		"%v": fmt.Sprint(args...),
	}["%v"], ""), tag)
}

func head() string {
	return html("head", html("title", "My Website"))
}

func body() string {
	return html("body",
		html("h1", "Welcome!"),
		html("p", "This is a simple web page using a DSL in Go."),
	)
}

func main() {
	// Generate the HTML template using the DSL
	htmlContent := fmt.Sprintf("<!DOCTYPE html>\n%s\n%s", head(), body())

	// Output the generated HTML
	fmt.Println(htmlContent)
}
