package middleware_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type CorrectRequest struct {
	FieldA *string `json:"field_a" binding:"required"`
	FieldB *string `json:"field_b" binding:"e164,omitempty"`
}

type WrongTypeRequest struct {
	FieldA *int `json:"field_a" binding:"required"`
	FieldB *int `json:"field_b" binding:"e164,omitempty"`
}

type MissingRequiredFieldRequest struct {
	FieldB *string `json:"field_b" validate:"e164,omitempty"`
}

type UnknownValidationTagRequest struct {
	FieldA *string `json:"field_a" binding:"lte=3"`
}

type ErrorResponse struct {
	Code     string   `json:"code"`
	Messages []string `json:"messages"`
	Status   int      `json:"status"`
}

func TestValidationMiddleware(t *testing.T) {
	fieldA := "Field A"
	fieldB := "+5500123456789"

	t.Run("Should have success on validation", func(t *testing.T) {
		request := createCorrectRequest(fieldA, fieldB)
		context, recorder := createValidationContext(request)

		middleware.Validation[CorrectRequest]()(context)
		gotRequest := context.MustGet("Request").(CorrectRequest)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, request, gotRequest)
		assert.False(t, context.IsAborted())
	})

	t.Run("Should have error when try parse a request field with a syntax error", func(t *testing.T) {
		request := createRequestWithSyntaxError()
		context, recorder := createValidationContextWithStringRequest(request)

		middleware.Validation[CorrectRequest]()(context)

		var response ErrorResponse
		json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
		assert.Len(t, response.Messages, 1)
		assert.Contains(t, response.Messages[0], "erro de sintaxe na posição")
		assert.True(t, context.IsAborted())
	})

	t.Run("Should have error when try parse a request field with a wrong type", func(t *testing.T) {
		request := createWrongTypeRequest(1, 1)
		context, recorder := createValidationContext(request)

		middleware.Validation[CorrectRequest]()(context)

		var response ErrorResponse
		json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
		assert.Len(t, response.Messages, 1)
		assert.Equal(t, "o campo 'field_a' deve ser 'string'", response.Messages[0])
		assert.True(t, context.IsAborted())
	})

	t.Run("Should have error when try parse a request with a missing required field", func(t *testing.T) {
		request := createMissingRequiredFieldRequest(fieldB)
		context, recorder := createValidationContext(request)

		middleware.Validation[CorrectRequest]()(context)

		var response ErrorResponse
		json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
		assert.Len(t, response.Messages, 1)
		assert.Equal(t, "'field_a' é obrigatório", response.Messages[0])
		assert.True(t, context.IsAborted())
	})

	t.Run("Should have error when try parse a request with a wrong phone format", func(t *testing.T) {
		request := createCorrectRequest(fieldA, "Phone")
		context, recorder := createValidationContext(request)

		middleware.Validation[CorrectRequest]()(context)

		var response ErrorResponse
		json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
		assert.Len(t, response.Messages, 1)
		assert.Equal(t, "'field_b' precisa estar no formato +<country_code><zone_code><phone_number> sem espaços ou caracteres especiais, por exemplo: +5500123456789", response.Messages[0])
		assert.True(t, context.IsAborted())
	})

	t.Run("Should have error when try parse a request with a unknown validation tag", func(t *testing.T) {
		request := createUnknownValidationTagRequest(fieldA)
		context, recorder := createValidationContext(request)

		middleware.Validation[UnknownValidationTagRequest]()(context)

		var response ErrorResponse
		json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
		assert.Len(t, response.Messages, 1)
		assert.Equal(t, "erro desconhecido", response.Messages[0])
		assert.True(t, context.IsAborted())
	})
}

func createValidationContext(request interface{}) (*gin.Context, *httptest.ResponseRecorder) {
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	requestInBytes, _ := json.Marshal(request)
	bodyBuffer := bytes.NewBuffer(requestInBytes)
	context.Request = &http.Request{
		Body: io.NopCloser(bodyBuffer),
	}

	return context, recorder
}

func createValidationContextWithStringRequest(request string) (*gin.Context, *httptest.ResponseRecorder) {
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	requestInBytes := []byte(request)
	bodyBuffer := bytes.NewBuffer(requestInBytes)
	context.Request = &http.Request{
		Body: io.NopCloser(bodyBuffer),
	}

	return context, recorder
}

func createCorrectRequest(fieldA string, fieldB string) CorrectRequest {
	return CorrectRequest{&fieldA, &fieldB}
}

func createWrongTypeRequest(fieldA int, fieldB int) WrongTypeRequest {
	return WrongTypeRequest{&fieldA, &fieldB}
}

func createMissingRequiredFieldRequest(fieldB string) MissingRequiredFieldRequest {
	return MissingRequiredFieldRequest{&fieldB}
}

func createUnknownValidationTagRequest(fieldA string) UnknownValidationTagRequest {
	return UnknownValidationTagRequest{&fieldA}
}

func createRequestWithSyntaxError() string {
	return `{
    "field_a": "Field A",,
    "field_b": "Field B"
	}`
}
