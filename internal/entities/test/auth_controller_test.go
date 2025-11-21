package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


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

// Helper: login and return cookie header string.
func loginAndGetCookie(t *testing.T, r *gin.Engine, email, password string) string {
	t.Helper()

	form := url.Values{}
	form.Add("email", email)
	form.Add("password", password)

	req, _ := http.NewRequest(http.MethodPost, "/login", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusFound {
		t.Fatalf("expected login redirect, got %d", w.Code)
	}

	cookie := w.Header().Get("Set-Cookie")
	if cookie == "" {
		t.Fatalf("expected Set-Cookie after login")
	}
	return cookie
}

func TestHandleLogin (t *testing.T) {
	router, db := setupTests(t)
	setupAuthenticationTests(t, router, db)

	cookie := loginAndGetCookie(t, router, "murilo.anderson.souza@img.com.br", "Rrqmss1#")
	// if cookie == "" {
	// 	t.Fatalf("expected cookie after login")
	// }
	assert.NotEqual(t, cookie, "", "a and c should not be equal")
}
