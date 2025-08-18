package model

import (
	"time"

	"github.com/luizmarinhojr/metrics/internal/http/api/view/response"
	"gorm.io/gorm"
)

type Broker struct {
	gorm.Model
	ID        uint     `gorm:"autoIncrement"`
	Nome      string   `gorm:"unique;not null"`
	UserID    uint     `gorm:"unique;not null"`
	Metricas  []Metric `gorm:"foreignKey:CorretorID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (b *Broker) NewResponse() *response.Broker {
	return &response.Broker{
		ID:        b.ID,
		Nome:      b.Nome,
		CreatedAt: &b.CreatedAt,
		UpdatedAt: &b.UpdatedAt,
	}
}
