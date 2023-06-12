package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type Seller struct {
	sellerService seller.Service
	// sellerService seller.Service
}

func NewSeller(s seller.Service) *Seller {
	return &Seller{
		sellerService: s,
		// sellerService: s,
	}

}

func (s *Seller) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		sellers, err := s.sellerService.GetAll(c.Request.Context())
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "internal server error")
			return
		}
		if len(sellers) == 0 {
			web.Response(c, 204, sellers)
			return
		}
		web.Success(c, http.StatusOK, sellers)
	}
}

func (s *Seller) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := c.Params.Get("id")
		parsedId, err := strconv.Atoi(id)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "id received is invalid")
			return
		}
		foundSeller, err := s.sellerService.Get(c, parsedId)
		if err != nil {
			if errors.Is(err, seller.ErrNotFound) {
				web.Error(c, http.StatusNotFound, "could not find id %v", parsedId)
				return
			} else {
				web.Error(c, http.StatusInternalServerError, "internal server error")
				return
			}
		}
		web.Response(c, 200, foundSeller)
	}
}

func (s *Seller) Create() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (s *Seller) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (s *Seller) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
