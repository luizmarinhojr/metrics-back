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

func (bh *BrokerHandler) GetAll(ctx *gin.Context) {
	brokers := bh.usecase.GetAll()
	ctx.JSON(http.StatusOK, brokers)
}

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
