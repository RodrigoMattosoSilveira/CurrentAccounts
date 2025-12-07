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

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/test"
)


func TestShowLogin(t *testing.T) {
	
	var testCases = []test.TestCase{
		{Name: "Login Page Test", Rest: "GET", Path: "/login"},
		{Name: "Logon Page Test", Rest: "GET", Path: "/logon"},
	}

	router := test.SetupGinTests(t)
	// setupAuthenticationTests(t, router, db)
	for _, tc := range testCases {
		log.Println("Running TestShowLogin: ", tc.Name)
		t.Run(tc.Name, func(t *testing.T) {
			// Use the reusable helper to perform the golden file test
			
			test.AssertGoldenFile(t, router, tc.Rest, tc.Path, tc.Name, nil)
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
	router := test.SetupGinTests(t)
	w := loginTestHelper(t, router, "murilo.anderson.souza@img.com.br", "Rrqmss1#")
	assert.Equal(t, http.StatusOK, w.Code, "expected status OK after login")
}
func TestHandleLoginInvalidEmail (t *testing.T) {
	router := test.SetupGinTests(t)

	w := loginTestHelper(t, router, "muriloooo.anderson.souza@img.com.br", "Rrqmss1#")
	// if cookie == "" {
	// 	t.Fatalf("expected cookie after login")
	// }
	assert.Equal(t, http.StatusUnauthorized, w.Code, "expected status OK after login")
}
func TestHandleLoginInvalidPassword (t *testing.T) {
	router := test.SetupGinTests(t)

	w := loginTestHelper(t, router, "murilo.anderson.souza@img.com.br", "Rrqmss1!")
	assert.Equal(t, http.StatusUnauthorized, w.Code, "expected status OK after login")
}
func TestShowLogonPost(t *testing.T) {
	router := test.SetupGinTests(t)

	Name := "Igor Gabriel Marcos Vinicius Rezend"
	address := "Rua Águia, 743"
	email := "igor-rezende72@fertility.com.br"	
	cell := "(11) 99826-4206"	
	password := "Rrqmss1#"

	form := url.Values{}
	form.Add("fullName", Name)
	form.Add("address", address)
	form.Add("email", email)
	form.Add("cell", cell)
	form.Add("password", password)

	var tc test.TestCase
	tc.Name =  "Logon Post Valid New Person"
	tc.Rest = "POST"
	tc.Path = "/register"

	log.Println("Running TestShowLogin: ", tc.Name)
	t.Run(tc.Name, func(t *testing.T) {
		test.AssertGoldenFile(t, router, tc.Rest, tc.Path, tc.Name, strings.NewReader(form.Encode()))
	})
}
func TestShowLogonPostHaveAccouunt(t *testing.T) {
	router := test.SetupGinTests(t)

	var tc test.TestCase
	tc.Name =  "Logon Post Have Accouunt"
	tc.Rest = "GET"
	tc.Path = "/login"

	log.Println("Running TestShowLogin: ", tc.Name)
	t.Run(tc.Name, func(t *testing.T) {
		test.AssertGoldenFile(t, router, tc.Rest, tc.Path, tc.Name, nil)
	})
}
func TestHandleLogout(t *testing.T) {
	router := test.SetupGinTests(t)

	w := loginTestHelper(t, router, "murilo.anderson.souza@img.com.br", "Rrqmss1#")
	assert.Equal(t, http.StatusOK, w.Code, "expected status OK after login")

	tc := testCaseHelper("Logout Page Test", "GET", "/logout")
	t.Run(tc.Name, func(t *testing.T) {
		test.AssertGoldenFile(t, router, tc.Rest, tc.Path, tc.Name, nil)
	})
}

func testCaseHelper(Name string, Rest string, Path string) *test.TestCase {
	var tc test.TestCase
	tc.Name =  Name
	tc.Rest = Rest
	tc.Path = Path
	return &tc
}