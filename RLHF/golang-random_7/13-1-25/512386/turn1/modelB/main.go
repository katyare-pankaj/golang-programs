package main

import (
	"fmt"
)

// Comment represents a single comment in the thread
type Comment struct {
	ID       int
	Text     string
	Author   string
	Replies  []*Comment // Slice of replies
	ParentID int        // ID of the parent comment (0 for root comments)
}

// BuildThread builds a recursive comment thread from a flat list of comments
func BuildThread(comments []*Comment) []*Comment {
	var thread []*Comment
	// Create a map to look up comments by their ID
	commentMap := make(map[int]*Comment)

	for _, c := range comments {
		commentMap[c.ID] = c
	}

	// Build the thread by linking replies to their parent comments
	for _, c := range comments {
		if c.ParentID == 0 {
			// Root comment, add it to the thread
			thread = append(thread, c)
		} else {
			// Reply, find the parent comment and add it to its Replies slice
			parent, ok := commentMap[c.ParentID]
			if ok {
				parent.Replies = append(parent.Replies, c)
			}
		}
	}

	return thread
}

func main() {
	// Sample comments with their parent IDs
	comments := []*Comment{
		{ID: 1, Text: "Comment 1", Author: "User1"},
		{ID: 2, Text: "Reply to Comment 1", Author: "User2", ParentID: 1},
		{ID: 3, Text: "Comment 3", Author: "User3"},
		{ID: 4, Text: "Reply to Comment 1", Author: "User4", ParentID: 1},
		{ID: 5, Text: "Reply to Reply to Comment 1", Author: "User5", ParentID: 2},
		{ID: 6, Text: "Comment 6", Author: "User6", ParentID: 3},
	}

	thread := BuildThread(comments)

	// Display the comment thread
	fmt.Println("Comment Thread:")
	displayThread(thread, 0)
}

func displayThread(comments []*Comment, indent int) {
	for _, c := range comments {
		fmt.Printf("%s%s - %s\n", string(make([]rune, indent*2)), c.Author, c.Text)
		displayThread(c.Replies, indent+1)
	}
}
