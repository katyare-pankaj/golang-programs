// main.go

package main

import (
	"fmt"
	"go-programs/RLHF/golang_alignment_5/07-11-24/106270/turn2/modelA/blog"
	"go-programs/RLHF/golang_alignment_5/07-11-24/106270/turn2/modelA/userauth"
)

func main() {
	auth := userauth.NewAuthenticator()
	if err := blog.CreatePost(auth, "Hello, World!", "My first blog post."); err != nil {
		fmt.Println("Error creating post:", err)
	}
}
