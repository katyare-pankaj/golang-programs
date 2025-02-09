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
	
	Type      EventType
	Index     int
	Value     interface{}
	IsCheckpoint bool
}

// SliceStore holds the slice and its event history.
type SliceStore struct {
	slice  []interface{}
	events []Event
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

// MarkCheckpoint adds a marker in the events to declare a checkpoint state.
func (s *SliceStore) MarkCheckpoint() {
	s.events = append(s.events, Event{Type: -1, IsCheckpoint: true})
}

// GetSlice returns the current state of the slice.
func (s *SliceStore) GetSlice() []interface{} {
	return append([]interface{}{}, s.slice...)
}

// Restore resets the slice to the state after the specified number of events.
func (s *SliceStore) Restore(eventCount int) error {
	if eventCount < 0 || eventCount > len(s.events) {
		return errors.New("invalid event count")
	}

	// Rebuild the slice by reapplying events.
	s.slice = make([]interface{}, 0)
	for i := 0; i < eventCount; i++ {
		event := s.events[i]
		switch event.Type {
		case Add:
			s.slice = append(s.slice, event.Value)
		case Remove:
			s.slice = append(s.slice[:event.Index], s.slice[event.Index+1:]...)
		case Update:
			if event.Index < len(s.slice) {
				s.slice[event.Index] = event.Value
			}
		}
	}
	return nil
}

// RestoreToCheckpoint finds the latest checkpoint and restores to that state.
func (s *SliceStore) RestoreToCheckpoint() error {
	lastCheckpoint := -1

	for i, event := range s.events {
		if event.IsCheckpoint {
			lastCheckpoint = i
		}
	}

	if lastCheckpoint < 0 {
		return errors.New("no checkpoints available to restore to")
	}

	return s.Restore(lastCheckpoint + 1) // Restore to state right after the checkpoint
}

func main() {
	store := NewSliceStore()

	// Add elements to the slice.
	store.Add("A")
	store.Add("B")
	store.MarkCheckpoint() // Mark checkpoint
	store.Add("C")

	fmt.Println("Current Slice:", store.GetSlice())

	// Update an element in the slice.
	store.Update(1, "B-updated")
	fmt.Println("Updated Slice:", store.GetSlice())

	// Restore to the latest checkpoint.
	store.RestoreToCheckpoint()
	fmt.Println("Restored to Checkpoint Slice:", store.GetSlice())
}