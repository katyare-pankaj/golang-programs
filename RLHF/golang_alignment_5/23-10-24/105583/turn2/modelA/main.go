package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dghubble/go-twitter/twitter"

	"github.com/go-playground/assert"
	"github.com/stretchr/testify/mock"
)

// MockTwitterClient is a mock implementation of the TwitterClient interface
type MockTwitterClient struct {
	mock.Mock
}

// PostTweet mocks the PostTweet method
func (m *MockTwitterClient) PostTweet(ctx context.Context, tweet string) (*twitter.Tweet, *twitter.Response, error) {
	args := m.Called(ctx, tweet)
	return args.Get(0).(*twitter.Tweet), args.Get(1).(*twitter.Response), args.Error(2)
}

// TwitterClient defines the interface for interacting with Twitter
type TwitterClient interface {
	PostTweet(ctx context.Context, tweet string) (*twitter.Tweet, *twitter.Response, error)
}

// tweetHandler handles HTTP requests to post a tweet
func tweetHandler(w http.ResponseWriter, r *http.Request) {
	tweet := r.FormValue("tweet")
	err := postTweet(tweet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Tweet posted successfully!")
}

// postTweet posts a tweet using the TwitterClient
func postTweet(tweet string) error {
	// In a real implementation, you would inject the TwitterClient here
	tc := getTwitterClient() // <-- Dependency injection
	_, _, err := tc.PostTweet(context.Background(), tweet)
	return err
}

var (
	// replace these with your actual twitter consumer key, consumer secret, access token, and access token secret
	consumerKey       = "YOUR_CONSUMER_KEY"
	consumerSecret    = "YOUR_CONSUMER_SECRET"
	accessToken       = "YOUR_ACCESS_TOKEN"
	accessTokenSecret = "YOUR_ACCESS_TOKEN_SECRET"
)

// getTwitterClient returns a new TwitterClient instance
func getTwitterClient() TwitterClient {
	// In a real implementation, you would use the actual Twitter API client here
	return &realTwitterClient{}
}

type realTwitterClient struct{}

func (c *realTwitterClient) PostTweet(ctx context.Context, tweet string) (*twitter.Tweet, *twitter.Response, error) {
	config := oauth.NewConfig(consumerKey, consumerSecret)
	token := oauth.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth.NoContext, token)

	client := twitter.NewClient(httpClient)

	return client.Statuses.Update(tweet, nil)
}

func TestTweetHandler(t *testing.T) {
	// Test 1: Successful tweet post
	mockClient := &MockTwitterClient{}
	mockClient.On("PostTweet", mock.Anything, "Hello, TDD!").Return(&twitter.Tweet{}, &twitter.Response{}, nil)

	// Override the global getTwitterClient function for testing
	getTwitterClient = func() TwitterClient {
		return mockClient
	}

	req := httptest.NewRequest("POST", "/tweet?tweet=Hello%2C+TDD%21", nil)
	rr := httptest.NewRecorder()

	tweetHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Tweet posted successfully!\n", rr.Body.String())
	mockClient.AssertExpectations(t)

	// Test 2: Tweet post failure
	mockClient.Reset()
	mockClient.On("PostTweet", mock.Anything, "Failed tweet").Return(nil, nil, errors.New("tweet post failed"))

	getTwitterClient = func() TwitterClient {
		return mockClient
	}

	req = httptest.NewRequest("POST", "/tweet?tweet=Failed+tweet", nil)
	rr = httptest.NewRecorder()

	tweetHandler(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	mockClient.AssertExpectations(t)
}
