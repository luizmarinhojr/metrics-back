package handler

import "github.com/luizmarinhojr/metrics/internal/app/usecase"

type Handler struct {
	BrokerHandler *BrokerHandler
	MetricHandler *MetricHandler
	UserHandler   *UserHandler
}

func NewHandler(usecases *usecase.UseCase) *Handler {
	return &Handler{
		BrokerHandler: newBrokerHandler(usecases.BrokerUseCase),
		MetricHandler: newMetricHandler(usecases.MetricUseCase),
		UserHandler:   newUserHandler(usecases.UserUseCase),
	}
}
