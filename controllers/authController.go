package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"restAuth/database"
	"restAuth/models"
)

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func Register(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"status": "error",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(fiber.Map{
			"status": "Error hashing password",
			"error":  err.Error(),
		})
	}

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: string(hashedPassword),
	}

	database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"status": "success",
		"user":   user,
	})
}
