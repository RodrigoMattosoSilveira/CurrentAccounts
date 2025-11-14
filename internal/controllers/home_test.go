package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
 )

func TestHomeHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	// We only need a new router. The handler itself will find and parse templates.
	r := gin.New()
	r.GET("/", HomeHandler)

	req, _ := http.NewRequest(http.MethodGet, "/", nil )
	w := httptest.NewRecorder( )
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code )
	// Check for content from both the layout and the specific page.
	assert.Contains(t, w.Body.String(), "MC") // From layout.tmpl
	assert.Contains(t, w.Body.String(), "Madone Logistics") // From welcome.tmpl
}
