package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestPeopleControllerIndex(t *testing.T) {
	router := SetupFiberTests(t)
	
	var testCases = []TestCase{
		{Name: "People Controller Index", Rest: http.MethodGet, Path: "/people"},
	}

	// setupAuthenticationTests(t, router, db)
	for _, tc := range testCases {
		log.Println("Running TestPeopleControllerIndex: ", tc.Name)
		t.Run(tc.Name, func(t *testing.T) {
			// TODO Split AssertGoldenFileFiber: 
			// 1. buildString () string 
			// 2. executeRequest (string) resp 
			// 3. compare with golden file, 
			// This will enable me to test the content of the string, like does is has the <form> tag

			// Create a new HTTP request
			req := httptest.NewRequest(http.MethodGet, "/people", nil)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			// Perform the request using Fiber's built-in Test method
			resp, err := router.Test(req, -1); // -1 disables request timeout
			if err != nil {
				t.Fatalf("fiber login request failed: %v", err)
			}
			assert.Equal(t, http.StatusOK, resp.StatusCode, "expected status OK after getting people")

			AssertGoldenFileFiberNew(t, "People Controller Index", resp)

		})
	}
}