package repository

import "gorm.io/gorm"

type Repository struct {
	BrokerRepository *BrokerRepository
	MetricRepository *MetricRepository
	UserRepository   *UserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		BrokerRepository: newBrokerRepository(db),
		MetricRepository: newMetricRepository(db),
		UserRepository:   newUserRepository(db),
	}
}
