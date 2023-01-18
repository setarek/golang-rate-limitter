package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"golang-rate-limitter/config"
	"golang-rate-limitter/internal/blog/model"
	"golang-rate-limitter/pkg/logger"
)

func InitDB(config *config.Config, logger logger.Logger) (*gorm.DB, error) {

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s", config.GetString("db_user"),
		config.GetString("db_password"),
		config.GetString("db_host"),
		config.GetString("db_port"))
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		logger.Error("error while creating database connection")
		return nil, err
	}

	db.AutoMigrate(&model.BlogModel{})
	return db, nil

}
