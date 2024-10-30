package main

import (
	"fmt"
)

// Define the interface for likeable content
type Likeable interface {
	Like()
	Unlike()
}

// Basic implementation of Likeable for TextPost
type TextPost struct {
	content string
	likes   int
}

func (t *TextPost) Like() {
	t.likes++
}

func (t *TextPost) Unlike() {
	if t.likes > 0 {
		t.likes--
	}
}

// Implement Likeable for ImagePost
type ImagePost struct {
	imageURL string
	likes    int
}

func (i *ImagePost) Like() {
	i.likes++
}

func (i *ImagePost) Unlike() {
	if i.likes > 0 {
		i.likes--
	}
}

// **Uncomment the following lines to see the duplication:**

//func likePost(post *TextPost) {
//	post.Like()
//}
//
//func unlikePost(post *TextPost) {
//	post.Unlike()
//}
//
//func likePost(post *ImagePost) {
//	post.Like()
//}
//
//func unlikePost(post *ImagePost) {
//	post.Unlike()
//}

// We can use the 'reflect' package to achieve similar functionality without code duplication.
func likePost(post interface{}) {
	if value, ok := post.(Likeable); ok {
		value.Like()
	} else {
		fmt.Println("Can't like a non-likeable post.")
	}
}

func unlikePost(post interface{}) {
	if value, ok := post.(Likeable); ok {
		value.Unlike()
	} else {
		fmt.Println("Can't unlike a non-likeable post.")
	}
}

func main() {
	post1 := &TextPost{content: "Hello, world!"}
	post2 := &ImagePost{imageURL: "example.com/img.jpg"}

	likePost(post1)
	likePost(post2)

	unlikePost(post1)
	unlikePost(post2)

	// Print likes count
	fmt.Println("Post 1 likes:", post1.likes)
	fmt.Println("Post 2 likes:", post2.likes)
}
