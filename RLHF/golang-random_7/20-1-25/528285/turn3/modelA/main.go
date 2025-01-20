package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type QuizOption interface{}

type Question struct {
	QuestionText string
	Options      []QuizOption
	Answer       QuizOption
}

type Quiz struct {
	Questions []Question
}

func (q *Quiz) AddQuestion(question Question) {
	q.Questions = append(q.Questions, question)
}

func (q *Quiz) AskQuestion(index int) {
	question := q.Questions[index]
	fmt.Printf("Question %d: %s\n", index+1, question.QuestionText)
	for i, option := range question.Options {
		fmt.Printf("%d. %v\n", i+1, option)
	}
}

func (q *Quiz) ValidateAnswer(index int, userInput string) bool {
	question := q.Questions[index]
	answerType := reflect.TypeOf(question.Answer)

	var userAnswer QuizOption
	var err error

	switch answerType.Kind() {
	case reflect.String:
		userAnswer = userInput
	case reflect.Int:
		userAnswer, err = strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Invalid input type, expected an integer.")
			return false
		}
	case reflect.Bool:
		// Consider typical variations for boolean input
		lowerInput := strings.ToLower(userInput)
		if lowerInput == "true" || lowerInput == "false" {
			userAnswer = lowerInput == "true"
		} else {
			fmt.Println("Invalid input type, expected true/false.")
			return false
		}
	default:
		fmt.Println("Unhandled answer type.")
		return false
	}

	return reflect.DeepEqual(userAnswer, question.Answer)
}

func main() {
	quiz := Quiz{}

	quiz.AddQuestion(Question{
		QuestionText: "What is the capital of France?",
		Options:      []QuizOption{"Paris", "Berlin", "Rome", "Madrid"},
		Answer:       "Paris",
	})

	quiz.AddQuestion(Question{
		QuestionText: "Which number comes after 5?",
		Options:      []QuizOption{4, 5, 6, 7},
		Answer:       6,
	})

	quiz.AddQuestion(Question{
		QuestionText: "Is the sky blue?",
		Options:      []QuizOption{true, false},
		Answer:       true,
	})

	for i := 0; i < len(quiz.Questions); i++ {
		quiz.AskQuestion(i)
		fmt.Print("Enter your answer: ")
		var userInput string
		fmt.Scanln(&userInput)

		if quiz.ValidateAnswer(i, userInput) {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("Incorrect. The correct answer is %v\n", quiz.Questions[i].Answer)
		}
	}
}
