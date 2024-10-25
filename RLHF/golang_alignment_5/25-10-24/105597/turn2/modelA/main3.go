package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
)

type user struct {
	ID   int
	Name string
}

func getUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate fetching user data from somewhere
		userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
		ctx := context.WithValue(r.Context(), "user", &user{ID: userID, Name: "User " + strconv.Itoa(userID)})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value("user").(*user)
	fmt.Fprintf(w, "User Profile: ID=%d, Name=%s\n", user.ID, user.Name)
}

func main() {
	http.Handle("/profile", getUserMiddleware(http.HandlerFunc(handleProfile)))
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
