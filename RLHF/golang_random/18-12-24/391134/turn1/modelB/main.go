package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func fetchUser(ctx context.Context, userID int) (*User, error) {
	// Create a new HTTP client
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Create a new request
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://api.example.com/users/%d", userID), nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request")
	}

	// Perform the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error performing request")
	}

	// Defer the closure of the response body
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("expected status code 200, got %d", resp.StatusCode)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "error reading response body")
	}

	// Decode the response body into a User struct
	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, errors.Wrap(err, "error decoding response body")
	}

	return &user, nil
}

func main() {
	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a slice of user IDs to fetch
	userIDs := []int{1, 2, 3, 4, 5}

	// Create a channel to receive fetched users
	usersCh := make(chan *User, len(userIDs))

	// Spawn Goroutines to fetch users concurrently
	for _, userID := range userIDs {
		go func(userID int) {
			user, err := fetchUser(ctx, userID)
			if err != nil {
				// Log the error and continue with the next user
				fmt.Printf("Error fetching user %d: %v\n", userID, err)
				return
			}

			// Send the fetched user to the channel
			usersCh <- user
		}(userID)
	}

	// Receive fetched users from the channel and print them
	for i := 0; i < len(userIDs); i++ {
		user := <-usersCh
		if user != nil {
			fmt.Printf("User %d: %s, %d\n", userIDs[i], user.Name, user.Age)
		}
	}
}
