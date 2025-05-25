package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/metrics/internal/app/usecase"
)

type MetricHandler struct {
	usecase *usecase.MetricUseCase
}

func newMetricHandler(us *usecase.MetricUseCase) *MetricHandler {
	return &MetricHandler{
		usecase: us,
	}
}

func (mh *MetricHandler) GetAll(ctx *gin.Context) {
	metrics := mh.usecase.GetAll()
	ctx.JSON(http.StatusOK, metrics)
}
