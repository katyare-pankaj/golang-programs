package main

import (
	"fmt"
	"reflect"
)

type QuizOption interface{}

type Question struct {
	Question string
	Options  []QuizOption
	Answer   QuizOption
}

type Quiz struct {
	Questions []Question
}

func (q *Quiz) AddQuestion(question Question) {
	q.Questions = append(q.Questions, question)
}

func (q *Quiz) AskQuestion(index int) {
	question := q.Questions[index]
	fmt.Printf("Question %d: %s\n", index+1, question.Question)
	for i, option := range question.Options {
		fmt.Printf("%d. %v\n", i+1, option)
	}
}

func (q *Quiz) ValidateAnswer(index int, answer interface{}) bool {
	expectedAnswer := q.Questions[index].Answer
	return reflect.DeepEqual(expectedAnswer, answer)
}

func main() {
	quiz := Quiz{}

	quiz.AddQuestion(Question{
		Question: "What is the capital of France?",
		Options:  []QuizOption{"Paris", "Berlin", "Rome", "Madrid"},
		Answer:   "Paris",
	})

	quiz.AddQuestion(Question{
		Question: "Which number comes after 5?",
		Options:  []QuizOption{4, 5, 6, 7},
		Answer:   6,
	})

	quiz.AddQuestion(Question{
		Question: "Is the sky blue?",
		Options:  []QuizOption{true, false},
		Answer:   true,
	})

	for i := 0; i < len(quiz.Questions); i++ {
		quiz.AskQuestion(i)
		fmt.Print("Enter your answer: ")
		var userInput interface{}
		fmt.Scanln(&userInput)

		answerCorrect := quiz.ValidateAnswer(i, userInput)
		if answerCorrect {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("Incorrect. The correct answer is %v\n", quiz.Questions[i].Answer)
		}
	}
}
