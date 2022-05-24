package test

import (
	"encoding/json"
	"io"
	"jabutech/blog-restful-api/handler"
	"jabutech/blog-restful-api/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Setup router
func setupRouter() *gin.Engine {
	// Use handler
	appHealthHandler := handler.NewPingHandler()

	// Use router
	router := router.NewRouter(appHealthHandler)

	return router
}

func TestPing(t *testing.T) {
	// Use router
	router := setupRouter()

	// Create request
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/ping", nil)

	// Create recorder for recorder
	recorder := httptest.NewRecorder()

	// Run test
	router.ServeHTTP(recorder, request)

	// Get result
	response := recorder.Result()
	// Read response body json
	body, _ := io.ReadAll(response.Body)
	// Create var responseBody
	var responseBody map[string]interface{}
	// Decode json
	json.Unmarshal(body, &responseBody)

	// response status must be 200 (ok)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "pong", responseBody["status"])
}
