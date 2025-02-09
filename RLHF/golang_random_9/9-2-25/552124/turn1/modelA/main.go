package main

import (
	"fmt"
)

// EventType defines the types of events that can occur in the slice store.
type EventType int

const (
	Add EventType = iota
	Remove
	Update
)

// Event represents a change to the slice.
type Event struct {
	Type  EventType
	Index int
	Value interface{}
}

// SliceStore manages a slice and its event log.
type SliceStore struct {
	slice  []interface{}
	events []Event
}

// NewSliceStore creates a new SliceStore.
func NewSliceStore() *SliceStore {
	return &SliceStore{
		slice:  make([]interface{}, 0),
		events: make([]Event, 0),
	}
}

// Add appends a new element to the slice.
func (s *SliceStore) Add(value interface{}) {
	s.events = append(s.events, Event{Type: Add, Index: len(s.slice), Value: value})
	s.slice = append(s.slice, value)
}

// Remove deletes the element at the specified index.
func (s *SliceStore) Remove(index int) {
	if index < 0 || index >= len(s.slice) {
		fmt.Println("Index out of range")
		return
	}
	s.events = append(s.events, Event{Type: Remove, Index: index, Value: s.slice[index]})
	s.slice = append(s.slice[:index], s.slice[index+1:]...)
}

// Update modifies the element at the specified index.
func (s *SliceStore) Update(index int, value interface{}) {
	if index < 0 || index >= len(s.slice) {
		fmt.Println("Index out of range")
		return
	}
	s.events = append(s.events, Event{Type: Update, Index: index, Value: value})
	s.slice[index] = value
}

// GetSlice returns a copy of the current slice.
func (s *SliceStore) GetSlice() []interface{} {
	return append([]interface{}{}, s.slice...)
}

// Restore resets the slice to the state after the specified number of events.
func (s *SliceStore) Restore(eventCount int) {
	if eventCount < 0 || eventCount > len(s.events) {
		fmt.Println("Invalid event count")
		return
	}

	s.slice = make([]interface{}, 0)
	for i := 0; i < eventCount; i++ {
		event := s.events[i]
		switch event.Type {
		case Add:
			s.slice = append(s.slice, event.Value)
		case Remove:
			s.slice = append(s.slice[:event.Index], s.slice[event.Index+1:]...)
		case Update:
			s.slice[event.Index] = event.Value
		}
	}
}

func main() {
	store := NewSliceStore()

	// Add elements
	store.Add("A")
	store.Add("B")
	store.Add("C")

	fmt.Println("Current Slice:", store.GetSlice())

	// Update an element
	store.Update(1, "B-updated")
	fmt.Println("Updated Slice:", store.GetSlice())

	// Remove an element
	store.Remove(0)
	fmt.Println("Modified Slice:", store.GetSlice())

	// Restore to an earlier state
	store.Restore(2)
	fmt.Println("Restored Slice:", store.GetSlice())
}
