package main

import (
	"fmt"
	"reflect"
)

// Define a common interface for all content types
type Content interface {
	Display() string
}

// Different content types implementing the Content interface
type TextContent struct {
	Text string
}

func (t TextContent) Display() string {
	return fmt.Sprintf("Text: %s", t.Text)
}

type ImageContent struct {
	URL string
}

func (i ImageContent) Display() string {
	return fmt.Sprintf("Image: %s", i.URL)
}

type VideoContent struct {
	Title    string
	Duration int
}

func (v VideoContent) Display() string {
	return fmt.Sprintf("Video: %s, Duration: %d seconds", v.Title, v.Duration)
}

// Function to handle displaying content using reflection
func DisplayContent(content interface{}) {
	value := reflect.ValueOf(content)
	if value.Kind() == reflect.Ptr && !value.IsNil() {
		value = value.Elem()
	}

	if value.Kind() == reflect.Struct {
		method := value.MethodByName("Display")
		if method.IsValid() {
			results := method.Call(nil)
			if len(results) > 0 && results[0].Kind() == reflect.String {
				fmt.Println(results[0].String())
				return
			}
		}
	}

	fmt.Println("Unsupported content type")
}

func main() {
	textContent := &TextContent{Text: "Hello, World!"}
	imageContent := &ImageContent{URL: "example.com/image.jpg"}
	videoContent := &VideoContent{Title: "Sample Video", Duration: 120}

	DisplayContent(textContent)
	DisplayContent(imageContent)
	DisplayContent(videoContent)

	// Unsupported type
	DisplayContent(10)
}
