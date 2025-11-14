package controllers

import (
	"testing"

	"github.com/gin-gonic/gin"
 )

// 1. Define the test cases in a table
// type TestCase struct {
// 	name string
// 	path  string
// }
 


func TestHomeHandler(t *testing.T) {
	var testCases = []TestCase{
		{"Home Page Test", "/"},
	}
	gin.SetMode(gin.TestMode)

	// We only need a new router. The handler itself will find and parse templates.
	router := SetupTestRouter("/", HomeHandler)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Use the reusable helper to perform the golden file test
			assertGoldenFile(t, router, "GET", tc.path, tc.name)
		})
	}
}
