package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// ConsentRepository represents an interface for managing consents
type ConsentRepository interface {
	Save(consent *UserConsent) error
	GetAll() ([]*UserConsent, error)
}

// UserConsent represents a user's consent for specific data processing activities
type UserConsent struct {
	UserID           string     `json:"user_id"`
	ActivityID       string     `json:"activity_id"`
	Description      string     `json:"description"`
	GivenAt          time.Time  `json:"given_at"`
	RevokedAt        *time.Time `json:"revoked_at,omitempty"`
	IsConsentGranted bool       `json:"is_consent_granted"`
}

// InMemoryConsentRepository is a simple in-memory implementation of ConsentRepository
type InMemoryConsentRepository struct {
	consents []*UserConsent
}

func (repo *InMemoryConsentRepository) Save(consent *UserConsent) error {
	consent.GivenAt = time.Now()
	repo.consents = append(repo.consents, consent)
	return nil
}

func (repo *InMemoryConsentRepository) GetAll() ([]*UserConsent, error) {
	return repo.consents, nil
}

func main() {
	repo := &InMemoryConsentRepository{}

	// Save consent for processing user location data
	consent1 := &UserConsent{
		UserID:           "user123",
		ActivityID:       "process_location_data",
		Description:      "Allow processing of user location data for personalized experiences.",
		IsConsentGranted: true,
	}

	if err := repo.Save(consent1); err != nil {
		fmt.Println("Error saving consent:", err)
	} else {
		fmt.Println("Consent saved successfully.")
	}

	// Get all consents for user123
	consents, err := repo.GetAll()
	if err != nil {
		fmt.Println("Error retrieving consents:", err)
	} else {
		fmt.Println("Consents for user123:")
		for _, consent := range consents {
			consentJSON, _ := json.MarshalIndent(consent, "", "  ")
			fmt.Println(string(consentJSON))
		}
	}
}
