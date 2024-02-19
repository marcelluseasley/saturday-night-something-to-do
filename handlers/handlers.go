package handlers

import (
	"encoding/json"
	"strconv"


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

func (u *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad Request"})
	}
	user, err := u.UserService.GetUser(uID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not Found"})
	}
	return c.JSON(user)
}

func (u *UserHandler) UpdateUser(c *fiber.Ctx) error {
	user := models.User{}
	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad Request"})
	}
	err = u.UserService.UpdateUser(user)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not Found"})
	}
	return c.JSON(user)
}

func (u *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Bad Request"})
	}
	err = u.UserService.DeleteUser(uID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not Found"})
	}
	return c.SendStatus(204)
}

func (u *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := u.UserService.ListUsers()
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Not Found"})
	}
	return c.JSON(users)
}
