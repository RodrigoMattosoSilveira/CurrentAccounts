package fiberapp

import (
	"log/slog"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/authentication"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/people"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities"
)

func StartFiber(port string, db *gorm.DB) {
	router := fiber.New()
	router.Use(logger.New(logger.Config{
		Format:     "[${time}] ${ip} - ${method} ${path} - ${status} ${latency}\n",
		TimeFormat: time.RFC3339,
		TimeZone:   "Local",
		Output:     os.Stdout, // Change to os.Stdout if you want console logs
	}))
	store := session.New()
	router.Use(utilities.WithSession(store))
	
	// // Define routes
	// app.Get("/new", func(c *fiber.Ctx) error {
	// 	return c.SendString("FIBER: New route")
	// })

	authentication.RegisterRoutes(router, db)
	people.RegisterRoutes(router, db)

	slog.Info("[Fiber] Listening on :" + port)
	router.Listen(":" + port)
}