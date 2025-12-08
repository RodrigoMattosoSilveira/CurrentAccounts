package test

import (
	"log"
	"testing"
)


func TestShowFiber(t *testing.T) {
	
	var testCases = []TestCase{
		{Name: "Show Fiber Test", Rest: "GET", Path: "/fiber"},
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