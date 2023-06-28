package middleware_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/middleware"
	"github.com/stretchr/testify/assert"
)

const (
	InvalidId = "o id '%s' é inválido"
)

func TestIdValidationMiddleware(t *testing.T) {
	t.Run("Should not set id on context", func(t *testing.T) {
		context, recorder, router := createValidationContext("", getStringRequestInBytes)
		request, _ := http.NewRequest("GET", "/", nil)
		router.Use(middleware.IdValidation())
		router.GET("/", successHandler())

		router.ServeHTTP(recorder, request)

		_, exists := context.Get("Id")

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.False(t, exists)
		assert.False(t, context.IsAborted())
	})

	t.Run("Should have error when id is not an integer", func(t *testing.T) {
		context, recorder, router := createValidationContext("", getStringRequestInBytes)
		request, _ := http.NewRequest("GET", "/abc", nil)
		context.Request = request
		router.Use(middleware.IdValidation())
		router.GET("/:id", successHandler())

		router.ServeHTTP(recorder, request)

		var response ErrorResponse
		json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Len(t, response.Messages, 1)
		assert.Equal(t, response.Messages[0], "o id 'abc' é inválido")
	})

	t.Run("Should have success", func(t *testing.T) {
		context, recorder, router := createValidationContext("", getStringRequestInBytes)
		request, _ := http.NewRequest("GET", "/1", nil)
		context.Request = request
		router.Use(middleware.IdValidation())
		router.GET("/:id", successHandler())

		router.ServeHTTP(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
