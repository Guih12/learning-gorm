package database

import (
	"errors"
	"go-orm/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase() (*gorm.DB, error) {
	url, _ := config.LoadingConfig()
	database, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		return nil, errors.New("Error! Connected not completed")
	}
	return database, nil
}
