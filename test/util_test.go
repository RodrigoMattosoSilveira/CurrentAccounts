package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
)

// Helper: login and return cookie header string.
func LoginTestHelperGin(t *testing.T, r *gin.Engine, email, password string) *httptest.ResponseRecorder{
	t.Helper()

	form := url.Values{}
	form.Add("email", email)
	form.Add("password", password)

	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}


// Helper: login and return cookie header string.
func LoginTestHelperFiber(t *testing.T, r *fiber.App, email, password string) *http.Response {
	t.Helper()

	form := url.Values{}
	form.Add("email", email)
	form.Add("password", password)

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Perform the request using Fiber's built-in Test method
	resp, err := r.Test(req, -1); // -1 disables request timeout
	if err != nil {
		t.Fatalf("fiber login request failed: %v", err)
	}
	return resp
}