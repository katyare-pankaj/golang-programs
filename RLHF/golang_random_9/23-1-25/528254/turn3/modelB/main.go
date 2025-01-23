package main

import (
	"fmt"
	"sort"
)

type Post struct {
	ID       int
	Title    string
	Likes    int
	Comments int // Added Comments field in the Post struct
}
type User struct {
	ID    int
	Name  string
	Likes int
}

var posts = []Post{
	{ID: 1, Title: "Post 1", Likes: 25, Comments: 10}, // Added comments for each post
	{ID: 2, Title: "Post 2", Likes: 10, Comments: 5},
	{ID: 3, Title: "Post 3", Likes: 50, Comments: 20},
	{ID: 4, Title: "Post 4", Likes: 30, Comments: 15},
	{ID: 5, Title: "Post 5", Likes: 18, Comments: 8},
}
var users = []User{
	{ID: 101, Name: "User A", Likes: 8},
	{ID: 102, Name: "User B", Likes: 20},
	{ID: 103, Name: "User C", Likes: 12},
	{ID: 104, Name: "User D", Likes: 35},
	{ID: 105, Name: "User E", Likes: 27},
}

func topThreeMostLikedPosts(posts []Post) []Post {
	//Sort the posts slice in descending order based on comments
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Comments > posts[j].Comments
	})

	return posts[:3]
}

func topThreeMostLikedUsers(users []User) []User {
	// Sort the users slice in descending order based on likes
	sort.Slice(users, func(i, j int) bool {
		return users[i].Likes > users[j].Likes
	})
	return users[:3]
}
func main() {
	// Display the top 3 most liked posts based on comments
	topCommentedPosts := topThreeMostLikedPosts(posts)
	fmt.Println("Top 3 Most Commented Posts:")
	for _, post := range topCommentedPosts {
		fmt.Printf("Post ID: %d, Title: %s, Likes: %d, Comments: %d\n", post.ID, post.Title, post.Likes, post.Comments)
	}

	fmt.Println()

	//Display the top 3 most liked users
	topLikedUsers := topThreeMostLikedUsers(users)
	fmt.Println("Top 3 Most Liked Users:")
	for _, user := range topLikedUsers {
		fmt.Printf("User ID: %d, Name: %s, Likes: %d\n", user.ID, user.Name, user.Likes)
	}
}
