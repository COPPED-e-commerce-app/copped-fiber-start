package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func InitServer() {
	app := initApp()
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

func initApp() *fiber.App {
	f := fiber.New()
	f.Use(cors.New())
	f.Use(logger.New())
	initRoutes(f)
	return f
}

func initRoutes(f *fiber.App) {
	f.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("ok")
	})
}
