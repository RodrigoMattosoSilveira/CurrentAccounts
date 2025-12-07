package fiberapp

import (
	"log/slog"
	"github.com/gofiber/fiber/v2"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/database"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/authentication"
)

func StartFiber(port string) {
	router := fiber.New()

	// // Define routes
	// app.Get("/new", func(c *fiber.Ctx) error {
	// 	return c.SendString("FIBER: New route")
	// })

	authentication.RegisterRoutesFiber(router, database.DB)
	slog.Info("[Fiber] Listening on :" + port)
	router.Listen(":" + port)
}
