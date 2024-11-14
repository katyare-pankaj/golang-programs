package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Employee struct
type Employee struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Skills []Skill `json:"skills"`
}

// Skill struct
type Skill struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Training struct
type Training struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Skills      []Skill    `json:"skills"`
	Employees   []Employee `json:"employees"`
}

func Test_MainModels(t *testing.T) {
	// Test cases for employee, skill, and training models
	emp := Employee{ID: "emp1", Name: "Jane Doe", Skills: []Skill{}}
	skill := Skill{ID: "skill1", Name: "Go Programming", Description: "Go programming language skills"}
	trng := Training{ID: "trng1", Name: "Go Bootcamp", Description: "Introduction to Go", Skills: []Skill{skill}, Employees: []Employee{emp}}

	if diff := cmp.Diff(emp, Employee{ID: "emp1", Name: "Jane Doe", Skills: []Skill{}}); diff != "" {
		t.Errorf("employee struct mismatch: %s", diff)
	}

	if diff := cmp.Diff(skill, Skill{ID: "skill1", Name: "Go Programming", Description: "Go programming language skills"}); diff != "" {
		t.Errorf("skill struct mismatch: %s", diff)
	}

	if diff := cmp.Diff(trng, Training{ID: "trng1", Name: "Go Bootcamp", Description: "Introduction to Go", Skills: []Skill{skill}, Employees: []Employee{emp}}); diff != "" {
		t.Errorf("training struct mismatch: %s", diff)
	}
}
