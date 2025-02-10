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

	// Create a map to store user votes
	userVotes := make(map[string]map[string]int)

	reader := bufio.NewReader(os.Stdin)

	var currentUser string
	for {
		fmt.Println("Current User:", currentUser)
		fmt.Println("Current Poll Status:")
		for i, opt := range options {
			fmt.Printf("[%d] %s - %d votes\n", i+1, opt.Option, opt.Votes)
		}

		fmt.Println("Menu:")
		fmt.Println("[1] Vote")
		fmt.Println("[2] Remove Vote")
		fmt.Println("[3] Change User")
		fmt.Println("[4] Exit")
		fmt.Print("Enter your choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			vote(&options, optionIndices, userVotes, currentUser, reader)
		case "2":
			removeVote(&options, userVotes, currentUser, reader)
		case "3":
			fmt.Print("Enter new user ID: ")
			currentUser, _ = reader.ReadString('\n')
			currentUser = strings.TrimSpace(currentUser)
		case "4":
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func vote(options *[]PollOption, optionIndices map[string]int, userVotes map[string]map[string]int, currentUser string, reader *bufio.Reader) {
	fmt.Println("Please enter the number of your choice to vote (or 'exit' to quit):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if strings.ToLower(input) == "exit" {
		return
	}

	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(*options) {
		fmt.Println("Invalid choice, please try again.")
		return
	}

	var optionToVoteFor string
	if choice == len(*options) {
		fmt.Print("Enter your preferred option for 'Other': ")
		optionToVoteFor, _ = reader.ReadString('\n')
		optionToVoteFor = strings.TrimSpace(optionToVoteFor)
		if _, exists := optionIndices[optionToVoteFor]; !exists {
			*options = append(*options, PollOption{Option: optionToVoteFor})
			optionIndices[optionToVoteFor] = len(*options) - 1
		}
	} else {
		optionToVoteFor = (*options)[choice-1].Option
	}

	// Update vote counts
	idx, ok := optionIndices[optionToVoteFor]
	if !ok {
		fmt.Println("Error updating vote count. Please try again.")
		return
	}

	if _, exists := userVotes[currentUser]; !exists {
		userVotes[currentUser] = make(map[string]int)
	}

	userVotes[currentUser][optionToVoteFor]++
	(*options)[idx].Votes++

	fmt.Println("Thank you for voting!\n")
}

func removeVote(options *[]PollOption, userVotes map[string]map[string]int, currentUser string, reader *bufio.Reader) {
	fmt.Println("Please enter the number of the option to remove your vote from (or 'exit' to quit):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if strings.ToLower(input) == "exit" {
		return
	}

	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(*options) {
		fmt.Println("Invalid choice, please try again.")
		return
	}

	optionToRemove := (*options)[choice-1].Option

	// Update vote counts
	idx, ok := optionIndices[optionToRemove]
	if !ok {
		fmt.Println("Error updating vote count. Please try again.")
		return
	}

	if _, exists := userVotes[currentUser]; exists {
		if voteCount, exists := userVotes[currentUser][optionToRemove]; exists {
			if voteCount > 0 {
				userVotes[currentUser][optionToRemove]--
				(*options)[idx].Votes--
				fmt.Println("Vote removed successfully!")
			} else {
				fmt.Println("You haven't voted for this option.")
			}
		} else {
			fmt.Println("You haven't voted for this option.")
		}
	} else {
		fmt.Println("You haven't voted in this poll yet.")
	}
}
