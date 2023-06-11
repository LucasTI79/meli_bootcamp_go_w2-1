package handler

import (
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type Product struct {
	service product.Service
}

func NewProduct(service product.Service) *Product {
	return &Product{service}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, _ := p.service.GetAll(c)
		web.Success(c, http.StatusOK, products)
	}
}

func (p *Product) Get() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (p *Product) Create() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
