package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luizmarinhojr/metrics/internal/app/usecase"
	"github.com/luizmarinhojr/metrics/internal/http/api/view/request"
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

func (mh *MetricHandler) Create(ctx *gin.Context) {
	var metric request.Metric
	err := ctx.BindJSON(&metric)
	err = metric.Validate()
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	corretorID := ctx.GetUint("corretor_id")
	response, erro := mh.usecase.Create(&metric, corretorID)
	if erro != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx.Status(http.StatusCreated)
	ctx.Header("id", fmt.Sprintf("%d", response.ID))
}

func (mh *MetricHandler) GetByID(ctx *gin.Context) {
	var metric request.MetricId
	err := ctx.BindJSON(&metric)
	err = metric.Validate()
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	response, erro := mh.usecase.GetById(&metric)
	if erro != nil {
		ctx.Writer.WriteHeader(http.StatusNotFound)
		return
	}
	ctx.JSON(http.StatusOK, response)
}
