package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// PollOption represents a single option in a poll.
type PollOption struct {
	Option string
	Votes  int
}

// DisplayPoll prints the current poll status to the console.
func DisplayPoll(options []PollOption) {
	fmt.Println("Current Poll Status:")
	for i, opt := range options {
		fmt.Printf("[%d] %s - %d votes\n", i+1, opt.Option, opt.Votes)
	}
	fmt.Println()
}

// GetVoteInput prompts the user for a vote and handles input.
func GetVoteInput() string {
	fmt.Println("Please enter the number of your choice to vote (or 'exit' to quit):")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// UpdateVote updates the vote count for a given choice.
func UpdateVote(choice int, options []PollOption) {
	if choice >= 1 && choice <= len(options) {
		options[choice-1].Votes++
		fmt.Println("Thank you for voting!\n")
	} else {
		fmt.Println("Invalid choice, please try again.\n")
	}
}

func main() {
	// Initialize the poll options
	options := []PollOption{
		{Option: "Go"},
		{Option: "Python"},
		{Option: "JavaScript"},
	}

	for {
		// Display the current poll status
		DisplayPoll(options)

		// Get user input for voting
		input := GetVoteInput()

		// Check for exit condition
		if strings.ToLower(input) == "exit" {
			break
		}

		// Convert input to integer and update votes if valid
		choice, err := strconv.Atoi(input)
		if err == nil {
			UpdateVote(choice, options)
		} else {
			fmt.Println("Invalid input, please enter a number.\n")
		}
	}

	// Display the final results
	fmt.Println("Final Poll Results:")
	DisplayPoll(options)
}
