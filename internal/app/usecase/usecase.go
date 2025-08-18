package usecase

import (
	"github.com/luizmarinhojr/metrics/internal/app/auth"
	"github.com/luizmarinhojr/metrics/internal/app/repository"
)

type UseCase struct {
	BrokerUseCase *BrokerUseCase
	MetricUseCase *MetricUseCase
	UserUseCase   *UserUseCase
}

func NewUseCase(repositories *repository.Repository) *UseCase {
	return &UseCase{
		BrokerUseCase: newBrokerUseCase(repositories.BrokerRepository),
		MetricUseCase: newMetricUseCase(repositories.MetricRepository),
		UserUseCase:   newUserUseCase(repositories.UserRepository, auth.NewAuth()),
	}
}
