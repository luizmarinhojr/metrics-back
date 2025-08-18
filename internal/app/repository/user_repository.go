package repository

import (
	"github.com/luizmarinhojr/metrics/internal/app/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func newUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) GetByEmail(user *model.User) error {
	return ur.db.Where("email = ?", user.Email).Preload("Corretor").First(user).Error
}

func (ur *UserRepository) Create(user *model.User) error {
	return ur.db.Create(user).Error
}
