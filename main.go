package main

import (
	"github.com/davidyanez27/go-gorm-restapi/src/data/models"
	"github.com/davidyanez27/go-gorm-restapi/src/db"
	"github.com/davidyanez27/go-gorm-restapi/src/presentation"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.Product{})
	db.DB.AutoMigrate(models.Product{})
	db.DB.AutoMigrate(models.ProductType{})

	server := presentation.Server(3000, presentation.SetupRoutes)
	server.Start()

}
