package api

import (
	jwt_config "github.com/CarlosPGit/golang_sface/internal/platform/jwtConfig"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const prefix = "/sfv1"

func MapUrls(app *fiber.App, handlers *InitConf) {

	app.Post("/login", handlers.userHandler.Login)
	app.Post("/register", handlers.userHandler.Register)

	apiGroup := app.Group(prefix)
	apiGroup.Use(jwt_config.NewJWT())
	apiGroup.Get("/", func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["name"].(string)
		return c.SendString("Welcome " + name)
	})
	app.Get("/a", func(c *fiber.Ctx) error {
		return c.SendString("access")
	})
	apiGroup.Get("/b", func(c *fiber.Ctx) error {
		return c.SendString("access")
	})
}
