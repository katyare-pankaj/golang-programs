package main

import (
	"fmt"
	"sort"
)

// Define structures for Posts and Users
type Post struct {
	ID    int
	Title string
	Likes int
}

type User struct {
	ID    int
	Name  string
	Likes int
}

// Sample data for posts and users
var posts = []Post{
	{ID: 1, Title: "Go Slices Explained", Likes: 120},
	{ID: 2, Title: "Understanding Goroutines", Likes: 150},
	{ID: 3, Title: "Mastering Interfaces", Likes: 110},
	{ID: 4, Title: "Go's Concurrency Model", Likes: 300},
}

var users = []User{
	{ID: 1, Name: "Alice", Likes: 75},
	{ID: 2, Name: "Bob", Likes: 120},
	{ID: 3, Name: "Charlie", Likes: 95},
	{ID: 4, Name: "David", Likes: 250},
}

// Function to get top 3 most liked posts
func getTopLikedPosts(posts []Post) []Post {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Likes > posts[j].Likes
	})
	if len(posts) > 3 {
		return posts[:3]
	}
	return posts
}

// Function to get top 3 users who liked the most
func getTopLikingUsers(users []User) []User {
	sort.Slice(users, func(i, j int) bool {
		return users[i].Likes > users[j].Likes
	})
	if len(users) > 3 {
		return users[:3]
	}
	return users
}

func main() {
	// Display top 3 most liked posts
	topPosts := getTopLikedPosts(posts)
	fmt.Println("Top 3 Most Liked Posts:")
	for _, post := range topPosts {
		fmt.Printf("Title: %s, Likes: %d\n", post.Title, post.Likes)
	}

	fmt.Println()

	// Display top 3 users who liked the most
	topUsers := getTopLikingUsers(users)
	fmt.Println("Top 3 Users With Most Likes:")
	for _, user := range topUsers {
		fmt.Printf("Name: %s, Likes: %d\n", user.Name, user.Likes)
	}
}
