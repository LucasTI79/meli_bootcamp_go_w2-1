package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type response struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Status   int      `json:"-"`
	Code     string   `json:"code"`
	Messages []string `json:"messages"`
}

type Data struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func Response(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func Success(c *gin.Context, status int, data interface{}) {
	Response(c, status, response{Data: data})
}

func Error(c *gin.Context, status int, format string, args ...interface{}) {
	err := ErrorResponse{
		Code: strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Messages: []string{
			fmt.Sprintf(format, args...),
		},
		Status: status,
	}

	Response(c, status, err)
}
