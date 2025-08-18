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
	mr.db.Preload("Corretor").Order("data desc").Find(&metrics)
	return &metrics
}

func (mr *MetricRepository) Create(metric *model.Metric) error {
	return mr.db.Create(metric).Error
}

func (mr *MetricRepository) GetById(metric *model.Metric) error {
	return mr.db.Preload("Corretor").First(&metric, metric.ID).Error
}
