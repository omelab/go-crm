// database/connection.go
package database

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/omelab/go-crm/config"
)

var DB *gorm.DB

func Connect() {
	//  DatabaseURL := "host=localhost user=user password=password dbname=database port=5432 sslmode=disable TimeZone=Asia/Dhaka"

	db, err := gorm.Open("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	DB = db
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}