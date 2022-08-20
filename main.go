package main

import (
	"go-fiber/database"
	"go-fiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.ConnectMongo()

	api := app.Group("/api/v1")

	routes.RegisterUserRoutes(api)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	routes.RegisterRoutes(api)

	app.Listen(":8080")
}
