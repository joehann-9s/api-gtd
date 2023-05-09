package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

//var DSN = os.Getenv("DSN")

func DBConnection() {
	DSN := os.Getenv("DSN")
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	}
	log.Println("DB connected")
}
