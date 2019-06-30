package model

type User struct {
	ID       uint   `gorm:"column:id;primary_key"`
	Username string `gorm:"column:user_name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	Bio      string `gorm:"column:bio"`
	// Image    URL    `gorm:"image"`
}
