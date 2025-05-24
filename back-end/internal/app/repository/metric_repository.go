package repository

import (
	"github.com/luizmarinhojr/metrics/internal/app/model"
	"gorm.io/gorm"
)

type MetricRepository struct {
	db *gorm.DB
}

func newMetricRepository(db *gorm.DB) *MetricRepository {
	return &MetricRepository{
		db: db,
	}
}

func (mr *MetricRepository) GetAll() *[]model.Metric {
	var metrics []model.Metric
	mr.db.Find(&metrics)
	return &metrics
}
