package main

import (
	"fmt"
	"go-orm/database"
	"go-orm/models"
)

func main() {
	db, err := database.ConnectDatabase()

	if err != nil {
		panic("Erro ao conectar database")
	}
	eroor := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{})

	if eroor != nil {
		fmt.Println("MIGRATE FEITO COM SUCESSO")
	} else {
		fmt.Println("MIGRAÇÃO FEITO COM SUCESSO")
	}
}
