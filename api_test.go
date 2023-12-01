// Add API endpoint tests here
package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	// Initialize the API routes
	InitAPI()

	// Create a test server
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Create a request to the "/ping" endpoint
	req, _ := http.NewRequest("GET", ts.URL+"/ping", nil)

	// Perform the request and capture the response
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Could not perform request: %v", err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("Could not read response body: %v", err)
	}

	// Assert the status code and response body
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.JSONEq(t, `{"message":"pong"}`, string(body))
}
