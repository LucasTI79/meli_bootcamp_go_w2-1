package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type response struct {
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Status   int      `json:"-"`
	Code     string   `json:"code"`
	Messages []string `json:"message"`
}

func Response(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func Success(c *gin.Context, status int, data interface{}) {
	Response(c, status, response{Data: data})
}

// NewErrorf creates a new error with the given status code and the message
// formatted according to args and format.
func Error(c *gin.Context, status int, format string, args ...interface{}) {
	err := errorResponse{
		Code: strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Messages: []string{
			fmt.Sprintf(format, args...),
		},
		Status: status,
	}

	Response(c, status, err)
}

func ValidationError(c *gin.Context, err error) {
	status := http.StatusUnprocessableEntity
	errorMessages := make([]string, 0)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {

		for _, err := range validationErrors {
			errorMessages = append(errorMessages, readableMessageFrom(err))
		}
	} else if marshallingError, ok := err.(*json.UnmarshalTypeError); ok {
		errorMessages = append(errorMessages, fmt.Sprintf("the field '%s' must be a '%s'", marshallingError.Field, marshallingError.Type.String()))
	}

	response := errorResponse{
		Code:     strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
		Messages: errorMessages,
		Status:   status,
	}

	Response(c, status, response)
}

func readableMessageFrom(fe validator.FieldError) string {
	var message string
	switch fe.Tag() {
	case "required":
		message = fmt.Sprintf("'%s' is required", fe.Field())
	case "lte":
		message = fmt.Sprintf("'%s' should be less than %s", fe.Field(), fe.Param())
	case "gte":
		message = fmt.Sprintf("'%s' should be greater than %s", fe.Field(), fe.Param())
	default:
		message = "unknown error"
	}

	return strings.ToLower(message)
}
