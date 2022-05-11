package main

import (
	"github.com/gofiber/fiber/v2"
	"restAuth/database"
	"restAuth/routes"
)

func main() {
	database.Connect()
	app := fiber.New()

	routes.Setup(app)

	app.Listen(":3000")
}
