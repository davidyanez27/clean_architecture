package db

import (
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DSN = "sqlserver://sa:yourStrong&Password@localhost:1433?database=inventario"

// var DSN = "host=localhost user=sa password=yourStrong&Password dbname=prueba port:1433"
var DB *gorm.DB

func DBConnection() {

	var error error
	DB, error = gorm.Open(sqlserver.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connected")
	}

}
