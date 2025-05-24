package repository

import "gorm.io/gorm"

type Repository struct {
	BrokerRepository *BrokerRepository
	MetricRepository *MetricRepository
}

func NewRepository(db *gorm.DB) *Repository {
	br := newBrokerRepository(db)
	mr := newMetricRepository(db)

	return &Repository{
		BrokerRepository: br,
		MetricRepository: mr,
	}
}
