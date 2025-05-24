package handler

import "github.com/luizmarinhojr/metrics/internal/app/usecase"

type MetricHandler struct {
	usecase *usecase.MetricUseCase
}

func newMetricHandler(us *usecase.MetricUseCase) *MetricHandler {
	return &MetricHandler{
		usecase: us,
	}
}
