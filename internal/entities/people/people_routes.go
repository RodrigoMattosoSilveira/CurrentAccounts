package people

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	controller := NewController(service)

	r.GET("/people", controller.Index)
	r.GET("/people/new", controller.New)
	r.POST("/people", controller.Create)
	r.GET("/people/:id", controller.Show)
	r.GET("/people/:id/edit", controller.Edit)
	r.POST("/people/:id", controller.Update)
	r.POST("/people/:id/delete", controller.Delete)
}
