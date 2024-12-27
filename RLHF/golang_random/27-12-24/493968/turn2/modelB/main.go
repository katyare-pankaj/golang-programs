package main

import (
	"errors"
	"fmt"
)

type LoginEvent struct {
	Username string
	Password string
}

// UserService handles user-related operations.
type UserService interface {
	Login(event LoginEvent) error
}

// In-memory user service for demonstration purposes.
type InMemoryUserService struct {
	users map[string]string
}

func NewInMemoryUserService() *InMemoryUserService {
	return &InMemoryUserService{users: make(map[string]string)}
}

func (s *InMemoryUserService) Login(event LoginEvent) error {
	expectedPassword, ok := s.users[event.Username]
	if !ok {
		return errors.New("user not found")
	}

	if expectedPassword != event.Password {
		return errors.New("invalid password")
	}

	return nil
}

// EventProcessor processes login events.
type EventProcessor struct {
	userService UserService
}

func NewEventProcessor(userService UserService) *EventProcessor {
	return &EventProcessor{userService: userService}
}

func (p *EventProcessor) ProcessLoginEvent(event LoginEvent) error {
	if err := p.validateLoginEvent(event); err != nil {
		return fmt.Errorf("event validation failed: %w", err)
	}

	err := p.userService.Login(event)
	if err != nil {
		return fmt.Errorf("user login failed: %w", err)
	}

	fmt.Println("User logged in successfully!")
	return nil
}

func (p *EventProcessor) validateLoginEvent(event LoginEvent) error {
	if event.Username == "" {
		return errors.New("username is required")
	}

	if event.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func main() {
	userService := NewInMemoryUserService()
	eventProcessor := NewEventProcessor(userService)

	loginEvent := LoginEvent{Username: "john", Password: "password"}
	if err := eventProcessor.ProcessLoginEvent(loginEvent); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
