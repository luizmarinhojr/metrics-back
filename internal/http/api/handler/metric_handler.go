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

// @Summary Obter todas as métricas
// @Description Retorna a lista de todas as métricas.
// @Tags metrics
// @Accept json
// @Produce json
// @Success 200 {array} response.Metric "Lista de métricas"
// @Router /metrics [get]
func (mh *MetricHandler) GetAll(ctx *gin.Context) {
	metrics := mh.usecase.GetAll()
	ctx.JSON(http.StatusOK, metrics)
}

// @Summary Criar uma nova métrica
// @Description Cria uma nova métrica com os dados fornecidos. Requer autenticação.
// @Tags metrics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param metric body request.Metric true "Métrica a ser criada"
// @Success 201 {object} string "ID da métrica criada"
// @Failure 400 {object} string "Erro de validação"
// @Failure 401 {string} string "Não autorizado"
// @Failure 500 {object} string "Erro interno do servidor"
// @Router /metric [post]
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

// @Summary Obter métrica pelo ID
// @Description Retorna uma métrica específica pelo ID. Requer autenticação.
// @Tags metrics
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param metric body request.MetricId true "ID da métrica"
// @Success 200 {object} response.Metric "Métrica encontrada"
// @Failure 400 {object} string "Erro de validação"
// @Failure 401 {string} string "Não autorizado"
// @Failure 404 {object} string "Métrica não encontrada"
// @Router /metric/id [get]
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
