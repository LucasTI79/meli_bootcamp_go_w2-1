package middleware_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInternalErrorMiddleware(t *testing.T) {
	t.Run("Should handlers execute with success", func(t *testing.T) {
		router, lastMiddlewareWasCalled := createRouter(successHandler)
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/", nil)

		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.True(t, *lastMiddlewareWasCalled)
		assert.Equal(t, "last middleware called", recorder.Body.String())
	})

	t.Run("Should write a internal server error response when capture a panic", func(t *testing.T) {
		router, lastMiddlewareWasCalled := createRouter(errorHandler)
		recorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", "/", nil)

		router.ServeHTTP(recorder, request)
		var response ErrorResponse
		_ = json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, http.StatusInternalServerError, recorder.Code)
		assert.Len(t, response.Messages, 1)
		assert.Equal(t, "an internal error ocurred", response.Messages[0])
		assert.False(t, *lastMiddlewareWasCalled)
	})
}

func successHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}

func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		panic("error")
	}
}

func lastMiddleware(wasCalled *bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		*wasCalled = true
		c.String(http.StatusOK, "last middleware called")
	}
}

func createRouter(handler func() gin.HandlerFunc) (*gin.Engine, *bool) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.Use(middleware.InternalError())
	wasCalled := false
	router.GET("/", handler(), lastMiddleware(&wasCalled))
	return router, &wasCalled
}
