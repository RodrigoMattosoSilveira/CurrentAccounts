package authentication

import (
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/people"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppFiber struct {
	DB *gorm.DB
}

func RegisterRoutesFiber(r *fiber.App, db *gorm.DB) {

	repo := people.NewRepository(db)
	service := people.NewService(repo)
	controller := NewPeopleController(service)
	
 	r.Get("/fiber", controller.ShowFiber)
	r.Get("/login", controller.ShowLogin)
	r.Post("/login", controller.HandleLogin)
	r.Get("/logout", controller.HandleLogout)	
}
