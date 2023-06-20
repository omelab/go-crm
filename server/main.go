// main.go
package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/omelab/go-crm/app/models"
	"github.com/omelab/go-crm/config"
	"github.com/omelab/go-crm/database"
	"github.com/omelab/go-crm/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	//database connection
	database.Connect()
	defer database.Close()

	// Migrate the User model
	db := database.DB
	db.AutoMigrate(&models.User{})
	 
	// Initialize routes
	routes.SetupAuthRoutes(app, db)
	
	port := fmt.Sprintf(":%d", config.AppPort)
	log.Printf("Starting %s on port %d\n", config.AppName, config.AppPort)
	log.Fatal(app.Listen(port))
 
}