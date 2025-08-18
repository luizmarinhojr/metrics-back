package usecase

import (
	"github.com/luizmarinhojr/metrics/internal/app/auth"
	"github.com/luizmarinhojr/metrics/internal/app/repository"
	"github.com/luizmarinhojr/metrics/internal/app/service"
	"github.com/luizmarinhojr/metrics/internal/http/api/view/request"
	"github.com/luizmarinhojr/metrics/internal/http/api/view/response"
)

type UserUseCase struct {
	repository      *repository.UserRepository
	passwordService *service.PasswordService
	auth            *auth.Auth
}

func newUserUseCase(rp *repository.UserRepository, auth *auth.Auth) *UserUseCase {
	return &UserUseCase{
		repository:      rp,
		passwordService: service.NewService().PasswordService,
		auth:            auth,
	}
}

func (us *UserUseCase) GetByEmail(user *request.User) (*string, *response.Broker, error) {
	userModel := user.New()
	err := us.repository.GetByEmail(userModel)
	if err != nil {
		return nil, nil, err
	}
	erro := us.passwordService.CheckPasswordHash([]byte(user.Password), userModel.Password)
	if erro != nil {
		return nil, nil, erro
	}
	token, erroa := us.auth.JWT.GenerateJWT(userModel)
	if erroa != nil {
		return nil, nil, erroa
	}
	return &token, userModel.Corretor.NewResponse(), nil
}

func (us *UserUseCase) Create(user *request.User) (*response.User, error) {
	userModel := user.New()
	hashedPassword, err := us.passwordService.HashPassword(userModel.Password)
	if err != nil {
		return nil, err
	}
	userModel.Password = hashedPassword
	erro := us.repository.Create(userModel)
	if erro != nil {
		return nil, erro
	}
	return userModel.NewResponse(), err
}

func (us *UserUseCase) ValidateJWT(token *string) error {
	_, err := us.auth.JWT.ValidateJWT(token)
	return err
}
