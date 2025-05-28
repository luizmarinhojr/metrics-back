package repository

import (
	"github.com/luizmarinhojr/metrics/internal/app/model"
	"gorm.io/gorm"
)

type BrokerRepository struct {
	db *gorm.DB
}

func newBrokerRepository(db *gorm.DB) *BrokerRepository {
	return &BrokerRepository{
		db: db,
	}
}

func (br *BrokerRepository) GetAll() *[]model.Broker {
	var brokers []model.Broker
	br.db.Find(&brokers)
	return &brokers
}

func (br *BrokerRepository) Create(broker *model.Broker) error {
	return br.db.Create(broker).Error
}

func (br *BrokerRepository) GetByName(broker *model.Broker) error {
	return br.db.Where("nome = ?", broker.Nome).First(broker).Error
}

func (br *BrokerRepository) GetById(broker *model.Broker) error {
	return br.db.First(&broker, broker.ID).Error
}
