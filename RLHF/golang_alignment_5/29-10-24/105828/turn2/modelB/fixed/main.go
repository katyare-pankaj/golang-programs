package main

import (
	"fmt"
	"reflect"
)

// Define a interface for content with similar operations
type Content interface {
	GetTitle() string
	SetTitle(string)
}

// Structs representing different types of content
type BlogPost struct {
	title string
}

func (b *BlogPost) GetTitle() string {
	return b.title
}

func (b *BlogPost) SetTitle(title string) {
	b.title = title
}

type Video struct {
	title string
}

func (v *Video) GetTitle() string {
	return v.title
}

func (v *Video) SetTitle(title string) {
	v.title = title
}

// Function to handle operation on any content type that implements the Content interface
func handleContent(content interface{}, operation string, args ...interface{}) {
	v := reflect.ValueOf(content)

	// Check if the content type implements the Content interface
	if !reflect.TypeOf(content).Implements(reflect.TypeOf((*Content)(nil)).Elem()) {
		fmt.Println("Error: Invalid content type.")
		return
	}

	// Get the method corresponding to the operation
	method := v.MethodByName(operation)
	if !method.IsValid() {
		fmt.Println("Error: Invalid operation.")
		return
	}

	// Convert arguments to reflect.Value and call the method
	argValues := make([]reflect.Value, len(args))
	for i, arg := range args {
		argValues[i] = reflect.ValueOf(arg)
	}
	result := method.Call(argValues)

	// Print the result (if the method returns a value)
	if len(result) > 0 {
		fmt.Println("Result:", result[0].Interface())
	}
}

func main() {
	post := &BlogPost{title: "My First Post"}
	video := &Video{title: "Introduction to Go"}

	handleContent(post, "GetTitle") // Output: Result: My First Post
	handleContent(post, "SetTitle", "New Title")
	handleContent(post, "GetTitle") // Output: Result: New Title

	handleContent(video, "GetTitle") // Output: Result: Introduction to Go
	handleContent(video, "SetTitle", "Updated Video Title")
	handleContent(video, "GetTitle") // Output: Result: Updated Video Title
}
