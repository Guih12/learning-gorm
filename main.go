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
	eroor := db.AutoMigrate(&models.User{}, &models.Product{}, &models.Order{}, &models.CreditCard{})

	if eroor != nil {
		fmt.Println("MIGRATE FEITO COM SUCESSO")
	} else {
		fmt.Println("MIGRAÇÃO FEITO COM SUCESSO")
	}
	//product := &models.Product{Name: "Produto 1", Price: 200}
	//products := []models.Product{{Name: "Produto 2", Price: 300}, {Name: "arroz", Price: 2}, {Name: "linguica", Price: 3}}

	//result := db.Create(product)
	//result := db.Create(products)
	//
	//if result.Error != nil {
	//	println("Error", err)
	//}
	//
	//for _, product := range products {
	//	fmt.Println(product.ID, product.Name)
	//}

	//Createwith associations - Método de inserção aninhada igual rails
	result := db.Create(&models.User{Name: "George", LastName: "Borsato", Cpf: "064.111.381-17",
		CreditCard: models.CreditCard{Number: "112-77"}})

	fmt.Println(result.RowsAffected)
}
