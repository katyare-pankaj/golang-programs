package main

import (
	"errors"
	"fmt"
)

// EventType defines the types of events that affect the slice.
type EventType int

const (
	// Event types for the slice operations.
	Add EventType = iota
	Remove
	Update
)

// Event represents a change made to the slice.
type Event struct {
	Type  EventType   // The type of operation: Add, Remove, Update
	Index int         // Index of the element affected by the operation
	Value interface{} // Value of the element involved in the operation
}

// SliceStore holds the slice and its event history.
type SliceStore struct {
	slice  []interface{} // The actual slice being managed
	events []Event       // The log of all events affecting the slice
}

// NewSliceStore creates and returns a new SliceStore.
func NewSliceStore() *SliceStore {
	return &SliceStore{
		slice:  make([]interface{}, 0),
		events: make([]Event, 0),
	}
}

// Add appends a new item to the slice.
func (s *SliceStore) Add(value interface{}) {
	s.events = append(s.events, Event{Type: Add, Index: len(s.slice), Value: value})
	s.slice = append(s.slice, value)
}

// Remove deletes an item from the slice at the specified index.
func (s *SliceStore) Remove(index int) error {
	if index < 0 || index >= len(s.slice) {
		return errors.New("index out of range")
	}
	s.events = append(s.events, Event{Type: Remove, Index: index, Value: s.slice[index]})
	s.slice = append(s.slice[:index], s.slice[index+1:]...)
	return nil
}

// Update changes the value of the slice at the specified index.
func (s *SliceStore) Update(index int, value interface{}) error {
	if index < 0 || index >= len(s.slice) {
		return errors.New("index out of range")
	}
	s.events = append(s.events, Event{Type: Update, Index: index, Value: value})
	s.slice[index] = value
	return nil
}

// GetSlice returns the current state of the slice.
func (s *SliceStore) GetSlice() []interface{} {
	return append([]interface{}{}, s.slice...)
}

// Restore reverts the slice to reflect the state after a specified number of events.
func (s *SliceStore) Restore(eventCount int) error {
	if eventCount < 0 || eventCount > len(s.events) {
		return errors.New("invalid event count")
	}

	// Re-initialize the slice and reapply events.
	s.slice = make([]interface{}, 0)
	for i := 0; i < eventCount; i++ {
		event := s.events[i]
		switch event.Type {
		case Add:
			s.slice = append(s.slice, event.Value)
		case Remove:
			// Every Remove event removes an element, so we don't apply it directly here.
			// Instead, we don't add the removed element to current scan of the.
		case Update:
			s.slice[event.Index] = event.Value
		}
	}
	// reapply remove logic to finalize state as events[idx].Index as index is not correct.
	for i := 0; i < eventCount; i++ {
		event := s.events[i]
		if event.Type == Remove {
			s.slice = append(s.slice[:event.Index], s.slice[event.Index+1:]...)
		}
	}
	return nil
}

func main() {
	store := NewSliceStore()

	// Add elements to the slice.
	store.Add("A")
	store.Add("B")
	store.Add("C")

	fmt.Println("Current Slice:", store.GetSlice())

	// Update an element in the slice.
	store.Update(1, "B-updated")
	fmt.Println("Updated Slice:", store.GetSlice())

	// Remove an element from the slice.
	store.Remove(0)
	fmt.Println("Modified Slice:", store.GetSlice())

	// Restore to an earlier state.
	store.Restore(2)
	fmt.Println("Restored Slice (after 2 events):", store.GetSlice())
}
