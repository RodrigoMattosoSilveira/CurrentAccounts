package ginapp

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/database"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/authentication"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/people"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/server"
)

func StartGin(port string) {
	// start the server
	router := server.SetupRouter()

	router.GET("/old", func(c *gin.Context) {
		c.String(http.StatusOK, "GIN: Old route")
	})

	// Define the routes
	authentication.RegisterRoutes(router, database.DB)
	people.RegisterRoutes(router, database.DB)

	slog.Info("[Gin] Listening on", "port", port)
	router.Run(":" + port)
}
