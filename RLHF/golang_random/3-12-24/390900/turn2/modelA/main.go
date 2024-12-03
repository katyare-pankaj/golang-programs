package main

import (
	"fmt"
)

// Define a simple interface for logging
type Logger interface {
	Log(message string)
}

// Concrete implementation of the Logger interface
type ConsoleLogger struct{}

func (l *ConsoleLogger) Log(message string) {
	fmt.Println("Console:", message)
}

// Define a separate interface for saving data
type Saver interface {
	Save(data interface{}) error
}

// Concrete implementation of the Saver interface
type FileSaver struct{}

func (s *FileSaver) Save(data interface{}) error {
	fmt.Println("File: Saving data...")
	return nil
}

// Use interface composition to combine Logger and Saver functionalities
type LogAndSave struct {
	Logger
	Saver
}

// Example function that uses the LogAndSave struct
func useLogAndSave(las LogAndSave) {
	las.Log("Starting operation...")
	err := las.Save("Sample data")
	if err != nil {
		las.Log("Error saving data:", err)
	} else {
		las.Log("Data saved successfully.")
	}
}

func main() {
	// Create a LogAndSave instance using composition
	las := LogAndSave{
		Logger: &ConsoleLogger{},
		Saver:  &FileSaver{},
	}

	// Use the LogAndSave instance
	useLogAndSave(las)
}
