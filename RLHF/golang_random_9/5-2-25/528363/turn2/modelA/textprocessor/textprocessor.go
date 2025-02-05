package textprocessor

// TextProcessor defines the interface for text processing plugins.
type TextProcessor interface {
	Process(text string) (string, error)
}
