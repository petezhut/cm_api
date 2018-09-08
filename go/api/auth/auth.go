package auth

import (
	"fmt"
)

func NewAuth() *Auth {
	return &Auth{Username: DEFAULT_ADMIN_USERNAME, Password: DEFAULT_ADMIN_PASSWORD}
}

func (auth *Auth) String() string {
	return fmt.Sprintf("Auth(User: %s Password: %s)", auth.Username, auth.Password)
}
