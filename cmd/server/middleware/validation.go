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

const (
	CannotBeBlank = "pelo menos um dos seguintes campos deve ser informado para modificações: %v"
)

func RequestValidation[T any](canBeBlank bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request T
		if err := ctx.ShouldBindJSON(&request); err != nil {
			status := http.StatusUnprocessableEntity
			errorMessages := make([]string, 0)

			if syntaxError, ok := err.(*json.SyntaxError); ok {
				errorMessages = append(errorMessages, fmt.Sprintf("erro de sintaxe na posição %d: %v", syntaxError.Offset, syntaxError.Error()))
			}

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
			return
		}

		if !canBeBlank {
			if isBlank(request) {
				web.Error(ctx, http.StatusBadRequest, CannotBeBlank, strings.Join(getFieldNames(request), ", "))
				ctx.Abort()
				return
			}
		}

		ctx.Set("Request", request)
	}
}

func readableMessageFrom(structValue interface{}, fe validator.FieldError) string {
	var message string
	switch fe.Tag() {
	case "required":
		message = fmt.Sprintf("'%s' é obrigatório", getFieldNameOfFieldError(structValue, fe))
	case "e164":
		message = fmt.Sprintf("'%s' precisa estar no formato +<country_code><zone_code><phone_number> sem espaços ou caracteres especiais, por exemplo: +5500123456789", getFieldNameOfFieldError(structValue, fe))
	default:
		message = "erro desconhecido"
	}

	return strings.ToLower(message)
}

func getFieldNameOfFieldError(structValue interface{}, err validator.FieldError) string {
	structType := reflect.TypeOf(structValue)
	fieldName := strings.SplitN(err.Namespace(), ".", 2)[1]
	field, _ := structType.FieldByName(fieldName)

	jsonTag := field.Tag.Get("json")
	jsonTag = strings.Split(jsonTag, ",")[0]
	return jsonTag
}

func getFieldNames(structValue interface{}) []string {
	structType := reflect.TypeOf(structValue)
	fieldNames := make([]string, 0)

	for i := 0; i < structType.NumField(); i++ {
		jsonTag := structType.Field(i).Tag.Get("json")
		jsonTag = strings.Split(jsonTag, ",")[0]
		fieldNames = append(fieldNames, jsonTag)
	}

	return fieldNames
}

func isBlank(request any) bool {
	valueOfRequest := reflect.ValueOf(request)

	for index := 0; index < valueOfRequest.NumField(); index++ {
		valueField := valueOfRequest.Field(index)

		if !valueField.IsNil() {
			return false
		}
	}

	return true
}
