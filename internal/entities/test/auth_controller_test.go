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
			
			assertGoldenFile(t, router, tc.rest, tc.path, tc.name, nil)
		})
	}
}

// Helper: login and return cookie header string.
func loginTestHelper(t *testing.T, r *gin.Engine, email, password string) *httptest.ResponseRecorder{
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

func TestHandleLogin (t *testing.T) {
	router, db := setupTests(t)
	setupAuthenticationTests(t, router, db)

	w := loginTestHelper(t, router, "murilo.anderson.souza@img.com.br", "Rrqmss1#")
	assert.Equal(t, http.StatusOK, w.Code, "expected status OK after login")
}
func TestHandleLoginInvalidEmail (t *testing.T) {
	router, db := setupTests(t)
	setupAuthenticationTests(t, router, db)

	w := loginTestHelper(t, router, "muriloooo.anderson.souza@img.com.br", "Rrqmss1#")
	// if cookie == "" {
	// 	t.Fatalf("expected cookie after login")
	// }
	assert.Equal(t, http.StatusUnauthorized, w.Code, "expected status OK after login")
}
func TestHandleLoginInvalidPassword (t *testing.T) {
	router, db := setupTests(t)
	setupAuthenticationTests(t, router, db)

	w := loginTestHelper(t, router, "murilo.anderson.souza@img.com.br", "Rrqmss1!")
	assert.Equal(t, http.StatusUnauthorized, w.Code, "expected status OK after login")
}
func TestShowLogonPost(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router, db:= setupTests(t)
	setupAuthenticationTests(t, router, db)

	name := "Igor Gabriel Marcos Vinicius Rezend"
	address := "Rua Águia, 743"
	email := "igor-rezende72@fertility.com.br"	
	cell := "(11) 99826-4206"	
	password := "Rrqmss1#"

	form := url.Values{}
	form.Add("fullname", name)
	form.Add("address", address)
	form.Add("email", email)
	form.Add("cell", cell)
	form.Add("password", password)

	var tc TestCase
	tc.name =  "Logon Post Valid New Person"
	tc.rest = "POST"
	tc.path = "/register"

	log.Println("Running TestShowLogin: ", tc.name)
	t.Run(tc.name, func(t *testing.T) {
		assertGoldenFile(t, router, tc.rest, tc.path, tc.name, strings.NewReader(form.Encode()))
	})
}
func TestShowLogonPostHaveAccouunt(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	router, db:= setupTests(t)
	setupAuthenticationTests(t, router, db)

	var tc TestCase
	tc.name =  "Logon Post Have Accouunt"
	tc.rest = "GET"
	tc.path = "/login"

	log.Println("Running TestShowLogin: ", tc.name)
	t.Run(tc.name, func(t *testing.T) {
		assertGoldenFile(t, router, tc.rest, tc.path, tc.name, nil)
	})
}
func TestHandleLogout(t *testing.T) {
	router, db := setupTests(t)
	setupAuthenticationTests(t, router, db)

	w := loginTestHelper(t, router, "murilo.anderson.souza@img.com.br", "Rrqmss1#")
	assert.Equal(t, http.StatusOK, w.Code, "expected status OK after login")

	tc := testCaseHelper("Logout Page Test", "GET", "/logout")
	t.Run(tc.name, func(t *testing.T) {
		assertGoldenFile(t, router, tc.rest, tc.path, tc.name, nil)
	})
}

func testCaseHelper(name string, rest string, path string) *TestCase {
	var tc TestCase
	tc.name =  name
	tc.rest = rest
	tc.path = path
	return &tc
}