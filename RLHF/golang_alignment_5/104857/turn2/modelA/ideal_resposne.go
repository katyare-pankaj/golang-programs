package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// BDD tests
func TestSynchronizeUser(t *testing.T) {
	t.Run("it should synchronize user name", func(t *testing.T) {
		existingUser := User{ID: 1, Name: "Alice", Email: "alice@example.com", version: 0}
		updatedFields := map[string]string{"name": "Bob"}
		expectedUser := User{ID: 1, Name: "Bob", Email: "alice@example.com", version: 1}
		syncedUser := synchronizeUser(existingUser, updatedFields)
		assert.Equal(t, expectedUser, syncedUser, "user should be updated with the new name")
	})

	t.Run("it should synchronize user email", func(t *testing.T) {
		existingUser := User{ID: 2, Name: "Charlie", Email: "charlie@example.com", version: 0}
		updatedFields := map[string]string{"email": "charlie_new@example.com"}
		expectedUser := User{ID: 2, Name: "Charlie", Email: "charlie_new@example.com", version: 1}
		syncedUser := synchronizeUser(existingUser, updatedFields)
		assert.Equal(t, expectedUser, syncedUser, "user should be updated with the new email")
	})

	t.Run("it should not modify the existing user", func(t *testing.T) {
		existingUser := User{ID: 3, Name: "David", Email: "david@example.com", version: 0}
		updatedFields := map[string]string{"name": "David Updated"}
		synchronizeUser(existingUser, updatedFields)
		assert.Equal(t, 0, existingUser.version, "existing user should not be modified")
	})
}
