package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/luizmarinhojr/metrics/internal/app/model"
)

var jwtKey = []byte("secret-key")

type claims struct {
	UserID   uint `json:"user_id"`
	BrokerID uint `json:"corretor_id"`
	jwt.RegisteredClaims
}

type JsonWebToken struct {
	Claims *claims
}

func newJWT() *JsonWebToken {
	return &JsonWebToken{
		Claims: &claims{},
	}
}

func (j *JsonWebToken) GenerateJWT(user *model.User) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24 * 15)
	j.Claims.UserID = user.ID
	j.Claims.BrokerID = user.Corretor.ID
	j.Claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j.Claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j *JsonWebToken) ValidateJWT(tokenString *string) (*claims, error) {
	claims := &claims{}

	token, err := jwt.ParseWithClaims(*tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}

	return claims, nil
}
