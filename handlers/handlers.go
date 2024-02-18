package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/marcelluseasley/saturday-night-something-to-do/models"
	"github.com/marcelluseasley/saturday-night-something-to-do/service"
)

type UserHandler struct {
	service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{s}
}

func (u *UserHandler) CreateUser(c *fiber.Ctx) error {
	user := models.User{}

	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad Request"})
	}
	err = u.UserService.CreateUser(user)
	return err
}
