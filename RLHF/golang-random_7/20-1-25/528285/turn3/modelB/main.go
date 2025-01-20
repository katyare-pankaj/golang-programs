package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

type Question interface {
	Ask() string
	Validate(answer string) bool
}

type StringQuestion struct {
	Question string
	Options  []string
	Answer   string
}

func (q StringQuestion) Ask() string {
	var options string
	for i, option := range q.Options {
		options += fmt.Sprintf("%d. %s\n", i+1, option)
	}
	return fmt.Sprintf("Question: %s\nOptions:\n%sEnter your answer: ", q.Question, options)
}

func (q StringQuestion) Validate(answer string) bool {
	for _, option := range q.Options {
		if option == answer {
			return true
		}
	}
	return false
}

type IntegerQuestion struct {
	Question string
	Options  []int
	Answer   int
}

func (q IntegerQuestion) Ask() string {
	var options string
	for i, option := range q.Options {
		options += fmt.Sprintf("%d. %d\n", i+1, option)
	}
	return fmt.Sprintf("Question: %s\nOptions:\n%sEnter your answer: ", q.Question, options)
}

func (q IntegerQuestion) Validate(answer string) bool {
	num, err := strconv.Atoi(answer)
	if err != nil {
		return false // Invalid integer input
	}
	for _, option := range q.Options {
		if option == num {
			return true
		}
	}
	return false
}

func createQuestion(questionType string) Question {
	switch questionType {
	case "string":
		return StringQuestion{
			Question: "What is the capital of France?",
			Options:  []string{"Paris", "Berlin", "Rome", "Madrid"},
			Answer:   "Paris",
		}
	case "integer":
		return IntegerQuestion{
			Question: "What is the sum of 10 and 20?",
			Options:  []int{25, 30, 35, 40},
			Answer:   30,
		}
	default:
		panic("Invalid question type: " + questionType)
	}
}

func main() {
	questionTypes := []string{"string", "integer"}
	rand.Seed(time.Now().UnixNano())

	for _, questionType := range questionTypes {
		q := createQuestion(questionType)
		fmt.Print(q.Ask())
		var userAnswer string
		fmt.Scanln(&userAnswer)

		if q.Validate(userAnswer) {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect. The correct answer is", reflect.ValueOf(q).Elem().FieldByName("Answer").Interface())
		}
	}
}
