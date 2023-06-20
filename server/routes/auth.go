// routes/auth.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/omelab/go-crm/app/handlers"
	"github.com/omelab/go-crm/app/repositories"
	"github.com/omelab/go-crm/app/services"
)

func SetupAuthRoutes(app *fiber.App, db *gorm.DB) {
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(db, authService)

	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)
}