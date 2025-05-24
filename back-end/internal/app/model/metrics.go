package model

import "gorm.io/gorm"

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
