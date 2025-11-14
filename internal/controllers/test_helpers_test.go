package controllers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
 )

 type TestCase struct {
	name string
	path  string
}
 
var testCases = []TestCase{
	{"Home Page Test", "/"},
}
// sanitizeFilename creates a safe filename from a test case name.
func sanitizeFilename(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "_")
	name = strings.ReplaceAll(name, "/", "_")
	reg := regexp.MustCompile(`[^a-z0-9_-]`)
	return reg.ReplaceAllString(name, "")
}

// assertGoldenFile performs a golden file test for a given Gin handler.
// It makes a request to the specified path and compares the response body
// to the content of a golden file.
func assertGoldenFile(t *testing.T, router *gin.Engine, method, path string, testName string) {
	// Create the HTTP request
	req, err := http.NewRequest(method, path, nil )
	require.NoError(t, err)

	// Use the response recorder to capture the response
	w := httptest.NewRecorder( )
	router.ServeHTTP(w, req)

	// Assert that the request was successful
	require.Equal(t, http.StatusOK, w.Code, "Expected HTTP status 200" )

	// Get the actual HTML response body
	actualHTML := w.Body.String()

	// Generate the golden file path from the test name
	sanitizedName := sanitizeFilename(testName)
	goldenFileName := sanitizedName + ".golden"
	goldenFilePath := filepath.Join("testdata", goldenFileName)

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

func SetupTestRouter(route string, routeHandler  gin.HandlerFunc) *gin.Engine {
	r := gin.New()
	r.GET(route, routeHandler)

	return r
}
