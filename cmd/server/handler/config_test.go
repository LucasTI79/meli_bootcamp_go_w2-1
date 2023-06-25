package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	BaseUri               = "/api/v1"
	ResourceAlreadyExists = "resource already exists"
	ResourceNotFound      = "resource not found"
)

func CreateServer() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func MakeRequest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func CreateBody(obj interface{}) string {
	encoded, _ := json.Marshal(obj)
	return string(encoded)
}

func DefinePath(resourceUri string) string {
	return BaseUri + resourceUri
}

func DefinePathWithId(resourceUri string, id int) string {
	return DefinePath(resourceUri) + "/" + strconv.Itoa(id)
}

func ValidationMiddleware(requestObject interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("Request", requestObject)
	}
}
