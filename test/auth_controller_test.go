package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

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

func TestShowLogonPost(t *testing.T) {
	router := SetupGinTests(t)

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

	var tc TestCase
	tc.Name =  "Logon Post Valid New Person"
	tc.Rest = "POST"
	tc.Path = "/register"

	log.Println("Running TestShowLogonPost Gin: ", tc.Name)
	t.Run(tc.Name, func(t *testing.T) {
		AssertGoldenFile(t, router, tc.Rest, tc.Path, tc.Name, strings.NewReader(form.Encode()))
	})
}
func TestShowLogonPostHaveAccouunt(t *testing.T) {
	router := SetupGinTests(t)

	var tc TestCase
	tc.Name =  "Logon Post Have Accouunt"
	tc.Rest = "GET"
	tc.Path = "/logon"

	log.Println("Running TestShowLogonPostHaveAccouunt Gin: ", tc.Name)
	t.Run(tc.Name, func(t *testing.T) {
		AssertGoldenFile(t, router, tc.Rest, tc.Path, tc.Name, nil)
	})
}

func testCaseHelper(Name string, Rest string, Path string) *TestCase {
	var tc TestCase
	tc.Name =  Name
	tc.Rest = Rest
	tc.Path = Path
	return &tc
}