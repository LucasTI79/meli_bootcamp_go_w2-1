package handler

import (
	"net/http"
	"strconv"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

const (
	RequestParamContext = "Request"
	InvalidId           = "o id '%s' é inválido"
	CannotBeBlank       = "pelo menos um campo deve ser informado para modificações"
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

func (r UpdateProductRequest) ToUpdateProduct() domain.UpdateProduct {
	return domain.UpdateProduct{
		Description:    r.Description,
		ExpirationRate: r.ExpirationRate,
		FreezingRate:   r.FreezingRate,
		Height:         r.Height,
		Length:         r.Length,
		Netweight:      r.Netweight,
		ProductCode:    r.ProductCode,
		RecomFreezTemp: r.RecomFreezTemp,
		Width:          r.Width,
		ProductTypeID:  r.ProductTypeID,
		SellerID:       r.SellerID,
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

// Create godoc
// @Summary List products
// @Description List all products
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Product "List of products"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /products [get]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products := p.service.GetAll(c.Request.Context())
		web.Success(c, http.StatusOK, products)
	}
}

// Get godoc
// @Summary Get a product by id
// @Description Get a product based on the provided id
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product Id"
// @Success 200 {object} []domain.Product "Created product"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /products/{id} [get]
func (p *Product) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Param("id")
		id, err := strconv.Atoi(requestId)

		if err != nil {
			web.Error(c, http.StatusBadRequest, InvalidId, requestId)
			return
		}

		product, err := p.service.Get(c.Request.Context(), id)

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(c, http.StatusOK, product)
	}
}

// Create godoc
// @Summary Create a new product
// @Description Create a new product based on the provided JSON payload
// @Tags Products
// @Accept json
// @Produce json
// @Param request body CreateProductRequest true "Product data"
// @Success 201 {object} domain.Product "Created product"
// @Failure 422 {object} web.ErrorResponse "Validation error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /products [post]
func (p *Product) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := c.MustGet(RequestParamContext).(CreateProductRequest)

		created, err := p.service.Create(c.Request.Context(), request.ToProduct())

		if err != nil {
			if apperr.Is[*apperr.ResourceAlreadyExists](err) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
		}

		web.Success(c, http.StatusCreated, created)
	}
}

// Update godoc
// @Summary Update a product
// @Description Update an existent product based on the provided JSON payload
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product id"
// @Param request body UpdateProductRequest true "Product data"
// @Success 200 {object} domain.Product "Updated product"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 422 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /products/{id} [patch]
func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Param("id")
		id, err := strconv.Atoi(requestId)

		if err != nil {
			web.Error(c, http.StatusBadRequest, InvalidId, requestId)
			return
		}

		request := c.MustGet(RequestParamContext).(UpdateProductRequest)

		if request.IsBlank() {
			web.Error(c, http.StatusBadRequest, CannotBeBlank)
			return
		}

		response, err := p.service.Update(c.Request.Context(), id, request.ToUpdateProduct())

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}

			if apperr.Is[*apperr.ResourceAlreadyExists](err) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
		}

		web.Success(c, http.StatusOK, response)
	}
}

// Delete godoc
// @Summary Delete a product
// @Description Delete a product based on the provided JSON payload
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product id"
// @Success 204
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /products/{id} [delete]
func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Param("id")
		id, err := strconv.Atoi(requestId)

		if err != nil {
			web.Error(c, http.StatusBadRequest, InvalidId, requestId)
			return
		}

		err = p.service.Delete(c.Request.Context(), id)

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(c, http.StatusNoContent, nil)
	}
}
