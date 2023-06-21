package middleware

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

func InternalError() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				web.Error(ctx, http.StatusInternalServerError, "internal server error")
			}
		}()

		ctx.Next()
	}
}
