package routes

import (
	"github.com/gofiber/fiber/v2"
	"restAuth/controllers"
)

func Setup(app *fiber.App) {
	app.Get("/", controllers.Hello)
	app.Post("api/register", controllers.Register)
	app.Post("api/login", controllers.Login)
	app.Get("api/user", controllers.User)
}
