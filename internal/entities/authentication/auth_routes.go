package authentication

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}
func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	// adapt converts a handler that returns an error into a gin.HandlerFunc
	// adapt := func(h func(*gin.Context) error) gin.HandlerFunc {
	// 	return func(c *gin.Context) {
	// 		if err := h(c); err != nil {
	// 			// customize error handling as needed (status code, logging, etc.)
	// 			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	// 		}
	// 	}
	// }

	controller := &App {
		DB: db,
	}
	
	r.GET("/", controller.ShowLogin)
	r.POST("/login", controller.HandleLogin)
	r.GET("/welcome", controller.HandleWelcome)
	r.GET("/logon", controller.ShowLogon)
	r.POST("/logon", controller.HandleLogon)
	r.POST("/logout", controller.HandleLogout)
	r.GET("/newpwd", controller.HandleNewPwd)
}
