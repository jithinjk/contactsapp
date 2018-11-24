package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func performAuthRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	req.SetBasicAuth("user1", "hello")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestPingRoute(t *testing.T) {
	r := setupRouter()

	w := performRequest(r, "GET", "/ping")

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestInvalidRoute(t *testing.T) {
	r := setupRouter()

	w := performRequest(r, "GET", "/invalid")

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "{\"message\":\"Page not found\",\"status\":404}", w.Body.String())
}

func TestUnauthorizedRequest(t *testing.T) {
	r := setupRouter()

	w := performRequest(r, "GET", "/api/v1/contacts/all")

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, "Basic realm=\"Authorization Required\"", w.Header().Get("Www-Authenticate"))
}

// TODO: fix it
func TestAuthorizedRequest(t *testing.T) {
	r := setupRouter()

	w := performAuthRequest(r, "GET", "/api/v1/contacts/all")
	assert.Equal(t, 200, w.Code)
}
