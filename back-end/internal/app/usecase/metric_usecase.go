package usecase

import "github.com/luizmarinhojr/metrics/internal/app/repository"

type MetricUseCase struct {
	repository *repository.MetricRepository
}

func newMetricUseCase(rp *repository.MetricRepository) *MetricUseCase {
	return &MetricUseCase{
		repository: rp,
	}
}
