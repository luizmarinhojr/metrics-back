package model

import (
	"time"

	"github.com/luizmarinhojr/metrics/internal/http/api/view/response"
	"gorm.io/gorm"
)

type Metric struct {
	gorm.Model
	ID             uint      `gorm:"autoIncrement"`
	Data           time.Time `gorm:"type:date"`
	CorretorID     uint      `gorm:"not null"`
	LeadsRecebidos int
	Ligacoes       int
	Espontaneo     int
	Captacoes      int
	Visitas        int
	Negociacoes    int
	Propostas      int
	Vendas         int
	Agendamentos   int
	Corretor       Broker `gorm:"foreignKey:CorretorID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

func (mt *Metric) NewResponse() *response.Metric {
	return &response.Metric{
		ID:             mt.ID,
		Data:           mt.Data,
		Corretor:       *mt.Corretor.NewResponse(),
		LeadsRecebidos: mt.LeadsRecebidos,
		Ligacoes:       mt.Ligacoes,
		Espontaneo:     mt.Espontaneo,
		Captacoes:      mt.Captacoes,
		Visitas:        mt.Visitas,
		Negociacoes:    mt.Negociacoes,
		Propostas:      mt.Propostas,
		Vendas:         mt.Vendas,
		Agendamentos:   mt.Agendamentos,
		CreatedAt:      &mt.CreatedAt,
		UpdatedAt:      &mt.UpdatedAt,
	}
}
