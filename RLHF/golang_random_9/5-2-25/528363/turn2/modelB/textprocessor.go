// textprocessor.go
package main

// TextProcessor defines the interface for plugins to implement
type TextProcessor interface {
	Process(text string) (string, error)
}
