package main

import (
	"fmt"
	"go-orm/database"
)

type User struct {
	ID   uint
	Name string
}

type Product struct {
	ID   uint
	Name string
}

type Order struct {
	ID    uint
	Order string
}

func main() {
	db, err := database.ConnectDatabase()

	if err != nil {
		panic("Erro ao conectar database")
	}
	eroor := db.AutoMigrate(&User{}, &Product{}, &Order{})

	if eroor != nil {
		fmt.Println("MIGRATE FEITO COM SUCESSO")
	} else {
		fmt.Println("MIGRAÇÃO FEITO COM SUCESSO")
	}
}
