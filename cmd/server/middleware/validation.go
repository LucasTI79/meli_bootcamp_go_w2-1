package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validation[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request T
		if err := ctx.ShouldBindJSON(&request); err != nil {
			status := http.StatusUnprocessableEntity
			errorMessages := make([]string, 0)

			if marshallingError, ok := err.(*json.UnmarshalTypeError); ok {
				errorMessages = append(errorMessages, fmt.Sprintf("o campo '%s' deve ser '%s'", marshallingError.Field, marshallingError.Type.String()))
			}

			if validationErrors, ok := err.(validator.ValidationErrors); ok {
				for _, err := range validationErrors {
					errorMessages = append(errorMessages, readableMessageFrom(request, err))
				}
			}

			response := web.ErrorResponse{
				Code:     strings.ReplaceAll(strings.ToLower(http.StatusText(status)), " ", "_"),
				Messages: errorMessages,
				Status:   status,
			}

			web.Response(ctx, status, response)
			ctx.Abort()
		}

		ctx.Set("Request", request)
	}
}

func readableMessageFrom(structValue interface{}, fe validator.FieldError) string {
	var message string
	switch fe.Tag() {
	case "required":
		message = fmt.Sprintf("'%s' é obrigatório", getFieldName(structValue, fe))
	case "e164":
		message = fmt.Sprintf("'%s' precisa estar no formato +<country_code><zone_code><phone_number> sem espaços ou caracteres especiais, por exemplo: +5500123456789", getFieldName(structValue, fe))
	default:
		message = "erro desconhecido"
	}

	return strings.ToLower(message)
}

func getFieldName(structValue interface{}, err validator.FieldError) string {
	structType := reflect.TypeOf(structValue)
	fieldName := strings.SplitN(err.Namespace(), ".", 2)[1]
	field, _ := structType.FieldByName(fieldName)

	jsonTag := field.Tag.Get("json")
	jsonTag = strings.Split(jsonTag, ",")[0]
	return jsonTag
}
