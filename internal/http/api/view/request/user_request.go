package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/luizmarinhojr/metrics/internal/app/model"
)

type User struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=40"`
}

func (u *User) New() *model.User {
	return &model.User{
		Email:    u.Email,
		Password: []byte(u.Password),
	}
}

func (br *User) Validate() error {
	validate := validator.New()
	err := validate.Struct(br)
	if err != nil {
		return err
	}
	return nil
}
