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
	Options  []string
	Answer   string
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
				options := field.Interface().([]string)
				question.Options = append(question.Options, options...)
			case "Answer":
				question.Answer = field.String()
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
		options += fmt.Sprintf("%d. %s\n", i+1, option)
	}
	return fmt.Sprintf("Question %d: %s\nOptions:\n%s", index+1, question.Question, options)
}

func (q *Quiz) CheckAnswer(index int, answer string) bool {
	return q.Questions[index].Answer == answer
}

func main() {
	quiz := Quiz{}
	quiz.GenerateQuiz([]interface{}{
		Question{
			Question: "What is the capital of France?",
			Options:  []string{"Paris", "Berlin", "Rome", "Madrid"},
			Answer:   "Paris",
		},
		Question{
			Question: "Which programming language is used for web development?",
			Options:  []string{"Python", "Java", "JavaScript", "C++"},
			Answer:   "JavaScript",
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
