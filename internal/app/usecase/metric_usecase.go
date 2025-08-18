package usecase

import (
	"github.com/luizmarinhojr/metrics/internal/app/repository"
	"github.com/luizmarinhojr/metrics/internal/http/api/view/request"
	"github.com/luizmarinhojr/metrics/internal/http/api/view/response"
)

type MetricUseCase struct {
	repository *repository.MetricRepository
}

func newMetricUseCase(rp *repository.MetricRepository) *MetricUseCase {
	return &MetricUseCase{
		repository: rp,
	}
}

func (mu *MetricUseCase) GetAll() *[]response.Metric {
	metrics := mu.repository.GetAll()

	var metricsResponse []response.Metric
	for _, metric := range *metrics {
		metricsResponse = append(metricsResponse, *metric.NewResponse())
	}
	return &metricsResponse
}

func (mu *MetricUseCase) Create(metric *request.Metric, id uint) (*response.Metric, error) {
	metricModel := metric.New()
	metricModel.CorretorID = id
	err := mu.repository.Create(metricModel)
	if err != nil {
		return nil, err
	}
	return metricModel.NewResponse(), nil
}

func (mu *MetricUseCase) GetById(metric *request.MetricId) (*response.Metric, error) {
	metricModel := metric.New()
	err := mu.repository.GetById(metricModel)
	if err != nil {
		return nil, err
	}
	return metricModel.NewResponse(), nil
}
