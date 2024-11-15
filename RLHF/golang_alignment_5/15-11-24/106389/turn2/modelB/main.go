package main

import (
	"fmt"
)

type WorkflowStep interface {
	Execute()
}

type EmailStep struct {
	message string
}

func (s *EmailStep) Execute() {
	fmt.Println("Sending email:", s.message)
}

func newEmailStep(message string) *EmailStep {
	return &EmailStep{message: message}
}

type DataProcessingStep struct {
	data string
}

func (s *DataProcessingStep) Execute() {
	fmt.Println("Processing data:", s.data)
}

func newDataProcessingStep(data string) *DataProcessingStep {
	return &DataProcessingStep{data: data}
}

type WorkflowSequence struct {
	steps []WorkflowStep
}

func (w *WorkflowSequence) Execute() {
	for _, step := range w.steps {
		step.Execute()
	}
}

func (w *WorkflowSequence) AddStep(step WorkflowStep) {
	w.steps = append(w.steps, step)
}

func main() {
	// Creating primitive workflow steps
	emailStep1 := newEmailStep("Welcome aboard!")
	dataProcessingStep1 := newDataProcessingStep("Customer data")
	emailStep2 := newEmailStep("Data processing complete")

	// Composite workflow step: Data processing pipeline
	dataProcessingPipeline := &WorkflowSequence{
		steps: []WorkflowStep{
			dataProcessingStep1,
		},
	}

	// Customer onboarding workflow
	customerOnboarding := &WorkflowSequence{
		steps: []WorkflowStep{
			emailStep1,
			dataProcessingPipeline,
			emailStep2,
		},
	}

	// Data backup workflow
	dataBackup := &WorkflowSequence{
		steps: []WorkflowStep{
			newDataProcessingStep("Backup data"),
		},
	}

	fmt.Println("Running Customer Onboarding Workflow:")
	customerOnboarding.Execute()

	fmt.Println("\nRunning Data Backup Workflow:")
	dataBackup.Execute()
}
