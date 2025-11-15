package controllers

import (
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities" // Use your module name

	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	// This handler uses a different layout and its own template.
	templateFiles := []string{
		"root/simple_layout.tmpl",
		"root/hello/hello.tmpl",
	}

	// Execute the "simple_layout.tmpl" as the base.
	utilities.RenderTemplate(c, "simple_layout", gin.H{
		"Title": "Hello, Gin!",
    	"Body":  "Welcome to the Gin web framework.",
	}, templateFiles...)
}
