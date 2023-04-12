package main

import (
	"log"

	"project/database"
	"project/models"
	"project/router"
)

func main() {

	database.StartDB()

	err := database.GetDB().AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatal("Error migrating database schema :", err)
	}

	route := router.SetupRoutes()
	route.Run(":8080")
}
