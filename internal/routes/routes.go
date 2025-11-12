package routes	

import (
	"github.com/gin-gonic/gin"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/controllers"
)

func SetupRoutes(router *gin.Engine) {
	HelloRoutes(router)
	AuthenticationRoutes(router)
}

func HelloRoutes(router *gin.Engine	) {
  router.GET("/hello", controllers.HomeHandler)
}

func AuthenticationRoutes(router *gin.Engine	) {
  router.GET("/", controllers.HomeHandler)
}