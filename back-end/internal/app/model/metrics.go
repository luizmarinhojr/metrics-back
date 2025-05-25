package model

import (
	"github.com/luizmarinhojr/metrics/internal/http/api/view/response"
	"gorm.io/gorm"
)

type Metric struct {
	gorm.Model
	ID             uint   `gorm:"autoIncrement"`
	Data           string `gorm:"not null"`
	CorretorID     uint   `gorm:"not null"`
	LeadsRecebidos int
	Ligacoes       int
	Espontaneo     int
	Captacoes      int
	Visitas        int
	Negociacoes    int
	Propostas      int
	Vendas         int
	Observacao     string
}

func (mt *Metric) NewResponse() *response.Metric {
	return &response.Metric{
		ID:             mt.ID,
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
