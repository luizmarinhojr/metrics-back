package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/metrics/config"
	"github.com/luizmarinhojr/metrics/internal/app/usecase"
	"github.com/luizmarinhojr/metrics/internal/http/api/view/request"
	"github.com/luizmarinhojr/metrics/internal/http/api/view/response"
)

type UserHandler struct {
	UserUseCase *usecase.UserUseCase
}

func newUserHandler(us *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		UserUseCase: us,
	}
}

// @Summary Login do usuário
// @Description Autentica um usuário e define o cookie de sessão com o token JWT.
// @Tags users
// @Accept json
// @Produce json
// @Param user body request.User true "Credenciais do usuário"
// @Success 200 {object} response.UserName "Login bem-sucedido"
// @Failure 400 {object} string "Erro de validação"
// @Failure 401 {string} string "Não autorizado"
// @Router /login [post]
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
	ctx.JSON(http.StatusOK, response.UserName{Corretor: corretor.Nome})
}

// @Summary Criar um novo usuário
// @Description Cria um novo usuário com os dados fornecidos.
// @Tags users
// @Accept json
// @Produce json
// @Param user body request.User true "Usuário a ser criado"
// @Success 201 {object} string "ID do usuário criado"
// @Failure 400 {object} string "Erro de validação"
// @Failure 500 {object} string "Erro interno do servidor"
// @Router /signup [post]
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

// @Summary Validar token de autenticação
// @Description Valida o token JWT contido no cookie.
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} string "Token válido"
// @Router /validate-token [get]
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

// @Summary Logout do usuário
// @Description Remove o cookie de sessão para efetuar o logout. Requer autenticação.
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {string} string "Logout bem-sucedido"
// @Router /logout [get]
func (uh *UserHandler) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", int(864000), config.PATH, config.DOMAIN, true, true)
	ctx.Writer.WriteHeader(http.StatusOK)
}
