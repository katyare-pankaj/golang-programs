package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Define a PollOption struct to hold each poll option
type PollOption struct {
	Name  string
	Votes int
}

// Define a Poll struct to hold a poll question and its options
type Poll struct {
	Question string
	Options  []PollOption
}

// Define a PollAppState struct to hold the state of the polling application
type PollAppState struct {
	Polls            []Poll
	CurrentPollIndex int
}

// Initialize the application state with some default polls
func initPollAppState() PollAppState {
	return PollAppState{
		Polls: []Poll{
			{
				Question: "What is your favorite programming language?",
				Options: []PollOption{
					{"Go", 0},
					{"Python", 0},
					{"JavaScript", 0},
				},
			},
		},
		CurrentPollIndex: 0,
	}
}

func main() {
	appState := initPollAppState()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Polling Application!")
	for {
		displayMenu()
		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			handleAddPoll(&appState, reader)
		case "2":
			handleViewPolls(appState)
		case "3":
			handleVote(&appState, reader)
		case "4":
			handleDisplayResults(appState)
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func displayMenu() {
	fmt.Println(`
Menu:
1. Add a new poll
2. View all polls
3. Vote for an option
4. Display results
5. Exit
`)
}

func handleAddPoll(appState *PollAppState, reader *bufio.Reader) {
	fmt.Print("Enter poll question: ")
	question, _ := reader.ReadString('\n')
	question = strings.TrimSpace(question)

	var options []PollOption
	fmt.Print("Enter poll options (separated by comma): ")
	optionsStr, _ := reader.ReadString('\n')
	optionsStr = strings.TrimSpace(optionsStr)
	optionNames := strings.Split(optionsStr, ",")
	for _, name := range optionNames {
		options = append(options, PollOption{Name: strings.TrimSpace(name), Votes: 0})
	}

	newPoll := Poll{Question: question, Options: options}
	appState.Polls = append(appState.Polls, newPoll)
	fmt.Println("Poll added successfully!")
}

func handleViewPolls(appState PollAppState) {
	fmt.Println("All Polls:")
	for index, poll := range appState.Polls {
		fmt.Printf("Poll %d: %s\n", index+1, poll.Question)
		for i, option := range poll.Options {
			fmt.Printf("  %d. %s (%d votes)\n", i+1, option.Name, option.Votes)
		}
	}
}

func handleVote(appState *PollAppState, reader *bufio.Reader) {
	handleViewPolls(*appState)
	fmt.Print("Enter poll number to vote for: ")
	pollNumStr, _ := reader.ReadString('\n')
	pollNumStr = strings.TrimSpace(pollNumStr)
	pollNum, err := strconv.Atoi(pollNumStr)
	if err != nil || pollNum < 1 || pollNum > len(appState.Polls) {
		fmt.Println("Invalid poll number. Please try again.")
		return
	}

	currentPoll := &appState.Polls[pollNum-1]
	fmt.Print("Enter option number to vote for: ")
	optionNumStr, _ := reader.ReadString('\n')
	optionNumStr = strings.TrimSpace(optionNumStr)
	optionNum, err := strconv.Atoi(optionNumStr)
	if err != nil || optionNum < 1 || optionNum > len(currentPoll.Options) {
		fmt.Println("Invalid option number. Please try again.")
		return
	}

	currentPoll.Options[optionNum-1].Votes++
	fmt.Println("Vote recorded successfully!")
}

func handleDisplayResults(appState PollAppState) {
	handleViewPolls(appState)
}
