package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type ApiResponse struct {
	Status  string      `json:"status"`
	Version string      `json:"version"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = map[int]User{
	1: {ID: 1, Name: "Alice"},
	2: {ID: 2, Name: "Bob"},
}

func respondWithJSON(w http.ResponseWriter, code int, response ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

func handleGetUser(w http.ResponseWriter, r *http.Request, version string) {
	id, err := getUserIdFromRequest(r)
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, ApiResponse{
			Status:  "failure",
			Version: version,
			Error:   err.Error(),
		})
		return
	}

	user, exists := users[id]
	if !exists {
		respondWithJSON(w, http.StatusNotFound, ApiResponse{
			Status:  "failure",
			Version: version,
			Error:   "User not found",
		})
		return
	}

	respondWithJSON(w, http.StatusOK, ApiResponse{
		Status:  "success",
		Version: version,
		Data:    user,
	})
}

func getUserIdFromRequest(r *http.Request) (int, error) {
	idStr := strings.TrimPrefix(r.URL.Path, "/v1/users/")
	var id int
	if _, err := fmt.Sscanf(idStr, "%d", &id); err != nil {
		return 0, fmt.Errorf("invalid user ID")
	}
	return id, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 2 || parts[0] != "v1" {
		respondWithJSON(w, http.StatusNotFound, ApiResponse{
			Status:  "failure",
			Version: "unknown",
			Error:   "API version not supported or invalid URL",
		})
		return
	}

	version := parts[0]
	resource := parts[1]

	switch resource {
	case "users":
		handleGetUser(w, r, version)
	default:
		respondWithJSON(w, http.StatusNotFound, ApiResponse{
			Status:  "failure",
			Version: version,
			Error:   "Resource not found",
		})
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
