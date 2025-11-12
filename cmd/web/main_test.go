package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200 but got %d", w.Code)
	}

	body := w.Body.String()
	if !strings.Contains(body, "Welcome to the Gin web framework.") {
		t.Errorf("Expected body to contain 'Welcome to the Gin web framework.' but got %s", body)
	}
}

// func TestHelloHandlerFail(t *testing.T) {
// 	router := setupRouter()

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/hello", nil)
// 	router.ServeHTTP(w, req)

// 	if w.Code != http.StatusOK {
// 		t.Fatalf("Expected status 200 but got %d", w.Code)
// 	}

// 	body := w.Body.String()
// 	if !strings.Contains(body, "Welcome to the Gin web framework.") {
// 		t.Errorf("Expected body to contain 'Welcome to the Gin web framework.' but got %s", body)
// 	}
// }