package main

import "fmt"

type WorkflowStep interface {
	Execute()
}

type PrimitiveStep struct {
	name   string
	action func()
}

func (s *PrimitiveStep) Execute() {
	fmt.Println(fmt.Sprintf("Executing step: %s", s.name))
	s.action()
}

func NewPrimitiveStep(name string, action func()) *PrimitiveStep {
	return &PrimitiveStep{name: name, action: action}
}

type CompositeStep struct {
	steps []WorkflowStep
}

func (s *CompositeStep) Execute() {
	fmt.Println("Executing composite step:")
	for _, step := range s.steps {
		step.Execute()
	}
}

func (s *CompositeStep) AddStep(step WorkflowStep) {
	s.steps = append(s.steps, step)
}

func main() {
	// Create some primitive steps
	sendEmail := NewPrimitiveStep("Send Email", func() {
		fmt.Println("Email sent.")
	})

	processData := NewPrimitiveStep("Process Data", func() {
		fmt.Println("Data processed.")
	})

	uploadData := NewPrimitiveStep("Upload Data", func() {
		fmt.Println("Data uploaded.")
	})

	validateInput := NewPrimitiveStep("Validate Input", func() {
		fmt.Println("Input validated.")
	})

	// Create composite steps
	dataProcessingPipeline := &CompositeStep{steps: []WorkflowStep{
		validateInput,
		processData,
		uploadData,
	}}

	customerOnboarding := &CompositeStep{steps: []WorkflowStep{
		sendEmail,
		dataProcessingPipeline,
	}}

	// Create a workflow by composing steps
	workflow := &CompositeStep{steps: []WorkflowStep{
		customerOnboarding,
	}}

	workflow.Execute()
}
