package usecase

import (
	"github.com/luizmarinhojr/metrics/internal/app/repository"
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

func (mu *MetricUseCase) GetAll() []response.Metric {
	metrics := mu.repository.GetAll()

	var metricsResponse []response.Metric
	for _, metric := range *metrics {
		metricsResponse = append(metricsResponse, *metric.NewResponse())
	}
	return metricsResponse
}
