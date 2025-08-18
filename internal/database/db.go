package database

import (
	"log"

	"github.com/luizmarinhojr/metrics/config"
	"github.com/luizmarinhojr/metrics/internal/app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DB_CONN), &gorm.Config{})

	if err != nil {
		log.Fatal("Error to connect to database:", err)
	}

	db.AutoMigrate(&model.User{}, &model.Broker{}, &model.Metric{})

	log.Println("Database Ready!")

	return db
}
