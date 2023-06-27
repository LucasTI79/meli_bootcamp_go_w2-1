package middleware_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func createValidationContext[T any](request T, getRequestInBytes func(T) []byte) (*gin.Context, *httptest.ResponseRecorder, *gin.Engine) {
	recorder := httptest.NewRecorder()
	context, router := gin.CreateTestContext(recorder)
	requestInBytes := getRequestInBytes(request)
	bodyBuffer := bytes.NewBuffer(requestInBytes)
	context.Request = &http.Request{
		Body: io.NopCloser(bodyBuffer),
	}

	return context, recorder, router
}

func getMarshaledRequestInBytes[T any](request T) []byte {
	requestInBytes, _ := json.Marshal(request)
	return requestInBytes
}

func getStringRequestInBytes(request string) []byte {
	return []byte(request)
}
