// business_logic/business_logic.go

package businesslogic

import (
	dataaccess "go-programs/RLHF/golang_alignment_5/07-11-24/106270/turn2/modelB/data_access"
)

// UserService provides user-related business logic
type UserService interface {
	GetUserByID(id int) (*dataaccess.User, error)
}

// NewUserService creates a new UserService instance using the provided data access layer
func NewUserService(da dataaccess.DataAccess) UserService {
	return &userServiceImpl{da: da}
}

type userServiceImpl struct {
	da dataaccess.DataAccess
}

func (u *userServiceImpl) GetUserByID(id int) (*dataaccess.User, error) {
	// Implementation details of the business logic for getting a user by ID
	// For simplicity, we'll just delegate to the data access layer.
	return u.da.GetUserByID(id)
}
