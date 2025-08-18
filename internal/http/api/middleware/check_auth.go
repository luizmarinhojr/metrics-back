package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/metrics/internal/app/auth"
)

type checkAuth struct {
	jwt *auth.JsonWebToken
}

func newCheckAuth(jwt *auth.JsonWebToken) *checkAuth {
	return &checkAuth{
		jwt: jwt,
	}
}

func (ca *checkAuth) Auth(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, erro := ca.jwt.ValidateJWT(&token)
	if erro != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set("user_id", claims.UserID)
	c.Set("corretor_id", claims.BrokerID)
	c.Next()
}
