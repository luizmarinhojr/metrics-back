package model

import "gorm.io/gorm"

type Broker struct {
	gorm.Model
	ID       uint     `gorm:"autoIncrement"`
	Nome     string   `gorm:"unique;not null"`
	Metricas []Metric `gorm:"foreignKey:CorretorID"`
}
