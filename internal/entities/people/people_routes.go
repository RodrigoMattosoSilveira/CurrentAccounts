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
	r.Get("/people/:id/name", controller.EditName)
	r.Get("/people/:id/address", controller.EditAddress)
	r.Get("/people/:id/email", controller.EditEmail)
	r.Get("/people/:id/cell", controller.EditCell)
	r.Get("/people/:id/role", controller.EditRole)
	r.Post("/people/:id/name", controller.UpdateName)
	r.Post("/people/:id/address", controller.UpdateAddress)
	r.Post("/people/:id/email", controller.UpdateEmail)
	r.Post("/people/:id/cell", controller.UpdateCell)
	r.Post("/people/:id/role", controller.UpdateRole)

	r.Post("/people/:id", controller.Update)
	r.Post("/people/:id/delete", controller.Delete)
}
