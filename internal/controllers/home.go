package controllers

import (
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities" // Use your module name

	"github.com/gin-gonic/gin"
)

func HomeHandler(c *gin.Context) {
	// We need the layout and the specific welcome page.
	// The paths are relative to the 'templates' directory.
	templateFiles := []string{
		"root/layout.tmpl",
		"root/authentication/login.tmpl",
	}

	// Call our custom renderer.
	// The name "layout.tmpl" tells the template engine which template definition to execute first.
	utilities.RenderTemplate(c, "layout", gin.H{
		"Tenant": "MC",
		"Host":   "Madone Logistics",
	}, templateFiles...)
}
