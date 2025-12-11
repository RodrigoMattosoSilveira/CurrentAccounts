package ginapp

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/people"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/server"
)

func StartGin(port string, db *gorm.DB) {
	// start the server
	router := server.SetupRouter()

	router.GET("/old", func(c *gin.Context) {
		c.String(http.StatusOK, "GIN: Old route")
	})

	// Define the routes
	people.RegisterRoutes(router, db)

	slog.Info("[Gin] Listening on", "port", port)
	router.Run(":" + port)
}
