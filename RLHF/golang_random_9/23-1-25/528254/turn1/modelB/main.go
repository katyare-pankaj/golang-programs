package main

import (
	"fmt"
	"sort"
)

type Post struct {
	ID    int
	Title string
	Views int
}

type Comment struct {
	PostID int
	User   string
	Text   string
}

type User struct {
	Username      string
	CommentsCount int
}

func main() {
	// Sample data
	posts := []Post{
		{1, "Post 1", 100},
		{2, "Post 2", 50},
		{3, "Post 3", 150},
	}

	comments := []Comment{
		{1, "User1", "Comment 1.1"},
		{1, "User2", "Comment 1.2"},
		{2, "User1", "Comment 2.1"},
		{3, "User3", "Comment 3.1"},
		{3, "User3", "Comment 3.2"},
	}

	users := []User{}
	// Calculate user engagement
	for _, comment := range comments {
		userFound := false
		for i, user := range users {
			if user.Username == comment.User {
				users[i].CommentsCount++
				userFound = true
				break
			}
		}
		if !userFound {
			users = append(users, User{Username: comment.User, CommentsCount: 1})
		}
	}
	// Sort posts by views in descending order
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Views > posts[j].Views
	})

	// Sort users by comments count in descending order
	sort.Slice(users, func(i, j int) bool {
		return users[i].CommentsCount > users[j].CommentsCount
	})

	// Display results
	fmt.Println("Most Viewed Posts:")
	for _, post := range posts[:3] {
		fmt.Printf("%d. %s - %d views\n", post.ID, post.Title, post.Views)
	}

	fmt.Println("\nTop Commenters:")
	for _, user := range users[:3] {
		fmt.Printf("%s - %d comments\n", user.Username, user.CommentsCount)
	}

	fmt.Println("\nUser Engagement Trends:")
	for _, user := range users {
		fmt.Printf("%s - %d comments\n", user.Username, user.CommentsCount)
	}
}
