package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PollOption struct {
	Option string
	Votes  int
}

func main() {
	options := []PollOption{
		{Option: "Go"},
		{Option: "Python"},
		{Option: "JavaScript"},
		{Option: "C#"},
		{Option: "Other"},
	}

	// Create a map to quickly access options by name
	optionIndices := make(map[string]int)
	for idx, opt := range options {
		optionIndices[opt.Option] = idx
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Current Poll Status:")
		for i, opt := range options {
			fmt.Printf("[%d] %s - %d votes\n", i+1, opt.Option, opt.Votes)
		}

		fmt.Println("Please enter the number of your choice to vote, or 'exit' to quit:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			break
		}

		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > len(options) {
			fmt.Println("Invalid choice, please try again.")
			continue
		}

		// Updated vote logic
		var optionToVoteFor string
		if choice == len(options) {
			fmt.Print("Enter your preferred option for 'Other': ")
			optionToVoteFor, _ = reader.ReadString('\n')
			optionToVoteFor = strings.TrimSpace(optionToVoteFor)
			// Add the new option to the slice if it doesn't exist
			if _, exists := optionIndices[optionToVoteFor]; !exists {
				options = append(options, PollOption{Option: optionToVoteFor})
				optionIndices[optionToVoteFor] = len(options) - 1
			}
		} else {
			optionToVoteFor = options[choice-1].Option
		}

		// Find the index of the voted option in the slice using the map
		idx, ok := optionIndices[optionToVoteFor]
		if !ok {
			fmt.Println("Error updating vote count. Please try again.")
			continue
		}

		options[idx].Votes++

		fmt.Println("Thank you for voting!\n")
	}

	fmt.Println("Final Poll Results:")
	for _, opt := range options {
		fmt.Printf("%s - %d votes\n", opt.Option, opt.Votes)
	}
}
