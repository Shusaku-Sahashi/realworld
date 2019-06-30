package resource

/*
	REST での Request body を定義する。
*/

type LoginRequest struct {
	User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"user"`
}

func (req *LoginRequest) Email() string {
	return req.User.Email
}

func (req *LoginRequest) Password() string {
	return req.User.Password
}
