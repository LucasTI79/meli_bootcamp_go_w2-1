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

type CreateProductRequest struct {
	Description    *string  `json:"description" binding:"required"`
	ExpirationRate *float32 `json:"expiration_rate" binding:"required"`
	FreezingRate   *float32 `json:"freezing_rate" binding:"required"`
	Height         *float32 `json:"height" binding:"required"`
	Length         *float32 `json:"length" binding:"required"`
	Netweight      *float32 `json:"netweight" binding:"required"`
	ProductCode    *string  `json:"product_code" binding:"required"`
	RecomFreezTemp *float32 `json:"recommended_freezing_temperature" binding:"required"`
	Width          *float32 `json:"width" binding:"required"`
	ProductTypeID  *int     `json:"product_type_id" binding:"required"`
	SellerID       *int     `json:"seller_id"`
}

func (r CreateProductRequest) ToProduct() domain.Product {
	return domain.Product{
		ID:             0,
		Description:    *r.Description,
		ExpirationRate: *r.ExpirationRate,
		FreezingRate:   *r.FreezingRate,
		Height:         *r.Height,
		Length:         *r.Length,
		Netweight:      *r.Netweight,
		ProductCode:    *r.ProductCode,
		RecomFreezTemp: *r.RecomFreezTemp,
		Width:          *r.Width,
		ProductTypeID:  *r.ProductTypeID,
		SellerID:       *r.SellerID,
	}
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
	return func(c *gin.Context) {
		var request CreateProductRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			web.ValidationError(c, err)
			return
		}

		created, err := p.service.Create(c, request.ToProduct())

		if err != nil {
			if _, ok := err.(*apperr.ResourceAlreadyExists); ok {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}

			web.Error(c, http.StatusInternalServerError, apperr.NewInternalServerError().Error())
			return
		}

		web.Success(c, http.StatusCreated, created)
	}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
