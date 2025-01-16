package main

import (
	"fmt"
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
}

func (h Html) Render() string {
	attrs := ""
	for key, value := range h.Attrs {
		attrs += fmt.Sprintf(" %s=\"%s\"", key, value)
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
			Attrs:   map[string]string{"class": "title", "id": "main-heading"},
			Content: "Hello, World!",
		},
		Html{
			Tag:     "p",
			Attrs:   map[string]string{"class": "intro"},
			Content: "This is a simple GoWeb DSL example with dynamic attributes.",
		},
		Text{
			Content: "Using fmt.Sprintf for dynamic template generation.",
		},
	)

	// Render the DSL and print the result
	fmt.Println(gw.Render())
}
