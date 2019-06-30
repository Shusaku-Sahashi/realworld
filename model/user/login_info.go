package user

type LoginInfo struct {
	Email    string
	Username string
	Bio      string
}

func NewLoginInfo(u *User) LoginInfo {
	return LoginInfo{
		Email:    u.Email,
		Username: u.Username,
		Bio:      u.Bio,
	}
}
