package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
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
	return func(c *gin.Context) {

		var req domain.Request
		err := c.Bind(&req)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "Existem erros na formatacao do Json. Nao foi possível realizar o parse.")
			return	
		}

		if req.CardNumberID == "" {
			web.Error(c, http.StatusUnprocessableEntity, "O Card Number do comprador é obrigatório")
			return
		}
		if req.FirstName == "" {
			web.Error(c, http.StatusUnprocessableEntity, "O Nome do comprador é obrigatório")
			return
		}
		if req.LastName== "" {
			web.Error(c, http.StatusUnprocessableEntity, "O Sobrenome do comprador é obrigatório")
			return
		}

		buyerId, err := b.buyerService.Save(c, req)
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		buyer := domain.Buyer{}
		buyer.ID = buyerId
		buyer.CardNumberID = req.CardNumberID
		buyer.FirstName = req.FirstName
		buyer.LastName = req.LastName
		
		web.Success(c, http.StatusCreated, buyer)

	}
}

func (b *Buyer) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (b *Buyer) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
