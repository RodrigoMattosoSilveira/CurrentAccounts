package test

import (
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestShowFiber(t *testing.T) {
	
	var testCases = []TestCase{
		{Name: "Show Fiber Test", Rest: "GET", Path: "/fiber"},
	}

	router := SetupFiberTests(t)
	// setupAuthenticationTests(t, router, db)
	for _, tc := range testCases {
		log.Println("Running TestShowFiber: ", tc.Name)
		t.Run(tc.Name, func(t *testing.T) {
			// Use the reusable helper to perform the golden file test
			
			AssertGoldenFileFiber(t, router, tc.Rest, tc.Path, tc.Name, nil)
		})
	}
}
func TestShowLogin(t *testing.T) {
	
	var testCases = []TestCase{
		{Name: "Login Page Test", Rest: "GET", Path: "/login"},
		// {Name: "Logon Page Test", Rest: "GET", Path: "/logon"},
	}

	router := SetupFiberTests(t)
	// setupAuthenticationTests(t, router, db)
	for _, tc := range testCases {
		log.Println("Running TestShowLogin: ", tc.Name)
		t.Run(tc.Name, func(t *testing.T) {
			// Use the reusable helper to perform the golden file test
			
			AssertGoldenFileFiber(t, router, tc.Rest, tc.Path, tc.Name, nil)
		})
	}
}
func TestHandleLogin (t *testing.T) {
	router := SetupFiberTests(t)
	resp := LoginTestHelperFiber(t, router, "murilo.anderson.souza@img.com.br", "Rrqmss1#")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "expected status OK after login")
}

func TestHandleLoginInvalidUserName (t *testing.T) {
	router := SetupFiberTests(t)
	resp := LoginTestHelperFiber(t, router, "murilo.anderson@img.com.br", "Rrqmss1!")
	assert.Equal(t, http.StatusUnauthorized,resp.StatusCode, "expected status OK after login")
}

func TestHandleLoginInvalidPassword (t *testing.T) {
	router := SetupFiberTests(t)
	resp := LoginTestHelperFiber(t, router, "murilo.anderson.souza@img.com.br", "Rrqmss1!")
	assert.Equal(t, http.StatusUnauthorized,resp.StatusCode, "expected status OK after login")
}

func TestHandleLogout(t *testing.T) {
	router := SetupFiberTests(t)
	resp := LoginTestHelperFiber(t, router, "murilo.anderson.souza@img.com.br", "Rrqmss1#")
	assert.Equal(t, http.StatusOK,resp.StatusCode, "expected status OK after login")

	tc := testCaseHelper("Logout Page Test", "GET", "/logout")
	t.Run(tc.Name, func(t *testing.T) {
		AssertGoldenFileFiber(t, router, tc.Rest, tc.Path, tc.Name, nil)
	})
}