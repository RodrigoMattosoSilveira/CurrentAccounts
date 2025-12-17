package test

import (
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)	

// TODO build a module consisting of it, GetHTTPResponse, AssertGoldenFileFiberNew	
func GetHTTPResponse(t *testing.T, router *fiber.App, method, path string, body io.Reader) *http.Response {
	t.Helper()

	log.Println("Performing HTTP request:", method, path)
	req, err := http.NewRequest(method, path, body )
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	require.NoError(t, err)



	resp, err := router.Test(req, -1) // -1 disables request timeout
	if err != nil {
		t.Fatalf("fiber request failed: %v", err)
	}
	assert.NotNil(t, resp, "expected non-nil response")
	return resp
}