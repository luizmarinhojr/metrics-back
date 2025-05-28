package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/luizmarinhojr/metrics/internal/app/model"
)

type Metric struct {
	Data           string `json:"data" validate:"required,min=3,max=30"`
	CorretorID     uint   `json:"id_corretor" validate:"required"`
	LeadsRecebidos int    `json:"leads_recebidos" validate:"max=2000"`
	Ligacoes       int    `json:"ligacoes" validate:"max=2000"`
	Espontaneo     int    `json:"espontaneo" validate:"max=2000"`
	Captacoes      int    `json:"captacoes" validate:"max=2000"`
	Visitas        int    `json:"visitas" validate:"max=2000"`
	Negociacoes    int    `json:"negociacoes" validate:"max=2000"`
	Propostas      int    `json:"propostas" validate:"max=2000"`
	Vendas         int    `json:"vendas" validate:"max=2000"`
	Observacao     string `json:"observacao" validate:"max=2000"`
}

func (mt *Metric) New() *model.Metric {
	return &model.Metric{
		Data:           mt.Data,
		CorretorID:     mt.CorretorID,
		LeadsRecebidos: mt.LeadsRecebidos,
		Ligacoes:       mt.Ligacoes,
		Espontaneo:     mt.Espontaneo,
		Captacoes:      mt.Captacoes,
		Visitas:        mt.Visitas,
		Negociacoes:    mt.Negociacoes,
		Propostas:      mt.Propostas,
		Vendas:         mt.Vendas,
		Observacao:     mt.Observacao,
	}
}

func (mt *Metric) Validate() error {
	validator := validator.New()
	err := validator.Struct(mt)
	if err != nil {
		return err
	}
	return nil
}

type MetricId struct {
	ID uint `json:"id" validate:"required"`
}

func (mt *MetricId) New() *model.Metric {
	return &model.Metric{
		ID: mt.ID,
	}
}

func (mt *MetricId) Validate() error {
	validator := validator.New()
	err := validator.Struct(mt)
	if err != nil {
		return err
	}
	return nil
}
