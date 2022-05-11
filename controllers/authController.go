package controllers

import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func Register(c *fiber.Ctx) error {

	var data map[string]interface{}
	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"status": "error",
		})
	}
	return c.JSON(fiber.Map{
		"status": "success",
		"user":   data,
	})
}
