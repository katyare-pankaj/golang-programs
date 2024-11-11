// main.go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang/groupcache/singleflight"
)

// Post represents a social media post
type Post struct {
	ID        int
	Likes     int
	LikedBy   map[int]bool // User IDs
	CreatedAt time.Time
}

// LikeService handles liking and unliking of posts
type LikeService struct {
	posts map[int]*Post
	sg    *singleflight.Group
}

// NewLikeService creates a new LikeService
func NewLikeService() *LikeService {
	return &LikeService{
		posts: make(map[int]*Post),
		sg:    &singleflight.Group{},
	}
}

// LikePost likes a post by userID
func (ls *LikeService) LikePost(postID, userID int) error {
	_, err := ls.sg.Do(fmt.Sprintf("like-%d-%d", postID, userID), func() (interface{}, error) {
		p, ok := ls.posts[postID]
		if !ok {
			return nil, fmt.Errorf("post not found: %d", postID)
		}

		if _, liked := p.LikedBy[userID]; liked {
			return nil, fmt.Errorf("user already liked post: %d", postID)
		}

		p.Likes++
		p.LikedBy[userID] = true
		return nil, nil
	})
	return err
}

func main() {
	ls := NewLikeService()

	// Initialize posts for demonstration
	ls.posts[1] = &Post{ID: 1, Likes: 0, LikedBy: make(map[int]bool)}

	err := ls.LikePost(1, 100)
	if err != nil {
		log.Fatalf("Error liking post: %v", err)
	}

	fmt.Println("Post 1 likes:", ls.posts[1].Likes)
}
