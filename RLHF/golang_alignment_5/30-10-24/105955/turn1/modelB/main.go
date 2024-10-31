package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/google/go-github/v32/github"
	"github.com/robfig/cron"
)

type CodeReviewTool struct {
	client  *github.Client
	mu      sync.Mutex
	metrics map[string]int
}

func NewCodeReviewTool(token string) *CodeReviewTool {
	// Create a new GitHub client using the provided access token
	client := github.NewClient(nil)
	client.SetOAuthToken(token)

	return &CodeReviewTool{
		client:  client,
		metrics: make(map[string]int),
	}
}

func (cr *CodeReviewTool) runSamplingTask() {
	// Define the time interval for sampling (every 10 seconds)
	samplingInterval := "*/10 * * * *"

	// Use the cron package to schedule the sampling task
	c := cron.New()
	c.AddFunc(samplingInterval, cr.sampleMetrics)
	c.Start()
}

func (cr *CodeReviewTool) sampleMetrics() {
	// Collect data from GitHub API to calculate metrics
	repos, _, err := cr.client.Repositories.List("", nil)
	if err != nil {
		log.Printf("Error fetching repositories: %v", err)
		return
	}

	for _, repo := range repos {
		// Your code to calculate metrics based on the repositories and their pulls/issues
		// For simplicity, let's just count the number of repositories
		cr.mu.Lock()
		cr.metrics["repositories"]++
		cr.mu.Unlock()
	}
}

func (cr *CodeReviewTool) displayMetrics() {
	// Display the calculated metrics
	for metric, value := range cr.metrics {
		fmt.Printf("%s: %d\n", metric, value)
	}
}

func main() {
	// Get the GitHub access token from the environment variable
	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		log.Fatal("GITHUB_TOKEN environment variable not set")
	}

	// Create a new CodeReviewTool instance
	cr := NewCodeReviewTool(githubToken)

	// Start the sampling task
	cr.runSamplingTask()

	// Display metrics every 30 seconds
	for {
		time.Sleep(30 * time.Second)
		cr.displayMetrics()
	}
}
