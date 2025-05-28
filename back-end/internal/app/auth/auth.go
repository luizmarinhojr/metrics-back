package auth

type Auth struct {
	JWT *JsonWebToken
}

func NewAuth() *Auth {
	return &Auth{
		JWT: newJWT(),
	}
}
