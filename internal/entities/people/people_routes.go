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
	// Edit Name
	r.Get("/person/:id/name",  controller.EditName)
	r.Post("/person/:id/name", controller.UpdateName)
	r.Get("/person/:id/name/esc",  controller.EscUpdateName)
	// Edit Address
	r.Get("/person/:id/address",  controller.EditAddress)
	r.Post("/person/:id/address", controller.UpdateAddress)
	r.Get("/person/:id/address/esc",  controller.EscUpdateAddress)
	// Edit Cell
	r.Get("/person/:id/cell",  controller.EditCell)
	r.Post("/person/:id/cell", controller.UpdateCell)
	r.Get("/person/:id/cell/esc",  controller.EscUpdateCell)
	// Edit Email
	r.Get("/person/:id/email",  controller.EditEmail)
	r.Post("/person/:id/email", controller.UpdateEmail)
	r.Get("/person/:id/email/esc",  controller.EscUpdateEmail)
	// Edit Role
	r.Get("/person/:id/role",  controller.EditRole)
	r.Post("/person/:id/role", controller.UpdateRole)
	r.Get("/person/:id/role/esc",  controller.EscUpdateRole)

	r.Post("/people/:id/delete", controller.Delete)
}
