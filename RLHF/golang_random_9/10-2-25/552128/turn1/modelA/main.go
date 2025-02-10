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

func main() {
	// Define the poll options
	options := []PollOption{
		{Option: "Go"},
		{Option: "Python"},
		{Option: "JavaScript"},
	}

	// Create a map to quickly access option indices by option name
	optionIndices := make(map[string]int)
	for idx, opt := range options {
		optionIndices[opt.Option] = idx
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		// Display the poll options and current votes
		fmt.Println("Current Poll Status:")
		for i, opt := range options {
			fmt.Printf("[%d] %s - %d votes\n", i+1, opt.Option, opt.Votes)
		}

		// Ask the user to vote
		fmt.Println("Please enter the number of your choice to vote (or 'exit' to quit):")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Check for exit condition
		if strings.ToLower(input) == "exit" {
			break
		}

		// Convert input to an index
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > len(options) {
			fmt.Println("Invalid choice, please try again.")
			continue
		}

		// Update the votes for the chosen option
		options[choice-1].Votes++

		fmt.Println("Thank you for voting!\n")
	}

	fmt.Println("Final Poll Results:")
	for _, opt := range options {
		fmt.Printf("%s - %d votes\n", opt.Option, opt.Votes)
	}
}
