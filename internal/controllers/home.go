package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities"
)

// HomeHandler handles the home route
func HomeHandler(c *gin.Context) {
	utilities.Render(c, "home/welcome.tmpl", gin.H{
		"Tenant": "MC",
		"Host":   "Madone Logistics",
	})
}
