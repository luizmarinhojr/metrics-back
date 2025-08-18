package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/metrics/config"
	"github.com/luizmarinhojr/metrics/internal/app/usecase"
	"github.com/luizmarinhojr/metrics/internal/http/api/view/request"
)

type UserHandler struct {
	UserUseCase *usecase.UserUseCase
}

func newUserHandler(us *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		UserUseCase: us,
	}
}

func (uh *UserHandler) Login(ctx *gin.Context) {
	var request request.User
	err := ctx.BindJSON(&request)
	err = request.Validate()
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	token, corretor, erro := uh.UserUseCase.GetByEmail(&request)
	if erro != nil {
		ctx.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	ctx.SetCookie("token", *token, int(864000), config.PATH, config.DOMAIN, false, true)
	ctx.JSON(http.StatusOK, gin.H{"corretor": corretor.Nome})
}

func (uh *UserHandler) Create(ctx *gin.Context) {
	var request request.User
	err := ctx.BindJSON(&request)
	err = request.Validate()
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	user, erro := uh.UserUseCase.Create(&request)
	if erro != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusCreated)
	ctx.Header("id", fmt.Sprintf("%d", user.ID))
}

func (uh *UserHandler) ValidateToken(ctx *gin.Context) {
	token, err := ctx.Cookie("token")
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"valid": false})
		return
	}
	erro := uh.UserUseCase.ValidateJWT(&token)
	if erro != nil {
		ctx.JSON(http.StatusOK, gin.H{"valid": false})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"valid": true})
}

func (uh *UserHandler) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", int(864000), config.PATH, config.DOMAIN, true, true)
	ctx.Writer.WriteHeader(http.StatusOK)
}
