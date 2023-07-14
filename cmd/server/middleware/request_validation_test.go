package middleware_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/middleware"
	"github.com/stretchr/testify/assert"
)

type CorrectRequest struct {
	FieldA *string `json:"field_a" binding:"required"`
	FieldB *string `json:"field_b" binding:"e164,omitempty"`
	FieldC *string `json:"field_c" binding:"datetime=2006-01-02 15:04:05,omitempty"`
	FieldD *string `json:"field_d" binding:"gt=3,omitempty"`
}

type WrongTypeRequest struct {
	FieldA *int `json:"field_a" binding:"required"`
	FieldB *int `json:"field_b" binding:"e164,omitempty"`
	FieldD *int `json:"field_c" binding:"gt=3,omitempty"`
}

type MissingRequiredFieldRequest struct {
	FieldB *string `json:"field_b" validate:"e164,omitempty"`
	FieldC *string `json:"field_c" validate:"datetime=2006-01-02 15:04:05,omitempty"`
	FieldD *string `json:"field_d" validate:"gt=3,omitempty"`
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
	fieldC := "2023-01-01 00:00:00"
	fieldD := "Field D"

	t.Run("Should have success on validation", func(t *testing.T) {
		request := createCorrectRequest(fieldA, fieldB, fieldC, fieldD)
		context, recorder, _ := createValidationContext(request, getMarshaledRequestInBytes[CorrectRequest])

		middleware.RequestValidation[CorrectRequest](true)(context)
		gotRequest := context.MustGet("Request").(CorrectRequest)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, request, gotRequest)
		assert.False(t, context.IsAborted())
	})

	t.Run("Should have error when try parse a empty body request", func(t *testing.T) {
		request := ""
		context, recorder, _ := createValidationContext(request, getStringRequestInBytes)

		middleware.RequestValidation[CorrectRequest](true)(context)

		var response ErrorResponse
		_ = json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
		assert.Len(t, response.Messages, 1)
		assert.Equal(t, "o corpo da requisição está vazio e precisa ser um objeto JSON válido", response.Messages[0])
		assert.True(t, context.IsAborted())
	})

	t.Run("Should have error when try parse a request field with a syntax error", func(t *testing.T) {
		request := createRequestWithSyntaxError()
		context, recorder, _ := createValidationContext(request, getStringRequestInBytes)

		middleware.RequestValidation[CorrectRequest](true)(context)

		var response ErrorResponse
		_ = json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
		assert.Len(t, response.Messages, 1)
		assert.Contains(t, response.Messages[0], "erro de sintaxe na posição")
		assert.True(t, context.IsAborted())
	})

	t.Run("Should have error when try parse a request field with a wrong type", func(t *testing.T) {
		request := createWrongTypeRequest(1, 1, 1)
		context, recorder, _ := createValidationContext(request, getMarshaledRequestInBytes[WrongTypeRequest])

		middleware.RequestValidation[CorrectRequest](true)(context)

		var response ErrorResponse
		_ = json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
		assert.Len(t, response.Messages, 1)
		assert.Equal(t, "o campo 'field_a' deve ser 'string'", response.Messages[0])
		assert.True(t, context.IsAborted())
	})

	t.Run("Should have error when try parse a request with a missing required field", func(t *testing.T) {
		request := createMissingRequiredFieldRequest(fieldB, fieldC, fieldD)
		context, recorder, _ := createValidationContext(request, getMarshaledRequestInBytes[MissingRequiredFieldRequest])

		middleware.RequestValidation[CorrectRequest](true)(context)

		var response ErrorResponse
		_ = json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
		assert.Len(t, response.Messages, 1)
		assert.Equal(t, "'field_a' é obrigatório", response.Messages[0])
		assert.True(t, context.IsAborted())
	})

	t.Run("Should have error when try parse a request with a wrong phone format", func(t *testing.T) {
		request := createCorrectRequest(fieldA, "Phone", fieldC, fieldD)
		context, recorder, _ := createValidationContext(request, getMarshaledRequestInBytes[CorrectRequest])

		middleware.RequestValidation[CorrectRequest](true)(context)

		var response ErrorResponse
		_ = json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
		assert.Len(t, response.Messages, 1)
		assert.Equal(t, "'field_b' precisa estar no formato +<country_code><zone_code><phone_number> sem espaços ou caracteres especiais, por exemplo: +5500123456789", response.Messages[0])
		assert.True(t, context.IsAborted())
	})

	t.Run("Should have error when try parse a request with a wrong datetime format", func(t *testing.T) {
		request := createCorrectRequest(fieldA, fieldB, "Date", fieldD)
		context, recorder, _ := createValidationContext(request, getMarshaledRequestInBytes[CorrectRequest])

		middleware.RequestValidation[CorrectRequest](true)(context)

		var response ErrorResponse
		_ = json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
		assert.Len(t, response.Messages, 1)
		assert.Equal(t, "'field_c' precisa estar no formato yyyy-mm-dd hh:mm:ss", response.Messages[0])
		assert.True(t, context.IsAborted())
	})

	t.Run("Should have error when try parse a request field with less or equal 3 characters", func(t *testing.T) {
		request := createCorrectRequest(fieldA, fieldB, fieldC, "a")
		context, recorder, _ := createValidationContext(request, getMarshaledRequestInBytes[CorrectRequest])

		middleware.RequestValidation[CorrectRequest](true)(context)

		var response ErrorResponse
		_ = json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
		assert.Len(t, response.Messages, 1)
		assert.Equal(t, "'field_d' precisa ter mais de 3 caracteres", response.Messages[0])
		assert.True(t, context.IsAborted())
	})

	t.Run("Should have error when try parse a request with a unknown validation tag", func(t *testing.T) {
		request := createUnknownValidationTagRequest(fieldA)
		context, recorder, _ := createValidationContext(request, getMarshaledRequestInBytes[UnknownValidationTagRequest])

		middleware.RequestValidation[UnknownValidationTagRequest](true)(context)

		var response ErrorResponse
		_ = json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
		assert.Len(t, response.Messages, 1)
		assert.Equal(t, "erro desconhecido", response.Messages[0])
		assert.True(t, context.IsAborted())
	})

	t.Run("Should have error when request that 'cannot be blank' is blank", func(t *testing.T) {
		request := createEmptyRequest()
		context, recorder, _ := createValidationContext(request, getStringRequestInBytes)

		middleware.RequestValidation[MissingRequiredFieldRequest](false)(context)

		var response ErrorResponse
		_ = json.Unmarshal(recorder.Body.Bytes(), &response)

		assert.Equal(t, http.StatusBadRequest, recorder.Code)
		assert.Len(t, response.Messages, 1)
		assert.Contains(t, response.Messages[0], "pelo menos um dos seguintes campos deve ser informado para modificações:")
		assert.True(t, context.IsAborted())
	})

	t.Run("Should have success when request that 'cannot be blank' is not blank", func(t *testing.T) {
		request := createMissingRequiredFieldRequest(fieldB, fieldC, fieldD)
		context, recorder, _ := createValidationContext(request, getMarshaledRequestInBytes[MissingRequiredFieldRequest])

		middleware.RequestValidation[MissingRequiredFieldRequest](false)(context)
		gotRequest := context.MustGet("Request").(MissingRequiredFieldRequest)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, request, gotRequest)
		assert.False(t, context.IsAborted())
	})
}

func createCorrectRequest(fieldA string, fieldB string, fieldC string, fieldD string) CorrectRequest {
	return CorrectRequest{&fieldA, &fieldB, &fieldC, &fieldD}
}

func createWrongTypeRequest(fieldA int, fieldB int, fieldD int) WrongTypeRequest {
	return WrongTypeRequest{&fieldA, &fieldB, &fieldD}
}

func createMissingRequiredFieldRequest(fieldB string, fieldC string, fieldD string) MissingRequiredFieldRequest {
	return MissingRequiredFieldRequest{&fieldB, &fieldC, &fieldD}
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

func createEmptyRequest() string {
	return `{
	}`
}
