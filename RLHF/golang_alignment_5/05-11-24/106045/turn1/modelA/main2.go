package main

import "fmt"

type Competency struct {
	Name string
}

type Employee struct {
	Name         string
	Competencies []*Competency
}

type Role struct {
	Name                 string
	RequiredCompetencies []*Competency
}

func main() {
	leadership := &Competency{Name: "Leadership"}
	communication := &Competency{Name: "Communication"}
	technicalProficiency := &Competency{Name: "Technical Proficiency"}
	strategicThinking := &Competency{Name: "Strategic Thinking"}

	emp := Employee{
		Name:         "Alice",
		Competencies: []*Competency{leadership, communication, technicalProficiency},
	}

	role := Role{
		Name:                 "Manager",
		RequiredCompetencies: []*Competency{leadership, communication, strategicThinking},
	}

	fmt.Println("Employee Competencies:", getCompetencyNames(emp.Competencies))
	fmt.Println("Role Required Competencies:", getCompetencyNames(role.RequiredCompetencies))
}

func getCompetencyNames(competencies []*Competency) []string {
	names := make([]string, len(competencies))
	for i, comp := range competencies {
		names[i] = comp.Name
	}
	return names
}
