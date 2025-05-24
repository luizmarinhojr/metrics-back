package usecase

import "github.com/luizmarinhojr/metrics/internal/app/repository"

type UseCase struct {
	BrokerUseCase *BrokerUseCase
	MetricUseCase *MetricUseCase
}

func NewUseCase(repositories *repository.Repository) *UseCase {
	bu := newBrokerUseCase(repositories.BrokerRepository)
	mu := newMetricUseCase(repositories.MetricRepository)

	return &UseCase{
		BrokerUseCase: bu,
		MetricUseCase: mu,
	}
}
