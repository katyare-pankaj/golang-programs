package main

import (
	"fmt"
	"sort"
)

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

var posts = []Post{
	{ID: 1, Title: "Post 1", Likes: 25},
	{ID: 2, Title: "Post 2", Likes: 10},
	{ID: 3, Title: "Post 3", Likes: 50},
	{ID: 4, Title: "Post 4", Likes: 30},
	{ID: 5, Title: "Post 5", Likes: 18},
}
var users = []User{
	{ID: 101, Name: "User A", Likes: 8},
	{ID: 102, Name: "User B", Likes: 20},
	{ID: 103, Name: "User C", Likes: 12},
	{ID: 104, Name: "User D", Likes: 35},
	{ID: 105, Name: "User E", Likes: 27},
}

func topThreeMostLikedPosts(posts []Post) []Post {
	// Sort the posts slice in descending order based on likes
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Likes > posts[j].Likes
	})

	// Return the top 3 posts
	return posts[:3]
}
func topThreeMostLikedUsers(users []User) []User {
	// Sort the users slice in descending order based on likes
	sort.Slice(users, func(i, j int) bool {
		return users[i].Likes > users[j].Likes
	})

	// Return the top 3 users
	return users[:3]
}
func main() {
	// Display the top 3 most liked posts
	topLikedPosts := topThreeMostLikedPosts(posts)
	fmt.Println("Top 3 Most Liked Posts:")
	for _, post := range topLikedPosts {
		fmt.Printf("Post ID: %d, Title: %s, Likes: %d\n", post.ID, post.Title, post.Likes)
	}

	fmt.Println()

	// Display the top 3 most liked users
	topLikedUsers := topThreeMostLikedUsers(users)
	fmt.Println("Top 3 Most Liked Users:")
	for _, user := range topLikedUsers {
		fmt.Printf("User ID: %d, Name: %s, Likes: %d\n", user.ID, user.Name, user.Likes)
	}
}
