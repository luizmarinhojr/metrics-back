package database

import (
	"fmt"
	"log"

	"github.com/luizmarinhojr/metrics/config"
	"github.com/luizmarinhojr/metrics/internal/app/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	dbPath := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.PSQL_HOST_DEV, config.PSQL_PORT_DEV, config.PSQL_USER_DEV, config.PSQL_PASS_DEV, config.PSQL_DBNAME_DEV)
	db, err := gorm.Open(postgres.Open(dbPath), &gorm.Config{})

	if err != nil {
		log.Fatal("Error to connect to database:", err)
	}

	db.AutoMigrate(&model.Broker{}, &model.Metric{})

	log.Println("Database Ready!")

	return db
}
