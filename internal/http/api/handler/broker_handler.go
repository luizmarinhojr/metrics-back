package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/metrics/internal/app/usecase"
	"github.com/luizmarinhojr/metrics/internal/http/api/view/request"
)

type BrokerHandler struct {
	usecase *usecase.BrokerUseCase
}

func newBrokerHandler(us *usecase.BrokerUseCase) *BrokerHandler {
	return &BrokerHandler{
		usecase: us,
	}
}

// @Summary Obter todos os brokers
// @Description Retorna a lista de todos os brokers. Requer autenticação.
// @Tags brokers
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} response.Broker "Lista de brokers"
// @Failure 401 {string} string "Não autorizado"
// @Router /brokers [get]
func (bh *BrokerHandler) GetAll(ctx *gin.Context) {
	brokers := bh.usecase.GetAll()
	ctx.JSON(http.StatusOK, brokers)
}

// @Summary Criar um novo broker
// @Description Cria um novo broker com os dados fornecidos. Requer autenticação.
// @Tags brokers
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param broker body request.Broker true "Broker a ser criado"
// @Success 201 {object} string "ID do broker criado"
// @Failure 400 {object} string "Erro de validação"
// @Failure 401 {string} string "Não autorizado"
// @Failure 500 {object} string "Erro interno do servidor"
// @Router /broker [post]
func (bh *BrokerHandler) Create(ctx *gin.Context) {
	var broker request.Broker
	err := ctx.BindJSON(&broker)
	err = broker.Validate()
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	br, erro := bh.usecase.Create(&broker)
	if erro != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusCreated)
	ctx.Header("id", fmt.Sprintf("%d", br.ID))
}

// @Summary Obter broker pelo nome
// @Description Retorna um broker específico pelo nome. Requer autenticação.
// @Tags brokers
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param broker body request.BrokerName true "Nome do broker"
// @Success 200 {object} response.Broker "Broker encontrado"
// @Failure 400 {object} string "Erro de validação"
// @Failure 401 {string} string "Não autorizado"
// @Failure 404 {object} string "Broker não encontrado"
// @Router /broker/name [post]
func (bh *BrokerHandler) GetByName(ctx *gin.Context) {
	var broker request.BrokerName
	err := ctx.BindJSON(&broker)
	err = broker.Validate()
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	br, erro := bh.usecase.GetByName(&broker)
	if erro != nil {
		ctx.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, br)
}

// @Summary Obter broker pelo ID
// @Description Retorna um broker específico pelo ID. Requer autenticação.
// @Tags brokers
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param broker body request.BrokerId true "ID do broker"
// @Success 200 {object} response.Broker "Broker encontrado"
// @Failure 400 {object} string "Erro de validação"
// @Failure 401 {string} string "Não autorizado"
// @Failure 404 {object} string "Broker não encontrado"
// @Router /broker/id [post]
func (bh *BrokerHandler) GetById(ctx *gin.Context) {
	var broker request.BrokerId
	err := ctx.BindJSON(&broker)
	err = broker.Validate()
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	br, erro := bh.usecase.GetById(&broker)
	if erro != nil {
		ctx.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, br)
}
