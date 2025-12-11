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
	assert.Equal(t, http.StatusUnprocessableEntity,resp.StatusCode, "expected status OK after login")
}

func TestHandleLoginInvalidPassword (t *testing.T) {
	router := SetupFiberTests(t)
	resp := LoginTestHelperFiber(t, router, "murilo.anderson.souza@img.com.br", "Rrqmss1!")
	assert.Equal(t, http.StatusUnprocessableEntity,resp.StatusCode, "expected status OK after login")
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

func TestShowLogon(t *testing.T) {
	router := SetupFiberTests(t)
	
	var testCases = []TestCase{
		{Name: "Logon Page Test", Rest: "GET", Path: "/logon"},
	}

	// setupAuthenticationTests(t, router, db)
	for _, tc := range testCases {
		log.Println("Running TestShowLogon: ", tc.Name)
		t.Run(tc.Name, func(t *testing.T) {
			// TODO Split AssertGoldenFileFiber: 
			// 1. buildString () string 
			// 2. executeRequest (string) resp 
			// 3. compare with golden file, 
			// This will enable me to test the content of the string, like does is has the <form> tag
			AssertGoldenFileFiber(t, router, tc.Rest, tc.Path, tc.Name, nil)
		})
	}
}

func TestRegisterPerson(t *testing.T) {
	router := SetupFiberTests(t)

	var formStruc LogonForm = LogonForm{	
		name:     "Igor Gabriel Marcos Vinicius Rezend",
		address:  "Rua Águia, 743",
		email:    "igor-rezende72@fertility.com.br",	
		cell :    "(11) 99826-4206",	
		password: "Rrqmss1#",
	}

	resp := LogonTestHelperFiber(t, router, &formStruc)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "expected status OK after logon")

	AssertGoldenFileFiberNew(t, "Register Person", resp)
}
func TestRegisterPersonInvalidName(t *testing.T) {
	// very rudimentary test to check empty email handling; add validate.validator later
	router := SetupFiberTests(t)

	var formStruc LogonForm = LogonForm{	
		name:     "",
		address:  "Rua Águia, 743",
		email:    "igor-rezende72@fertility.com.br",	
		cell :    "(11) 99826-4206",	
		password: "Rrqmss1#",
	}

	resp := LogonTestHelperFiber(t, router, &formStruc)
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode, "expected status OK after logon")

	AssertGoldenFileFiberNew(t, "Register Person Invalid Name", resp)
}
func TestRegisterPersonInvalidEmail(t *testing.T) {
	// very rudimentary test to check empty email handling; add validate.validator later
	router := SetupFiberTests(t)

	var formStruc LogonForm = LogonForm{	
		name:     "Igor Gabriel Marcos Vinicius Rezend",
		address:  "Rua Águia, 743",
		email:    "",	
		cell :    "(11) 99826-4206",	
		password: "Rrqmss1#",
	}

	resp := LogonTestHelperFiber(t, router, &formStruc)
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode, "expected status OK after logon")

	AssertGoldenFileFiberNew(t, "Register Person Invalid Email", resp)
}
func TestRegisterPersonInvalidAddress(t *testing.T) {
	// very rudimentary test to check empty email handling; add validate.validator later
	router := SetupFiberTests(t)

	var formStruc LogonForm = LogonForm{	
		name:     "Igor Gabriel Marcos Vinicius Rezend",
		address:  "",
		email:    "igor-rezende72@fertility.com.br",	
		cell :    "(11) 99826-4206",	
		password: "Rrqmss1#",
	}

	resp := LogonTestHelperFiber(t, router, &formStruc)
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode, "expected status OK after logon")

	AssertGoldenFileFiberNew(t, "Register Person Invalid Address", resp)
}
func TestRegisterPersonInvalidPassword(t *testing.T) {
	// very rudimentary test to check empty email handling; add validate.validator later
	router := SetupFiberTests(t)

	var formStruc LogonForm = LogonForm{	
		name:     "Igor Gabriel Marcos Vinicius Rezend",
		address:  "Rua Águia, 743",
		email:    "igor-rezende72@fertility.com.br",	
		cell :    "(11) 99826-4206",	
		password: "",
	}

	resp := LogonTestHelperFiber(t, router, &formStruc)
	assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode, "expected status OK after logon")

	AssertGoldenFileFiberNew(t, "Register Person Invalid Password", resp)
}

func TestRegisterPersonAlreadyRegistered(t *testing.T) {
	// very rudimentary test to check empty email handling; add validate.validator later
	router := SetupFiberTests(t)

	var formStruc LogonForm = LogonForm{	
		name:     "Levi Rodrigo Diogo Araújo",
		address:  "Rua Itajubá, 669",
		email:    "thiago_sergio_novaes@dpi.indl.com.br",	
		cell :    "(31) 99637-4395",	
		password: "Rrqmss1#",
	}

	resp := LogonTestHelperFiber(t, router, &formStruc)
	assert.Equal(t, http.StatusConflict, resp.StatusCode, "expected status OK after logon")

	AssertGoldenFileFiberNew(t, "Register Person Already Registered", resp)
}