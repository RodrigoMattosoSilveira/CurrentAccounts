package test

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/constants"
	"github.com/RodrigoMattosoSilveira/CurrentAccounts/internal/utilities"
)

type LogonForm struct {
	name     string
	address  string
	email    string
	cell     string
	password string
}

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

func FiberRequest(app *fiber.App, method, path string, body io.Reader) (*http.Response, string) {
	req := httptest.NewRequest(method, path, body)
	resp, err := app.Test(req, -1)
	if err != nil {
		log.Fatalf("fiber test request failed: %v", err)
	}
	bodyBytes, _ := io.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	return resp, bodyString
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

// Helper: login and return cookie header string.
func LogonTestHelperFiber(t *testing.T, r *fiber.App, formStruct *LogonForm) *http.Response {
	t.Helper()

	form := url.Values{}
	form.Set("fullname", formStruct.name)
	form.Set("address", formStruct.address)
	form.Set("email", formStruct.email)
	form.Set("cell", formStruct.cell)
	form.Set("password", formStruct.password)

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Perform the request using Fiber's built-in Test method
	resp, err := r.Test(req, -1); // -1 disables request timeout
	if err != nil {
		t.Fatalf("fiber login request failed: %v", err)
	}
	return resp
}

func AssertGoldenFileFiber(t *testing.T, app *fiber.App, method, path string, testName string, body io.Reader) {
	// Create the HTTP request
	req, err := http.NewRequest(method, path, body )
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	require.NoError(t, err)

	// Use the response recorder to capture the response
	resp, _ := app.Test(req)

	// Assert that the request was successful
	require.Equal(t, http.StatusOK, resp.StatusCode, method + path + ": Expected HTTP status 200" )

	// Get the actual HTML response body
	actualHTMLBytes, _ := io.ReadAll(resp.Body)
	actualHTML := string(actualHTMLBytes)


	// Generate the golden file path from the test name
	projectRoot, err := utilities.FindProjectRoot()
	if err != nil {
		log.Printf("ERROR: Failed to find project root: %v", err)
		assert.Equal(t,err, nil)
	}
	sanitizedName := SanitizeFilename(testName)
	goldenFileName := sanitizedName + ".golden"
	goldenFilePath := filepath.Join(projectRoot, string(constants.TEST_GOLDEN_FOLDER), goldenFileName)

	// Update logic for golden files
	if os.Getenv("UPDATE_GOLDEN_FILES") != "" {
		t.Logf("Updating golden file: %s", goldenFilePath)
		err := os.MkdirAll(filepath.Dir(goldenFilePath), 0755)
		require.NoError(t, err)
		err = os.WriteFile(goldenFilePath, []byte(actualHTML), 0644)
		require.NoError(t, err)
	}

	// Read the golden file
	expectedHTML, err := os.ReadFile(goldenFilePath)
	require.NoError(t, err, "Failed to read golden file. Run with UPDATE_GOLDEN_FILES=true to create it.")

	// Compare the actual response to the golden file
	assert.Equal(t, string(expectedHTML), actualHTML)
}

func AssertGoldenFileFiberNew(t *testing.T, testName string, resp *http.Response) {

	// Get the actual HTML response body
	actualHTMLBytes, _ := io.ReadAll(resp.Body)
	actualHTML := string(actualHTMLBytes)


	// Generate the golden file path from the test name
	projectRoot, err := utilities.FindProjectRoot()
	if err != nil {
		log.Printf("ERROR: Failed to find project root: %v", err)
		assert.Equal(t,err, nil)
	}
	sanitizedName := SanitizeFilename(testName)
	goldenFileName := sanitizedName + ".golden"
	goldenFilePath := filepath.Join(projectRoot, string(constants.TEST_GOLDEN_FOLDER), goldenFileName)

	// Update logic for golden files
	if os.Getenv("UPDATE_GOLDEN_FILES") != "" {
		t.Logf("Updating golden file: %s", goldenFilePath)
		err := os.MkdirAll(filepath.Dir(goldenFilePath), 0755)
		require.NoError(t, err)
		err = os.WriteFile(goldenFilePath, []byte(actualHTML), 0644)
		require.NoError(t, err)
	}

	// Read the golden file
	expectedHTML, err := os.ReadFile(goldenFilePath)
	require.NoError(t, err, "Failed to read golden file. Run with UPDATE_GOLDEN_FILES=true to create it.")

	// Compare the actual response to the golden file
	assert.Equal(t, string(expectedHTML), actualHTML)
}

func testCaseHelper(Name string, Rest string, Path string) *TestCase {
	var tc TestCase
	tc.Name =  Name
	tc.Rest = Rest
	tc.Path = Path
	return &tc
}