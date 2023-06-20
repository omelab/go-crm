// app/handlers/auth_handler.go
package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/omelab/go-crm/app/models"
	"github.com/omelab/go-crm/app/services"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	db            *gorm.DB
	authService   *services.AuthService
}

func NewAuthHandler(db *gorm.DB, authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		db:            db,
		authService:   authService,
	}
}
 

func (h *AuthHandler) Register(c *fiber.Ctx) error {

	type RegisterRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := new(RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal server error"})
	}

	// Create user
	user := models.User{Email: req.Email, Password: string(hashedPassword)}
	if err := h.db.Create(&user).Error; err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Internal server error"})
	}

	return c.JSON(fiber.Map{"message": "User registered successfully"})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	type LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		log.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	token, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid email or password"})
	}

	return c.JSON(fiber.Map{"token": token})
}