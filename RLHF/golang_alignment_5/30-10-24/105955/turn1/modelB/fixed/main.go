package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/google/go-github/v32/github"
	"github.com/robfig/cron/v3"
	"golang.org/x/oauth2"
)

type CodeReviewTool struct {
	client  *github.Client
	mu      sync.Mutex
	metrics map[string]int
}

func NewCodeReviewTool(token string) *CodeReviewTool {
	// Create a new GitHub client using the provided access token
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

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
	log.Println("Sampling task started...")
}

func (cr *CodeReviewTool) sampleMetrics() {
	// Collect data from GitHub API to calculate metrics
	ctx := context.Background()
	repos, _, err := cr.client.Repositories.List(ctx, "", nil)
	if err != nil {
		log.Printf("Error fetching repositories: %v", err)
		return
	}

	cr.mu.Lock()
	defer cr.mu.Unlock()
	for _, repo := range repos {
		// Count the number of repositories
		cr.metrics["repositories"]++
		log.Printf("Found repository: %s", *repo.Name) // Log the found repository
	}
}

func (cr *CodeReviewTool) displayMetrics() {
	// Display the calculated metrics
	cr.mu.Lock()
	defer cr.mu.Unlock()
	fmt.Println("Displaying metrics:")
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
		time.Sleep(2 * time.Second)
		cr.displayMetrics()
	}
}
