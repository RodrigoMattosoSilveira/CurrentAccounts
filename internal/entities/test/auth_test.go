package test

import (
	"log"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/authentication"
)


// setupTestRouter loads templates and returns a gin.Engine for tests.
func setupTestRouter(t *testing.T, db *gorm.DB) *gin.Engine {
	t.Helper()

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	authentication.RegisterRoutes(r, db)

	return r
}

func TestShowLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	var testCases = []TestCase{
		{"Login Page Test", "GET", "/"},
		{"Logon Page Test", "GET", "/logon"},
	}

	router, db:= setupTests(t)
	setupAuthenticationTests(t, router, db)
	for _, tc := range testCases {
		log.Println("Runninf TestShowLogin: ", tc.name)
		t.Run(tc.name, func(t *testing.T) {
			// Use the reusable helper to perform the golden file test
			
			assertGoldenFile(t, router, tc.rest, tc.path, tc.name)
		})
	}
}