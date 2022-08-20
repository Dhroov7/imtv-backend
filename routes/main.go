package routes

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(api fiber.Router) {
	RegisterPageRoutes(api)
	RegisterMovieRoutes(api)
}
