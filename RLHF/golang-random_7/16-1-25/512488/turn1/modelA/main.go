package main

import (
	"fmt"
)

// Define a struct representing your data
type PageData struct {
	Title   string
	Header  string
	Content string
	Footer  string
	Links   []string // Example of an array for multiple links
}

// Function to generate the HTML page
func generateHTML(data PageData) string {
	// Define the HTML template
	htmlTemplate := `<!DOCTYPE html>
<html>
<head>
	<title>%s</title>
</head>
<body>
	<header>%s</header>
	<main>%s</main>
	<footer>%s</footer>
	<ul>
		%s
	</ul>
</body>
</html>`

	// Create a list item string from the links array
	linkItems := ""
	for _, link := range data.Links {
		linkItems += fmt.Sprintf("<li><a href=\"%s\">%s</a></li>\n", link, link)
	}

	// Generate the HTML content using fmt.Sprintf
	htmlContent := fmt.Sprintf(htmlTemplate, data.Title, data.Header, data.Content, data.Footer, linkItems)

	return htmlContent
}

func main() {
	// Create an instance of the struct with your data
	pageData := PageData{
		Title:   "My Web Page",
		Header:  "<h1>Welcome!</h1>",
		Content: "<p>This is the content of my page.</p>",
		Footer:  "<p>Copyright &copy; 2023</p>",
		Links:   []string{"https://example.com", "https://example.org", "https://example.net"},
	}

	// Generate the HTML
	htmlOutput := generateHTML(pageData)

	// Print the HTML output
	fmt.Println(htmlOutput)
}
