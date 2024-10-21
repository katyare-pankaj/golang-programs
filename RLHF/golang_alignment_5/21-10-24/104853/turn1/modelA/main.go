package main

import (
	"fmt"
	"time"
)

// Meeting struct to hold meeting details
type Meeting struct {
	Date          time.Time
	Agenda        string
	Participants  []string
	Notes         string
	ActionItems   []string
	CodeSnippets  []string // To store code snippets for discussion
	ReviewedFiles []string // List of files reviewed
	Suggestions   []string // General suggestions or improvements
	Concerns      []string // Any issues or concerns raised
}

// CodeReview holds details of a code review
type CodeReview struct {
	Meeting
	Reviewer string
	Reviewee string
}

func conductCodeReview(review CodeReview) {
	fmt.Printf("\nCode Review between %s and %s\n", review.Reviewer, review.Reviewee)
	// Simulate code review discussion
	review.AddNotes("Discussed code structure and formatting.")
	review.AddActionItem("Refactor function to use better variable names.")
	review.CodeSnippets = append(review.CodeSnippets, `func calculateSomething() int { ... }`)
	review.Suggestions = append(review.Suggestions, "Add comments to explain the logic.")

	// Display the review details
	review.DisplayMeeting()
}

func main() {
	reviewer := "Alice"
	reviewee := "Bob"
	date := time.Now()
	agenda := "Code Review of updated functionality"

	review := CodeReview{
		Meeting:  NewMeeting(date, agenda, []string{reviewer, reviewee}),
		Reviewer: reviewer,
		Reviewee: reviewee,
	}

	conductCodeReview(review)
}

// NewMeeting creates a new meeting instance
func NewMeeting(date time.Time, agenda string, participants []string) Meeting {
	return Meeting{
		Date:         date,
		Agenda:       agenda,
		Participants: participants,
		Notes:        "",
		ActionItems:  []string{},
	}
}

// AddNotes adds meeting notes
func (m *Meeting) AddNotes(notes string) {
	m.Notes = notes
}

// AddActionItem adds an action item to the meeting
func (m *Meeting) AddActionItem(actionItem string) {
	m.ActionItems = append(m.ActionItems, actionItem)
}

// DisplayMeeting displays meeting information
func (m Meeting) DisplayMeeting() {
	fmt.Printf("Meeting Date: %s\n", m.Date.Format("2006-01-02 15:04"))
	fmt.Printf("Agenda: %s\n", m.Agenda)
	fmt.Printf("Participants: %s\n", m.Participants)
	fmt.Printf("Notes: %s\n", m.Notes)
	fmt.Printf("Action Items: %s\n", m.ActionItems)
}
