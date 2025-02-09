package main  
import (  
    "fmt"
)

type EventType int  
const (  
    Append EventType = iota
    Insert
    Remove
    Update
    Clear
)

type Event struct {  
    Type EventType
    Index int
    Value interface{}
    OldValue interface{} // For Update event
}

type SliceStore struct {  
    slice []interface{}
    events []Event
    undoIndex int
}

func NewSliceStore() *SliceStore {  
    return &SliceStore{  
        slice: make([]interface{}, 0),
        events: make([]Event, 0),
        undoIndex: -1,
    }
}

func (s *SliceStore) Append(value interface{}) {  
    s.events = append(s.events, Event{Type: Append, Index: len(s.slice), Value: value})
    s.slice = append(s.slice, value)
    s.undoIndex = len(s.events) - 1
}

func (s *SliceStore) Insert(index int, value interface{}) {  
    if index < 0 || index > len(s.slice) {
        fmt.Println("Index out of range")
        return
    }
    s.events = append(s.events, Event{Type: Insert, Index: index, Value: value})
    s.slice = append(s.slice[:index], append([]interface{}{value}, s.slice[index:]...)...)
    s.undoIndex = len(s.events) - 1
}

func (s *SliceStore) Remove(index int) {  
    if index < 0 || index >= len(s.slice) {
        fmt.Println("Index out of range")
        return
    }
    oldValue := s.slice[index]
    s.events = append(s.events, Event{Type: Remove, Index: index, Value: oldValue})
    s.slice = append(s.slice[:index], s.slice[index+1:]...)
    s.undoIndex = len(s.events) - 1
}

func (s *SliceStore) Update(index int, value interface{}) {  
    if index < 0 || index >= len(s.slice) {
        fmt.Println("Index out of range")
        return
    }
    oldValue := s.slice[index]
    s.events = append(s.events, Event{Type: Update, Index: index, Value: value, OldValue: oldValue})
    s.slice[index] = value
    s.undoIndex = len(s.events) - 1
}

func (s *SliceStore) Clear() {  
    if len(s.slice) > 0 {
        s.events = append(s.events, Event{Type: Clear})
        s.slice = make([]interface{}, 0)
        s.undoIndex = len(s.events) - 1
    }
}

func (s *SliceStore) Get(index int) interface{} {  
    if index < 0 || index >= len(s.slice) {
        fmt.Println("Index out of range")
        return nil
    }
    return s.slice[index]
}

func (s *SliceStore) Undo() {  
    if s.undoIndex >= 0 {
        event := s.events[s.undoIndex]
        s.undoIndex--
        switch event.Type {
        case Append:
            s.slice = s.slice[:len(s.slice)-1]
        case Insert:
            s.slice = append(s.slice[:event.Index], s.slice[event.Index+1:]...)
        case Remove:
            s.slice = append(s.slice[:event.Index], event.Value)
        case Update:
            s.slice[event.Index] = event.OldValue