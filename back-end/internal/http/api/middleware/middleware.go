package middleware

import "github.com/luizmarinhojr/metrics/internal/app/auth"

type Middleware struct {
	CheckAuth *checkAuth
}

func NewMiddleware(auth *auth.Auth) *Middleware {
	return &Middleware{
		CheckAuth: newCheckAuth(auth.JWT),
	}
}
