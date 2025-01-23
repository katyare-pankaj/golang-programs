package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
)

type Dashboard struct {
	Title   string
	Widgets []Widget
}

type Widget struct {
	Title   string
	Content string
}

func main() {
	// Create a dashboard
	dashboard := &Dashboard{
		Title: "My Dashboard",
		Widgets: []Widget{
			{Title: "Widget 1", Content: "This is widget 1 content"},
			{Title: "Widget 2", Content: "This is widget 2 content"},
		},
	}

	// Add a new widget to the dashboard
	dashboard.Widgets = append(dashboard.Widgets, Widget{Title: "Widget 3", Content: "This is widget 3 content"})

	// Modify the content of a widget
	dashboard.Widgets[0].Content = "Updated content for Widget 1"

	// Create a template for the dashboard
	t, err := template.New("dashboard").Parse(`
        <html>
            <head>
                <title>{{ .Title }}</title>
            </head>
            <body>
                <h1>{{ .Title }}</h1>
                {{ range .Widgets }}
                    <h3>{{ .Title }}</h3>
                    <p>{{ .Content }}</p>
                {{ end }}
            </body>
        </html>
    `)
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Execute the template and write the output to a file
	f, err := os.Create("dashboard.html")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer f.Close()

	if err := t.Execute(f, dashboard); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	fmt.Println("Dashboard rendered successfully!")
}
