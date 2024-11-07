package userauth

type Authenticator interface {
	Login(username, password string) (userID int, err error)
	Logout(userID int) error
}

type basicAuthenticator struct{}

func (a *basicAuthenticator) Login(username, password string) (int, error) {
	// Implementation details for login
	return 1, nil
}

func (a *basicAuthenticator) Logout(userID int) error {
	// Implementation details for logout
	return nil
}

func NewAuthenticator() Authenticator {
	return &basicAuthenticator{}
}
