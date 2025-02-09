package main

import (
	"errors"
	"fmt"
)

// Existing code remains the same...

// RestoreToSpecificState reverts the slice to reflect the state after applying the provided sequence of events.
func (s *SliceStore) RestoreToSpecificState(events []Event) error {
	// Re-initialize the slice and reapply events.
	s.slice = make([]interface{}, 0)
	for _, event := range events {
		switch event.Type {
		case Add:
			s.slice = append(s.slice, event.Value)
		case Remove:
			// Every Remove event removes an element, so we don't apply it directly here.
			// Instead, we don't add the removed element to the current scan of the slice.
		case Update:
			if event.Index < 0 || event.Index >= len(s.slice) {
				return fmt.Errorf("index %d out of range during restoration", event.Index)
			}
			s.slice[event.Index] = event.Value
		default:
			return errors.New("unknown event type")
		}
	}

	// reapply remove logic to finalize the state.
	for _, event := range events {
		if event.Type == Remove {
			if event.Index < 0 || event.Index >= len(s.slice) {
				return fmt.Errorf("index %d out of range during removal restoration", event.Index)
			}
			s.slice = append(s.slice[:event.Index], s.slice[event.Index+1:]...)
		}
	}

	return nil
}

func main() {
	// Existing code remains the same...

	// Restore the slice to a specific state
	events := []Event{
		{Type: Update, Index: 0, Value: "A-updated"}, // Change A to A-updated
		{Type: Add, Index: 2, Value: "D"},            // Add D after C
	}
	if err := store.RestoreToSpecificState(events); err != nil {
		fmt.Println("Restoration error:", err)
	} else {
		fmt.Println("Restored Slice to specific state:", store.GetSlice())
	}
}
