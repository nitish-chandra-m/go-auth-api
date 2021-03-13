package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nitish-chandra-m/go-auth/src/github.com/nitish-chandra-m/go-auth/database"
	"github.com/nitish-chandra-m/go-auth/src/github.com/nitish-chandra-m/go-auth/routes"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	app.Listen(":8000")
}
