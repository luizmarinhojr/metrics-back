package middleware

import (
	"log"
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
		log.Println("no token JWT available on request")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	log.Printf("token JWT: %v", token)

	claims, erro := ca.jwt.ValidateJWT(&token)
	if erro != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.Set("user_id", claims.UserID)
	c.Set("corretor_id", claims.BrokerID)
	log.Printf("user_id: %v", claims.UserID)
	c.Next()
}
