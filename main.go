package main

import (
	"go-orm/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	url, _ := config.LoadingConfig()
	_, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		panic("Erro ao conectar database")
	}

}
