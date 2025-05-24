package handler

import "github.com/luizmarinhojr/metrics/internal/app/usecase"

type Handler struct {
	BrokerHandler *BrokerHandler
	MetricHandler *MetricHandler
}

func NewHandler(usecases *usecase.UseCase) *Handler {
	bh := newBrokerHandler(usecases.BrokerUseCase)
	mu := newMetricHandler(usecases.MetricUseCase)

	return &Handler{
		BrokerHandler: bh,
		MetricHandler: mu,
	}
}
