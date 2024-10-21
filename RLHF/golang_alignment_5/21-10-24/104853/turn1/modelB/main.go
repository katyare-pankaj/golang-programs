package main

import "time"

// Meeting struct to hold meeting details.
type Meeting struct {
	Date         time.Time // Use time.Time for better date handling and comparison.
	Agenda       string    // Consider using a more structured format for agenda, e.g., a slice of strings.
	Participants []string  // Use a sorted slice for participants to maintain consistency.
	Notes        string    // Use a multi-line string for notes to allow for more detailed discussions.
	ActionItems  []string  // Consider using a struct for action items with due dates and assignees.
}

func NewMeeting(date time.Time, agenda string, participants []string) Meeting {
	return Meeting{
		Date:         date,
		Agenda:       agenda,
		Participants: participants,
		Notes:        "",         // Empty string literal can be replaced with nil for better clarity.
		ActionItems:  []string{}, // Empty slice literal can be replaced with nil for better clarity.
	}
}

// AddActionItem adds an action item to the meeting.
func (m *Meeting) AddActionItem(actionItem string) {
	m.ActionItems = append(m.ActionItems, actionItem)
}

// Meeting struct to hold meeting details.
type Meeting struct {
	Date         time.Time
	Agenda       string
	Participants []string
	Notes        *string
	ActionItems  []string
}

// NewMeeting creates a new meeting instance.
func NewMeeting(date time.Time, agenda string, participants []string) Meeting {
	return Meeting{
		Date:         date,
		Agenda:       agenda,
		Participants: participants,
		Notes:        nil,
		ActionItems:  nil,
	}
}

// AddActionItems adds multiple action items to the meeting.
func (m *Meeting) AddActionItems(actionItems ...string) {
	m.ActionItems = append(m.ActionItems, actionItems...)
}
