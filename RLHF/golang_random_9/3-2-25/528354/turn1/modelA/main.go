package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type User struct {
	Username string
	Password string
	Role     string
}

type File struct {
	Name string
	Path string
}

var (
	users = map[string]User{
		"admin": {"admin", "password", "admin"},
		"user":  {"user", "password", "user"},
	}

	files = map[string]File{
		"file1.txt": {"file1.txt", "files/file1.txt"},
		"file2.txt": {"file2.txt", "files/file2.txt"},
	}
)

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/rename", authMiddleware(authorizationMiddleware(renameFileHandler)))
	http.HandleFunc("/move", authMiddleware(authorizationMiddleware(moveFileHandler)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if user, exists := users[username]; exists && user.Password == password {
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: user.Role,
		})
		log.Printf("User %s logged in", username)
		fmt.Fprintf(w, "Login successful")
		return
	}

	http.Error(w, "Invalid username or password", http.StatusUnauthorized)
}

func renameFileHandler(w http.ResponseWriter, r *http.Request) {
	oldName := r.FormValue("oldName")
	newName := r.FormValue("newName")

	if file, exists := files[oldName]; exists {
		newPath := filepath.Join("files", newName)
		if os.Rename(file.Path, newPath) == nil {
			files[newName] = File{Name: newName, Path: newPath}
			delete(files, oldName)
			log.Printf("File renamed from %s to %s", oldName, newName)
			fmt.Fprintf(w, "File renamed successfully")
			return
		}
	}

	http.Error(w, "Failed to rename file", http.StatusBadRequest)
}

func moveFileHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	newPath := r.FormValue("newPath")

	if file, exists := files[name]; exists {
		if os.Rename(file.Path, newPath) == nil {
			file.Path = newPath
			files[name] = file
			log.Printf("File %s moved to %s", name, newPath)
			fmt.Fprintf(w, "File moved successfully")
			return
		}
	}

	http.Error(w, "Failed to move file", http.StatusBadRequest)
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil || cookie == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		r.Header.Set("Role", cookie.Value)
		next(w, r)
	}
}

func authorizationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role := r.Header.Get("Role")
		if !strings.Contains("admin", role) { // Only admin is allowed
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
