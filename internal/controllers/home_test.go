package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/server"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities"
)

func TestHomeHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testRoute := "/"
	router := server.SetupRouter()
	router.GET(testRoute, HomeHandler)

	req, _ := http.NewRequest("GET", testRoute, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Load expected output from file
	fileExpect := utilities.GetProjectRoot() + "/static/testData/controllers/home/home_expected.tmpl"
	expectedBytes, err := os.ReadFile(fileExpect)
	assert.NoError(t, err, "could not read expected HTML file")

	expected := strings.TrimSpace(string(expectedBytes))
	actual := strings.TrimSpace(w.Body.String())

	// Optional: normalize whitespace for safety
	clean := func(s string) string {
		return strings.Join(strings.Fields(s), " ")
	}

	assert.Equal(t, clean(expected), clean(actual), "rendered HTML does not match expected output file")
}
