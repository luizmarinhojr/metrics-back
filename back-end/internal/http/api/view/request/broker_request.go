package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/luizmarinhojr/metrics/internal/app/model"
)

type BrokerName struct {
	Nome string `json:"nome" validate:"required,min=2,max=50"`
}

func (br *BrokerName) New() *model.Broker {
	return &model.Broker{
		Nome: br.Nome,
	}
}

func (br *BrokerName) Validate() error {
	validate := validator.New()
	err := validate.Struct(br)
	if err != nil {
		return err
	}
	return nil
}

type BrokerId struct {
	ID uint `json:"id" validate:"required"`
}

func (br *BrokerId) Validate() error {
	validate := validator.New()
	err := validate.Struct(br)
	if err != nil {
		return err
	}
	return nil
}
