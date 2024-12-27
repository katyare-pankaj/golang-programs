package main

import (
	"fmt"
)

type EventHandler interface {
	HandleEvent(event interface{}) (success bool, message string, err error)
}

type SomeEvent struct {
	Data string
}

type MyEventHandler struct{}

func (h *MyEventHandler) HandleEvent(event interface{}) (success bool, message string, err error) {
	switch ev := event.(type) {
	case *SomeEvent:
		fmt.Println("Handling SomeEvent:", ev.Data)
		return true, "Event handled successfully", nil
	default:
		return false, "Unsupported event type", fmt.Errorf("unsupported event type: %T", event)
	}
}

func dispatchEvent(handler EventHandler, event interface{}) {
	success, message, err := handler.HandleEvent(event)
	if err != nil {
		fmt.Println("Error handling event:", err)
	} else if !success {
		fmt.Println("Event handling failed:", message)
	} else {
		fmt.Println("Event handled successfully:", message)
	}
}

func main() {
	handler := &MyEventHandler{}
	dispatchEvent(handler, &SomeEvent{Data: "Hello"})
	dispatchEvent(handler, 42) // Unsupported event type
}
