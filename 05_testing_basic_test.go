package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func performRequest(t *testing.T, app *fiber.App, method, path string, body io.Reader, headers map[string]string) *http.Response {
	req := httptest.NewRequest(method, path, body)

	if headers != nil {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}

	resp, err := app.Test(req, -1)
	require.NoError(t, err, "app.Test should not return an error")

	return resp
}

func TestGetRoot(t *testing.T) {
	app := fiber.New()
	setupRoutes(app)

	resp := performRequest(t, app, "GET", "/", nil, nil)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Status code should be 200 OK")

	bodyBytes, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "Should be able to read response body")
	assert.Equal(t, "Hello, World!", string(bodyBytes), "Response body should be 'Hello, World!'")
}

func TestPostEcho(t *testing.T) {
	app := fiber.New()
	setupRoutes(app)

	requestBodyJSON := `{"message": "ping"}`
	headersJSON := map[string]string{"Content-Type": fiber.MIMEApplicationJSON}

	respJSON := performRequest(t, app, "POST", "/echo", strings.NewReader(requestBodyJSON), headersJSON)
	defer respJSON.Body.Close()

	assert.Equal(t, http.StatusOK, respJSON.StatusCode, "[JSON] Status code should be 200 OK")
	assert.Equal(t, fiber.MIMEApplicationJSON, respJSON.Header.Get(fiber.HeaderContentType), "[JSON] Content-Type should be application/json")
	bodyBytesJSON, errJSON := io.ReadAll(respJSON.Body)
	require.NoError(t, errJSON)
	assert.JSONEq(t, requestBodyJSON, string(bodyBytesJSON), "[JSON] Response body should match request body")

	requestBodyText := "Just plain text"
	headersText := map[string]string{"Content-Type": fiber.MIMETextPlain}

	respText := performRequest(t, app, "POST", "/echo", strings.NewReader(requestBodyText), headersText)
	defer respText.Body.Close()

	assert.Equal(t, http.StatusOK, respText.StatusCode, "[Text] Status code should be 200 OK")
	assert.Equal(t, fiber.MIMETextPlain, respText.Header.Get(fiber.HeaderContentType), "[Text] Content-Type should be text/plain")
	bodyBytesText, errText := io.ReadAll(respText.Body)
	require.NoError(t, errText)
	assert.Equal(t, requestBodyText, string(bodyBytesText), "[Text] Response body should match request body")

}

func TestGetUserByName(t *testing.T) {
	app := fiber.New()
	setupRoutes(app)

	userName := "alice"
	path := "/users/" + userName

	resp := performRequest(t, app, "GET", path, nil, nil)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Status code should be 200 OK")
	assert.Contains(t, resp.Header.Get(fiber.HeaderContentType), fiber.MIMEApplicationJSON, "Content-Type should contain application/json")

	var result map[string]string
	err := json.NewDecoder(resp.Body).Decode(&result)
	require.NoError(t, err, "Should be able to decode JSON response")

	assert.Equal(t, userName, result["user"], "Value of 'user' key in JSON should match path parameter")
}

func TestNotFound(t *testing.T) {
	app := fiber.New()

	resp := performRequest(t, app, "GET", "/path/that/does/not/exist", nil, nil)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode, "Status code should be 404 Not Found")

	bodyBytes, _ := io.ReadAll(resp.Body)
	assert.Contains(t, string(bodyBytes), "Cannot GET /path/that/does/not/exist", "Default 404 body should contain error message")
}
