package blog

import (
	"fmt"
	"go-programs/RLHF/golang_alignment_5/07-11-24/106270/turn2/modelA/userauth"
)

type BlogPost struct {
	Title    string
	Content  string
	AuthorID int
}

func CreatePost(authenticator userauth.Authenticator, title, content string) error {
	userID, err := authenticator.Login("exampleUser", "secret")
	if err != nil {
		return err
	}
	defer authenticator.Logout(userID)

	// Create a new blog post using the authenticated user ID
	post := BlogPost{Title: title, Content: content, AuthorID: userID}
	fmt.Printf("New post created: %+v\n", post)
	return nil
}
