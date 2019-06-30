package resource

import (
	"github.com/app/realworld/model/user"
)

/*
	REST での Response body　を定義する。
*/

type LoginResponse struct {
	User struct {
		Token    string `json:"token"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Bio      string `json:"bio"`
		//Image    string `json:"image"`
	}
}

func NewLoginResponse(info user.LoginInfo, token string) LoginResponse {
	resp := new(LoginResponse)
	resp.User.Username = info.Username
	resp.User.Email = info.Email
	resp.User.Bio = info.Bio
	resp.User.Token = token

	return *resp
}
