package service

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct{}

func newPasswordService() *PasswordService {
	return &PasswordService{}
}

func (ps *PasswordService) HashPassword(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return hash, err
}

func (ps *PasswordService) CheckPasswordHash(password, hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}
