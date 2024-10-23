package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dghubble/oauth1/v2"
	"github.com/go-redis/redis/v8"
)

// RedisClient is a global Redis client
var RedisClient *redis.Client

// OAuthConfig is the OAuth1 configuration
var OAuthConfig oauth1.Config

func main() {
	// Initialize Redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := RedisClient.Ping(context.Background()).Result()
	fmt.Println(pong, err)
	// Initialize OAuth Config
	OAuthConfig = oauth1.Config{
		ConsumerKey:    os.Getenv("CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("CONSUMER_SECRET"),
		CallbackURL:    "http://your-callback-url",
		Endpoint: oauth1.Endpoint{
			RequestTokenURL: "https://api.twitter.com/oauth/request_token",
			AuthorizeURL:    "https://api.twitter.com/oauth/authorize",
			AccessTokenURL:  "https://api.twitter.com/oauth/access_token",
		},
	}
	// Set up HTTP server and routes
	http.HandleFunc("/start-auth", startAuth)
	http.HandleFunc("/callback", handleCallback)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func startAuth(w http.ResponseWriter, r *http.Request) {
	requestToken, err := OAuthConfig.RequestToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Store the request token in Redis for later use
	err = RedisClient.Set(context.Background(), requestToken.Token, requestToken.Secret, 30*time.Minute).Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	authURL := OAuthConfig.AuthorizeURL(requestToken.Token, nil)
	http.Redirect(w, r, authURL, http.StatusFound)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	_ = vals.Get("oauth_verifier")
	token := vals.Get("oauth_token")

	// Retrieve the request token from Redis
	storedSecret, _ := RedisClient.Get(context.Background(), token).Result()

	fmt.Println(storedSecret)
}
