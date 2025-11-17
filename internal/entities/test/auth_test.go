package test

import (
	"log"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/entities/authentication"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/config"
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
		{"Login Page Test", "/"},
		{"Logon Page Test", "/"},
	}

	_, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db := setupTestDB(t)
	router := setupTestRouter(t, db)

	for _, tc := range testCases {
		log.Println("Runninf TestShowLogin: ", tc.name)
		t.Run(tc.name, func(t *testing.T) {
			// Use the reusable helper to perform the golden file test
			
			assertGoldenFile(t, router, "GET", tc.path, tc.name)
		})
	}
}