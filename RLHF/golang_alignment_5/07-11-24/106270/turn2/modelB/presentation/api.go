// presentation/api.go

package presentation

import (
	"encoding/json"
	businesslogic "go-programs/RLHF/golang_alignment_5/07-11-24/106270/turn2/modelB/business_logic"
	"net/http"
	"strconv"
)

// UserAPI handles user-related API endpoints
type UserAPI struct {
	us businesslogic.UserService
}

// NewUserAPI creates a new UserAPI instance using the provided user service
func NewUserAPI(us businesslogic.UserService) *UserAPI {
	return &UserAPI{us: us}
}

func (u *UserAPI) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Extract the user ID from the request path
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Delegate the business logic to the user service
	user, err := u.us.GetUserByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the user data in JSON format
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
