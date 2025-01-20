package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type Quiz struct {
	Questions []Question
}

type Question struct {
	Question string
	Options  []interface{}
	Answer   interface{}
}

func (q *Quiz) GenerateQuiz(questions []interface{}) {
	rand.Seed(time.Now().UnixNano())
	for _, qi := range questions {
		qv := reflect.ValueOf(qi)
		if qv.Kind() != reflect.Struct {
			panic("Questions must be structs")
		}
		var question Question
		for i := 0; i < qv.NumField(); i++ {
			field := qv.Field(i)
			fieldType := qv.Type().Field(i).Name
			switch fieldType {
			case "Question":
				question.Question = field.String()
			case "Options":
				options := field.Interface().([]interface{})
				question.Options = append(question.Options, options...)
			case "Answer":
				question.Answer = field.Interface()
			default:
				panic("Unexpected field in question struct")
			}
		}
		q.Questions = append(q.Questions, question)
	}
}

func (q *Quiz) AskQuestion(index int) string {
	question := q.Questions[index]
	var options string
	for i, option := range question.Options {
		options += fmt.Sprintf("%d. %v\n", i+1, option)
	}
	return fmt.Sprintf("Question %d: %s\nOptions:\n%s", index+1, question.Question, options)
}

func (q *Quiz) CheckAnswer(index int, answer string) bool {
	question := q.Questions[index]
	userAnswer := answer

	// Handle different types of options
	answerType := reflect.TypeOf(question.Answer)
	switch answerType.Kind() {
	case reflect.Int:
		var intAnswer int
		_, err := fmt.Sscanf(userAnswer, "%d", &intAnswer)
		if err != nil {
			return false // User input is not an integer
		}
		userAnswer = fmt.Sprintf("%d", intAnswer)
	case reflect.String:
		// No need to do anything special for strings
	default:
		panic("Unsupported answer type")
	}

	// Validate the answer
	return reflect.DeepEqual(question.Answer, userAnswer)
}

func main() {
	quiz := Quiz{}
	quiz.GenerateQuiz([]interface{}{
		Question{
			Question: "What is the capital of France?",
			Options:  []interface{}{"Paris", "Berlin", "Rome", "Madrid"},
			Answer:   "Paris",
		},
		Question{
			Question: "In which decade was Go released?",
			Options:  []interface{}{1990, 2000, 2010, 2020},
			Answer:   2009,
		},
	})

	for i := 0; i < len(quiz.Questions); i++ {
		fmt.Println(quiz.AskQuestion(i))
		var userAnswer string
		fmt.Scanln(&userAnswer)
		if quiz.CheckAnswer(i, userAnswer) {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect. The correct answer is", quiz.Questions[i].Answer)
		}
	}
}
