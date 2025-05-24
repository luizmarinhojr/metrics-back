package usecase

import (
	"github.com/luizmarinhojr/metrics/internal/app/model"
	"github.com/luizmarinhojr/metrics/internal/app/repository"
)

type BrokerUseCase struct {
	repository *repository.BrokerRepository
}

func newBrokerUseCase(rp *repository.BrokerRepository) *BrokerUseCase {
	return &BrokerUseCase{
		repository: rp,
	}
}

func (bu *BrokerUseCase) GetAll() *[]model.Broker {
	return bu.repository.GetAll()
}
