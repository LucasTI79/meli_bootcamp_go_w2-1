package middleware

import (
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

const (
	InvalidId = "o id '%s' é inválido"
)

func IdValidation() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam := ctx.Param("id")
		if idParam != "" {
			id, err := strconv.Atoi(idParam)
			if err != nil {
				web.Error(ctx, http.StatusBadRequest, InvalidId, idParam)
				ctx.Abort()
				return
			}

			ctx.Set("Id", id)
		}
	}
}
