package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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

func (uh *UserHandler) GetByEmail(ctx *gin.Context) {
	var request request.User
	err := ctx.BindJSON(&request)
	err = request.Validate()
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	token, erro := uh.UserUseCase.GetByEmail(&request)
	if erro != nil {
		ctx.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	ctx.SetCookie("token", *token, int(time.Hour*24*15), "/", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "Login succesful!"})
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
