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

type UpdateProductRequest struct {
	Description    *string  `json:"description"`
	ExpirationRate *float32 `json:"expiration_rate"`
	FreezingRate   *float32 `json:"freezing_rate"`
	Height         *float32 `json:"height"`
	Length         *float32 `json:"length"`
	Netweight      *float32 `json:"netweight"`
	ProductCode    *string  `json:"product_code"`
	RecomFreezTemp *float32 `json:"recommended_freezing_temperature"`
	Width          *float32 `json:"width"`
	ProductTypeID  *int     `json:"product_type_id"`
	SellerID       *int     `json:"seller_id"`
}

func (r UpdateProductRequest) ToProductOptional() domain.ProductOptional {
	return domain.ProductOptional{
		Description:    domain.Optional{Value: r.Description, HasValue: r.Description != nil},
		ExpirationRate: domain.Optional{Value: r.ExpirationRate, HasValue: r.ExpirationRate != nil},
		FreezingRate:   domain.Optional{Value: r.FreezingRate, HasValue: r.FreezingRate != nil},
		Height:         domain.Optional{Value: r.Height, HasValue: r.Height != nil},
		Length:         domain.Optional{Value: r.Length, HasValue: r.Length != nil},
		Netweight:      domain.Optional{Value: r.Netweight, HasValue: r.Netweight != nil},
		ProductCode:    domain.Optional{Value: r.ProductCode, HasValue: r.ProductCode != nil},
		RecomFreezTemp: domain.Optional{Value: r.RecomFreezTemp, HasValue: r.RecomFreezTemp != nil},
		Width:          domain.Optional{Value: r.Width, HasValue: r.Width != nil},
		ProductTypeID:  domain.Optional{Value: r.ProductTypeID, HasValue: r.ProductTypeID != nil},
		SellerID:       domain.Optional{Value: r.SellerID, HasValue: r.SellerID != nil},
	}
}

func (productRequest UpdateProductRequest) IsBlank() bool {
	return productRequest.Description == nil &&
		productRequest.ExpirationRate == nil &&
		productRequest.FreezingRate == nil &&
		productRequest.Height == nil &&
		productRequest.Length == nil &&
		productRequest.Netweight == nil &&
		productRequest.ProductCode == nil &&
		productRequest.RecomFreezTemp == nil &&
		productRequest.Width == nil &&
		productRequest.ProductTypeID == nil &&
		productRequest.SellerID == nil
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
	return func(c *gin.Context) {
		requestId := c.Param("id")
		id, err := strconv.Atoi(requestId)

		if err != nil {
			web.Error(c, http.StatusBadRequest, "id '%s' is not valid", requestId)
			return
		}

		product, err := p.service.Get(c, id)

		if err != nil {
			if _, ok := err.(*apperr.ResourceNotFound); ok {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}

			web.Error(c, http.StatusInternalServerError, apperr.NewInternalServerError().Error())
			return
		}

		web.Success(c, http.StatusOK, product)
	}
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
	return func(c *gin.Context) {
		requestId := c.Param("id")
		id, err := strconv.Atoi(requestId)

		if err != nil {
			web.Error(c, http.StatusBadRequest, "id '%s' is not valid", requestId)
			return
		}

		var request UpdateProductRequest
		if err := c.ShouldBindJSON(&request); err != nil {
			web.ValidationError(c, err)
			return
		}

		if request.IsBlank() {
			web.Error(c, http.StatusBadRequest, "at least one field must be informed for modifications")
			return
		}

		response, err := p.service.Update(c, id, request.ToProductOptional())

		if err != nil {
			if _, ok := err.(*apperr.ResourceNotFound); ok {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
			if _, ok := err.(*apperr.ResourceAlreadyExists); ok {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
			web.Error(c, http.StatusInternalServerError, apperr.NewInternalServerError().Error())
			return
		}

		web.Success(c, http.StatusOK, response)
	}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Param("id")
		id, err := strconv.Atoi(requestId)

		if err != nil {
			web.Error(c, http.StatusBadRequest, "id '%s' is not valid", requestId)
			return
		}

		err = p.service.Delete(c, id)

		if err != nil {
			if _, ok := err.(*apperr.ResourceNotFound); ok {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}

			web.Error(c, http.StatusInternalServerError, apperr.NewInternalServerError().Error())
			return
		}

		web.Success(c, http.StatusNoContent, nil)
	}
}
