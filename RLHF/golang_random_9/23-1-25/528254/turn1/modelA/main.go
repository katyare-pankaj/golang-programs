package main

import (
	"fmt"
	"sort"
)

// Define the data structures
type Post struct {
	ID    int
	Title string
	Views int
}

type Comment struct {
	UserID int
	PostID int
}

type User struct {
	ID       int
	Name     string
	Comments int
}

// Example data
var posts = []Post{
	{ID: 1, Title: "Go Slices Tutorial", Views: 150},
	{ID: 2, Title: "Concurrency in Go", Views: 200},
	{ID: 3, Title: "Understanding Interfaces", Views: 180},
}

var comments = []Comment{
	{UserID: 1, PostID: 1},
	{UserID: 1, PostID: 2},
	{UserID: 2, PostID: 2},
	{UserID: 3, PostID: 3},
	{UserID: 3, PostID: 1},
}

var users = []User{
	{ID: 1, Name: "Alice"},
	{ID: 2, Name: "Bob"},
	{ID: 3, Name: "Charlie"},
}

// Find most viewed posts
func getMostViewedPosts(posts []Post) []Post {
	// Sort posts by views in descending order
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Views > posts[j].Views
	})
	return posts
}

// Find top commenters
func getTopCommenters(comments []Comment, users []User) []User {
	userCommentCount := make(map[int]int)

	// Count comments per user
	for _, comment := range comments {
		userCommentCount[comment.UserID]++
	}

	// Update the users slice with the comment count
	for i, user := range users {
		if count, exists := userCommentCount[user.ID]; exists {
			users[i].Comments = count
		}
	}

	// Sort users by comment count in descending order
	sort.Slice(users, func(i, j int) bool {
		return users[i].Comments > users[j].Comments
	})
	return users
}

// Display user engagement trends
func displayUserEngagementTrends(comments []Comment) {
	postCommentCount := make(map[int]int)

	// Count comments per post
	for _, comment := range comments {
		postCommentCount[comment.PostID]++
	}

	fmt.Println("User Engagement Trends: Comments per Post")
	for postID, count := range postCommentCount {
		fmt.Printf("Post ID %d: %d comments\n", postID, count)
	}
}

func main() {
	// Calculate and display the most viewed posts
	mostViewedPosts := getMostViewedPosts(posts)
	fmt.Println("Most Viewed Posts:")
	for _, post := range mostViewedPosts {
		fmt.Printf("%s with %d views\n", post.Title, post.Views)
	}

	fmt.Println()

	// Calculate and display the top commenters
	topCommenters := getTopCommenters(comments, users)
	fmt.Println("Top Commenters:")
	for _, user := range topCommenters {
		fmt.Printf("%s with %d comments\n", user.Name, user.Comments)
	}

	fmt.Println()

	// Display user engagement trends
	displayUserEngagementTrends(comments)
}
