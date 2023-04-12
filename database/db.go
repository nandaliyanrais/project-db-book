package database

import (
	"fmt"
	"log"
	"project/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "/"
	dbPort   = "5432"
	dbname   = "db-book"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	db.AutoMigrate(models.Book{})
}

func GetDB() *gorm.DB {
	return db
}
