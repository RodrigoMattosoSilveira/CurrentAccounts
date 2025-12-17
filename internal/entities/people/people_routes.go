package people

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(r *fiber.App, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	controller := NewController(service)

		r.Get("/people", controller.Index)
		r.Get("/people/new", controller.New)
		r.Post("/people", controller.Create)
	r.Get("/people/:id", controller.Show)
	r.Post("/people/:id", controller.Update)
	r.Post("/people/:id/delete", controller.Delete)
}
