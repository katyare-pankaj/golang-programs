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

// Add Attr method to add dynamic attributes to Html
func (h *Html) Attr(key string, value interface{}) *Html {
	if h.Attrs == nil {
		h.Attrs = make(map[string]string)
	}
	h.Attrs[key] = fmt.Sprintf("%v", value)
	return h
}

// The main function starts here
func main() {
	gw := GoWeb{}

	gw.Add(
		Html{
			Tag:     "h1",
			Content: "Hello, World!",
		}.Attr("class", "title").Attr("id", "main-title"), // Add dynamic attributes here
		Html{
			Tag:     "p",
			Content: "This is a simple GoWeb DSL example.",
		},
		Text{
			Content: "Using fmt.Sprintf for dynamic template generation.",
		},
	)

	fmt.Println(gw.Render())
}
