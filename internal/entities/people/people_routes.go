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
	r.Get("/person/:id/name", controller.EditName)
	r.Get("/person/:id/address", controller.EditAddress)
	r.Get("/person/:id/email", controller.EditEmail)
	r.Get("/person/:id/cell", controller.EditCell)
	r.Get("/person/:id/role", controller.EditRole)
	r.Post("/person/:id/name", controller.UpdateName)
	r.Post("/person/:id/address", controller.UpdateAddress)
	r.Post("/person/:id/email", controller.UpdateEmail)
	r.Post("/person/:id/cell", controller.UpdateCell)
	r.Post("/person/:id/role", controller.UpdateRole)

	r.Post("/people/:id/delete", controller.Delete)
}
