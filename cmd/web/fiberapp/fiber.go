package fiberapp

import (
	"log/slog"
	"github.com/gofiber/fiber/v2"
)

func StartFiber(port string) {
	app := fiber.New()

	// Define routes
	app.Get("/new", func(c *fiber.Ctx) error {
		return c.SendString("FIBER: New route")
	})

	slog.Info("[Fiber] Listening on :" + port)
	app.Listen(":" + port)
}
