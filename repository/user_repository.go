package repository

import (
	"github.com/app/realworld/model"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (rep *UserRepository) GetById(id uint) (*model.User, error) {
	user := new(model.User)
	err := rep.DB.Find(user, "where = ?", id).Error
	return user, err
}

func (rep *UserRepository) Create(user model.User) (*model.User, error) {
	err := rep.DB.Create(user).Error
	return &user, err
}
