package main

import (
	"fmt"
	"strings"
)

// Widget represents a part of the dashboard
type Widget struct {
	Title string
	Data  string
}

// Dashboard represents the entire dashboard layout
type Dashboard struct {
	Widgets []Widget
}

// AddWidget adds a new widget to the dashboard
func (d *Dashboard) AddWidget(title, data string) {
	d.Widgets = append(d.Widgets, Widget{Title: title, Data: data})
}

// ModifyWidget modifies an existing widget based on index
func (d *Dashboard) ModifyWidget(index int, title, data string) error {
	if index < 0 || index >= len(d.Widgets) {
		return fmt.Errorf("widget index out of range")
	}
	d.Widgets[index] = Widget{Title: title, Data: data}
	return nil
}

// Render generates a string representation of the dashboard
func (d *Dashboard) Render() string {
	var sb strings.Builder
	sb.WriteString("Dashboard\n")
	sb.WriteString(strings.Repeat("=", 10) + "\n")
	for i, widget := range d.Widgets {
		sb.WriteString(fmt.Sprintf("Widget %d: %s\n", i+1, widget.Title))
		sb.WriteString(strings.Repeat("-", len(widget.Title)+10) + "\n")
		sb.WriteString(fmt.Sprintf("%s\n\n", widget.Data))
	}
	return sb.String()
}

func main() {
	// Create a new dashboard
	dashboard := &Dashboard{}

	// Add widgets to the dashboard
	dashboard.AddWidget("Sales Figures", "Q1: $100K\nQ2: $150K")
	dashboard.AddWidget("User Feedback", "Positive: 80%\nNegative: 10%")

	// Modify a widget
	err := dashboard.ModifyWidget(1, "Updated User Feedback", "Positive: 85%\nNegative: 5%\nNeutral: 10%")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Render and display the dashboard
	fmt.Println(dashboard.Render())
}
