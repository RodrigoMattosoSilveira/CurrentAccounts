package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities"

)

// HomeHandler handles the home route
func HelloHandler(c *gin.Context) {
utilities.Render(c, "hello/hello.tmpl", gin.H{
		"Title": "Hello, Gin!",
    	"Body":  "Welcome to the Gin web framework.",
	})
}