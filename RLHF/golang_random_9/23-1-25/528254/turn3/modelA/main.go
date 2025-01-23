package main

import (
	"fmt"
	"sort"
)

// Update the Post struct to include a Comments field.
type Post struct {
	ID       int
	Title    string
	Likes    int // This field is retained for completeness, though not used now
	Comments int
}

var posts = []Post{
	{ID: 1, Title: "Post 1", Likes: 25, Comments: 10},
	{ID: 2, Title: "Post 2", Likes: 10, Comments: 15},
	{ID: 3, Title: "Post 3", Likes: 50, Comments: 5},
	{ID: 4, Title: "Post 4", Likes: 30, Comments: 20},
	{ID: 5, Title: "Post 5", Likes: 18, Comments: 18},
}

// Function to find the top three posts based on the number of comments
func topThreeMostCommentedPosts(posts []Post) []Post {
	// Sort the posts slice in descending order based on comments
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Comments > posts[j].Comments
	})

	// Handle cases where there are fewer than 3 posts
	n := len(posts)
	if n < 3 {
		return posts // Return all posts if less than 3
	}
	return posts[:3] // Return the top 3 posts by comments
}

func main() {
	// Display the top 3 most commented posts
	mostCommentedPosts := topThreeMostCommentedPosts(posts)
	fmt.Println("Top 3 Most Commented Posts:")
	for _, post := range mostCommentedPosts {
		fmt.Printf("Post ID: %d, Title: %s, Comments: %d\n", post.ID, post.Title, post.Comments)
	}
}
