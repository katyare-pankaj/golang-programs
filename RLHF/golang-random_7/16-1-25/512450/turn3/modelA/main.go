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
}

func (h *Html) AddClass(classes ...string) {
	if _, ok := h.Attrs["class"]; !ok {
		h.Attrs["class"] = ""
	}
	existingClasses := strings.Fields(h.Attrs["class"])
	for _, class := range classes {
		if !strings.Contains(h.Attrs["class"], class) {
			existingClasses = append(existingClasses, class)
		}
	}
	h.Attrs["class"] = strings.Join(existingClasses, " ")
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
			Content: "Hello, World!",
		}.AddClass("title", "main-heading"),
		Html{
			Tag:     "p",
			Content: "This is a simple GoWeb DSL example with dynamic CSS classes.",
		}.AddClass("intro"),
		Text{
			Content: "Using fmt.Sprintf for dynamic template generation.",
		},
	)

	// Render the DSL and print the result
	fmt.Println(gw.Render())
}
