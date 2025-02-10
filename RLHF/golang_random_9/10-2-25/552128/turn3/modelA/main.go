package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// PollOption represents a voting option.
type PollOption struct {
	Option string
	Votes  int
}

// UserVotes tracks the vote for each user identified by a user identifier.
type UserVotes map[string]int

func main() {
	options := []PollOption{
		{Option: "Go"},
		{Option: "Python"},
		{Option: "JavaScript"},
		{Option: "C#"},
		{Option: "Other"},
	}

	optionIndices := make(map[string]int)
	for idx, opt := range options {
		optionIndices[opt.Option] = idx
	}

	userVotes := make(UserVotes)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Current Poll Status:")
		for i, opt := range options {
			fmt.Printf("[%d] %s - %d votes\n", i+1, opt.Option, opt.Votes)
		}

		fmt.Println("Enter your user ID (or 'exit' to quit):")
		userID, _ := reader.ReadString('\n')
		userID = strings.TrimSpace(userID)
		if strings.ToLower(userID) == "exit" {
			break
		}

		currentVoteIdx, hasVoted := userVotes[userID]
		if hasVoted {
			fmt.Printf("You have an existing vote for: %s\n", options[currentVoteIdx].Option)
			options[currentVoteIdx].Votes--
		}

		fmt.Println("Please enter the number of your new choice:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > len(options) {
			fmt.Println("Invalid choice, please try again.")
			continue
		}

		var optionToVoteFor string
		if choice == len(options) {
			fmt.Print("Enter your preferred option for 'Other': ")
			optionToVoteFor, _ = reader.ReadString('\n')
			optionToVoteFor = strings.TrimSpace(optionToVoteFor)

			if _, exists := optionIndices[optionToVoteFor]; !exists {
				options = append(options, PollOption{Option: optionToVoteFor})
				optionIndices[optionToVoteFor] = len(options) - 1
			}
		} else {
			optionToVoteFor = options[choice-1].Option
		}

		idx, ok := optionIndices[optionToVoteFor]
		if !ok {
			fmt.Println("Error updating vote count. Please try again.")
			continue
		}

		options[idx].Votes++
		userVotes[userID] = idx

		fmt.Println("Thank you for voting!\n")
	}

	fmt.Println("Final Poll Results:")
	for _, opt := range options {
		fmt.Printf("%s - %d votes\n", opt.Option, opt.Votes)
	}
}
