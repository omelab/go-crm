// app/handlers/user_handler.go
package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Profile(c *fiber.Ctx) error {
	user := c.Locals("user")
	return c.JSON(fiber.Map{"email": user})
}
