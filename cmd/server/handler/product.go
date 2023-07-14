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
	ProductCode    *string  `json:"product_code" binding:"required,gt=3"`
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

func NewProduct(service product.Service) *Product {
	return &Product{service}
}

// Get All products godoc
// @Summary List all products
// @Description Returns a collection of existing products.
// @Tags Products
// @Produce json
// @Success 200 {object} []domain.Product "List of all products"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /products [get]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products := p.service.GetAll()
		web.Success(c, http.StatusOK, products)
	}
}

// Get godoc
// @Summary Get a product by id
// @Description Get a product based on the provided id. Returns a not found error if the warehouse does not exist.
// @Tags Products
// @Produce json
// @Param id path int true "Product Id"
// @Success 200 {object} []domain.Product "Created product"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /products/{id} [get]
func (p *Product) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("Id")

		product, err := p.service.Get(id)

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
// @Summary Create a product
// @Description Create a new product based on the provided JSON payload.
// @Tags Products
// @Accept json
// @Produce json
// @Param request body CreateProductRequest true "Product to be created"
// @Success 201 {object} domain.Product "Created product"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 422 {object} web.ErrorResponse "Validation error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /products [post]
func (p *Product) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := c.MustGet(RequestParamContext).(CreateProductRequest)

		created, err := p.service.Create(request.ToProduct())

		if err != nil {
			if apperr.Is[*apperr.ResourceAlreadyExists](err) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
			if apperr.Is[*apperr.DependentResourceNotFound](err) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
		}

		web.Success(c, http.StatusCreated, created)
	}
}

// Update godoc
// @Summary Update a product
// @Description Update an existent product based on the provided id and JSON payload.
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product id"
// @Param request body UpdateProductRequest true "Product data"
// @Success 200 {object} domain.Product "Updated product"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 422 {object} web.ErrorResponse "Validation error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /products/{id} [patch]
func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("Id")
		request := c.MustGet(RequestParamContext).(UpdateProductRequest)

		response, err := p.service.Update(id, request.ToUpdateProduct())

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}

			if apperr.Is[*apperr.ResourceAlreadyExists](err) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}

			if apperr.Is[*apperr.DependentResourceNotFound](err) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
		}

		web.Success(c, http.StatusOK, response)
	}
}

// Delete godoc
// @Summary Delete a product
// @Description Delete a product based on the provided id.
// @Tags Products
// @Param id path int true "Product ID"
// @Success 204 "No content"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /products/{id} [delete]
func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.GetInt("Id")

		err := p.service.Delete(id)

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(c, http.StatusNoContent, nil)
	}
}

// Create godoc
// @Summary Count records by products
// @Description Record count by product.
// @Description If no query param is given, it brings the report to all product records.
// @Description If a product id is specified, it brings the number of records for this product.
// @Tags Products
// @Produce json
// @Param id query int false "Product ID"
// @Success 200 {object} []domain.RecordsByProductReport "Report of records by product"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /products/report-records [get]
func (p *Product) ReportRecords() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Request.URL.Query().Get("id")

		if idParam == "" {
			result := p.service.CountRecordsByAllProducts()
			web.Success(c, http.StatusOK, result)
			return
		}

		id, err := strconv.Atoi(idParam)

		if err != nil {
			web.Error(c, http.StatusBadRequest, InvalidId, idParam)
			return
		}

		productRecords, err := p.service.CountRecordsByProduct(id)

		if err != nil {
			if apperr.Is[*apperr.ResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(c, http.StatusOK, productRecords)
	}
}
