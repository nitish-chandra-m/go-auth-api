package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nitish-chandra-m/go-auth/src/github.com/nitish-chandra-m/go-auth/controllers"
)

func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

}
