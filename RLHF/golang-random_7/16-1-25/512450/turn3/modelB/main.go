package main

import (
	"fmt"
	"strings"
)

// The DSL starts here

type GoWeb struct {
	Components []Component
}

type Component interface {
	Render() string
}

// Define your DSL components

type Html struct {
	Tag     string
	Attrs   map[string]string
	Content string
	Classes []string
}

func (h Html) Render() string {
	attrs := ""
	for key, value := range h.Attrs {
		attrs += fmt.Sprintf(" %s=\"%s\"", key, value)
	}

	// Add classes attribute if there are any classes defined
	if len(h.Classes) > 0 {
		classes := strings.Join(h.Classes, " ")
		attrs += fmt.Sprintf(" class=\"%s\"", classes)
	}

	return fmt.Sprintf("<%s%s>%s</%s>", h.Tag, attrs, h.Content, h.Tag)
}

type Text struct {
	Content string
}

func (t Text) Render() string {
	return t.Content
}

// End of DSL components

func (g *GoWeb) Add(components ...Component) {
	g.Components = append(g.Components, components...)
}

func (g *GoWeb) Render() string {
	var result string
	for _, component := range g.Components {
		result += component.Render()
	}
	return result
}

// The main function starts here

func main() {
	// Create an instance of the GoWeb DSL
	gw := GoWeb{}

	// Define your components using the DSL
	gw.Add(
		Html{
			Tag:     "h1",
			Content: "Hello, World!",
			Classes: []string{"primary", "heading"},
		},
		Html{
			Tag:     "p",
			Content: "This is a simple GoWeb DSL example.",
			Classes: []string{"lead"},
		},
		Text{
			Content: "Using fmt.Sprintf for dynamic template generation.",
		},
	)

	// Render the DSL and print the result
	fmt.Println(gw.Render())
}
