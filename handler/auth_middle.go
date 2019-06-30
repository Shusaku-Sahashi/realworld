package handler

import (
	"github.com/app/realworld/model/user"
)

//go:generate mockery -name=AuthMiddle -outpkg=mock -output=mock
type AuthMiddle interface {
	Authenticate(email string, password string) (user.LoginInfo, string, error)
}
