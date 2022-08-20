package routes

import (
	model "go-fiber/database/models"
	"go-fiber/helper"
	service "go-fiber/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func RegisterUserRoutes(api fiber.Router) {
	api.Post("/login", func(c *fiber.Ctx) error {
		body := new(login)
		if err := c.BodyParser(body); err != nil {
			return err
		}
		user, err := service.GetUser(body.Email)

		// Throws Unauthorized error
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if user.Password != helper.GetHash(body.Password) {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		claims := jwt.MapClaims{
			"id":    user.ID,
			"email": user.Email,
			"admin": true,
			"exp":   time.Now().Add(time.Hour * 72).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte("secret"))

		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"token": t})
	})

	api.Post("/signup", func(c *fiber.Ctx) error {
		user := new(model.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		user.Password = helper.GetHash(user.Password)
		err := service.Save(user)
		if err != nil {
			return c.SendStatus(500)
		}
		return c.SendString("User Created!")
	})
}

func RegisterPageRoutes(api fiber.Router) {
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("home page")
	})

	api.Get("/genres", func(c *fiber.Ctx) error {
		genreList, err := service.GetGenres()
		if err != nil {
			return c.SendStatus(500)
		}
		return c.JSON(fiber.Map{"data": genreList})
	})

	api.Get("/:genre", func(c *fiber.Ctx) error {
		genre := c.Params("genre")
		movieList, err := service.GetMoviesFromGenre(genre)
		if err != nil {
			return c.SendStatus(500)
		}
		return c.JSON(fiber.Map{"data": movieList})
	})

	api.Get("/:category/list", func(c *fiber.Ctx) error {
		category := c.Params("category")
		movieList, err := service.GetMoviesFromCategory(category)
		if err != nil {
			return c.SendStatus(500)
		}
		return c.JSON(fiber.Map{"data": movieList})
	})

	api.Get("/:id/details", func(c *fiber.Ctx) error {
		movieId := c.Params("id")
		movie, err := service.GetMovie(movieId)
		if err != nil {
			return c.SendStatus(500)
		}
		return c.JSON(fiber.Map{"data": movie})
	})
}

func RegisterMovieRoutes(api fiber.Router) {
	api.Post("/movie/add", func(c *fiber.Ctx) error {
		body := new(model.Movie)
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(400)
		}
		err := service.SaveMovie(body)
		if err != nil {
			return c.SendStatus(500)
		}
		return c.SendString("Movie saved!")
	})

	api.Post("/movie/add/watchlist", func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		userId := user.Claims.(jwt.MapClaims)["id"].(string)
		body := new(model.Watchlist)
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(400)
		}
		body.UserId = userId
		service.AddWatchlist(body)
		return c.SendString("Movie saved in watchlist!")
	})

	api.Post("/genre/add", func(c *fiber.Ctx) error {
		body := new(model.Genre)
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(400)
		}
		err := service.SaveGenre(body)
		if err != nil {
			return c.SendStatus(500)
		}
		return c.SendString("Genre saved!")
	})
}
