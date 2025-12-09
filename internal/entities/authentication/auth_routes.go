package authentication

import (
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/people"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	repo := people.NewRepository(db)
	service := people.NewService(repo)
	controller := NewPeopleController(service)
	
	r.GET("/welcome", controller.HandleWelcome)
	r.GET("/logon", controller.ShowLogon)
	r.POST("/register", controller.HandleRegister)
	r.GET("/newpwd", controller.HandleNewPwd)
}
