package model

import (
	"time"

	"github.com/luizmarinhojr/metrics/internal/http/api/view/response"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"autoIncrement"`
	Email     string `gorm:"unique;not null"`
	Password  []byte `gorm:"not null"`
	Role      string `gorm:"default:corretor"`
	Corretor  Broker `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (u *User) NewResponse() *response.User {
	return &response.User{
		ID:        u.ID,
		Email:     u.Email,
		CreatedAt: &u.CreatedAt,
		UpdatedAt: &u.UpdatedAt,
	}
}
