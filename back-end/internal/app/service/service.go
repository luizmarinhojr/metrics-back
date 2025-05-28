package service

type Service struct {
	PasswordService *PasswordService
}

func NewService() *Service {
	return &Service{
		PasswordService: newPasswordService(),
	}
}
