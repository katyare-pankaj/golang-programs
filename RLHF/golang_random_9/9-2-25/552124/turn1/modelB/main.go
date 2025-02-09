package main

import (
	"fmt"
	"sync"
)

// Event represents a single change in the slice
type Event struct {
	Type  string      // Type of event (e.g., "append", "remove", "update")
	Index int         // Index of the element being affected
	Data  interface{} // The data of the element
}

// EventStore is an interface for storing events
type EventStore interface {
	AppendEvent(*Event) error
	GetEvents() ([]*Event, error)
}

// InMemoryEventStore is an in-memory implementation of EventStore
type InMemoryEventStore struct {
	events []*Event
	mu     sync.Mutex
}

func (s *InMemoryEventStore) AppendEvent(event *Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.events = append(s.events, event)
	return nil
}

func (s *InMemoryEventStore) GetEvents() ([]*Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.events, nil
}

// SliceStore manages the slice using events
type SliceStore struct {
	store     EventStore
	slice     []interface{}
	lastIndex int
	mu        sync.Mutex
}

// NewSliceStore creates a new SliceStore
func NewSliceStore(store EventStore) *SliceStore {
	return &SliceStore{
		store: store,
		slice: make([]interface{}, 0),
	}
}

// Append appends an element to the slice and stores an event
func (s *SliceStore) Append(data interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	event := &Event{
		Type:  "append",
		Index: s.lastIndex,
		Data:  data,
	}

	if err := s.store.AppendEvent(event); err != nil {
		return err
	}

	s.slice = append(s.slice, data)
	s.lastIndex++

	return nil
}

// Remove removes an element from the slice and stores an event
func (s *SliceStore) Remove(index int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.slice) {
		return fmt.Errorf("index %d is out of bounds", index)
	}

	event := &Event{
		Type:  "remove",
		Index: index,
		Data:  s.slice[index],
	}

	if err := s.store.AppendEvent(event); err != nil {
		return err
	}

	s.slice = append(s.slice[:index], s.slice[index+1:]...)

	return nil
}

// Update updates an element in the slice and stores an event
func (s *SliceStore) Update(index int, data interface{}) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if index < 0 || index >= len(s.slice) {
		return fmt.Errorf("index %d is out of bounds", index)
	}

	event := &Event{
		Type:  "update",
		Index: index,
		Data:  data,
	}

	if err := s.store.AppendEvent(event); err != nil {
		return err
	}

	s.slice[index] = data

	return nil
}

// GetSlice returns the current state of the slice
func (s *SliceStore) GetSlice() []interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	// Return a copy to avoid race conditions
	return append([]interface{}{}, s.slice...)
}

// RestoreFromEvents restores the slice from a sequence of events
