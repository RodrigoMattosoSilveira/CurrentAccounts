package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

)

func TestHHelloHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// We only need a new router. The handler itself will find and parse templates.
	r := gin.New()
	r.GET("/hello", HelloHandler)

	req, _ := http.NewRequest(http.MethodGet, "/hello", nil )
	w := httptest.NewRecorder( )
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code )
	// Check for content from both the layout and the specific page.
	assert.Contains(t, w.Body.String(), "Hello, Gin!") // From layout.tmpl
	assert.Contains(t, w.Body.String(), "Welcome to the Gin web framework.") // From welcome.tmpl
}
