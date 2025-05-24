package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/metrics/internal/app/usecase"
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
	ctx.JSON(http.StatusOK, bh.usecase.GetAll())
}
