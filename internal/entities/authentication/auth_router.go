package authentication

import (
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/people"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppFiber struct {
	DB *gorm.DB
}

func RegisterRoutes(app *fiber.App, db *gorm.DB) {

	repo := people.NewRepository(db)
	service := people.NewService(repo)
	controller := NewPeopleController(service)
	
 	app.Get("/fiber",     controller.ShowFiber).Name("ShowFiber")
	app.Get("/login",     controller.ShowLogin).Name("ShowLogin")
	app.Post("/login",    controller.HandleLogin).Name("HandleLogin")
	app.Get("/welcome",   controller.HandleWelcome).Name("HandleWelcome")
	app.Get("/logout",    controller.HandleLogout).Name("HandleLogouthan")	
	app.Get("/logon",     controller.ShowLogon).Name("ShowLogon")	
	app.Post("/register", controller.HandleRegister).Name("HandleRegister")

}
