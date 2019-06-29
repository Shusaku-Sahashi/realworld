package repositories

import (
	"github.com/app/realworld/models"
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

func (rep *UserRepository) GetById(id uint) (*models.User, error) {
	user := new(models.User)
	err := rep.DB.Find(user, "where = ?", id).Error
	return user, err
}

func (rep *UserRepository) Create(user models.User) (*models.User, error) {
	err := rep.DB.Create(user).Error
	return &user, err
}
