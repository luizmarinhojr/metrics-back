package usecase

import (
	"github.com/luizmarinhojr/metrics/internal/app/repository"
	"github.com/luizmarinhojr/metrics/internal/http/api/view/request"
	"github.com/luizmarinhojr/metrics/internal/http/api/view/response"
)

type BrokerUseCase struct {
	repository *repository.BrokerRepository
}

func newBrokerUseCase(rp *repository.BrokerRepository) *BrokerUseCase {
	return &BrokerUseCase{
		repository: rp,
	}
}

func (bu *BrokerUseCase) GetAll() *[]response.Broker {
	brokers := bu.repository.GetAll()

	var brokersResponse []response.Broker
	for _, broker := range *brokers {
		brokersResponse = append(brokersResponse, *broker.NewResponse())
	}
	return &brokersResponse
}

func (bu *BrokerUseCase) Create(broker *request.Broker) (*response.Broker, error) {
	brokerModel := broker.New()
	err := bu.repository.Create(brokerModel)
	if err != nil {
		return nil, err
	}
	return brokerModel.NewResponse(), nil
}

func (bu *BrokerUseCase) GetByName(broker *request.BrokerName) (*response.Broker, error) {
	brokerModel := broker.New()
	err := bu.repository.GetByName(brokerModel)
	if err != nil {
		return nil, err
	}
	return brokerModel.NewResponse(), nil
}

func (bu *BrokerUseCase) GetById(broker *request.BrokerId) (*response.Broker, error) {
	brokerModel := broker.New()
	err := bu.repository.GetById(brokerModel)
	if err != nil {
		return nil, err
	}
	return brokerModel.NewResponse(), nil
}
