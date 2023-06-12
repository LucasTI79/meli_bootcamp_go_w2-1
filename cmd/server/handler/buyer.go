package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer"
)

type Buyer struct {
	buyerService buyer.IService
}

func NewBuyer(b buyer.IService) *Buyer {
	return &Buyer{
		buyerService: b,
	}
}

func (b *Buyer) Get() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (b *Buyer) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		buyers, err := b.buyerService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		if len(buyers) == 0 {
			web.Success(c, http.StatusNoContent, buyers)
			return
		}

		web.Success(c, http.StatusOK, buyers)

	}
}

func (b *Buyer) Create() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (b *Buyer) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (b *Buyer) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
