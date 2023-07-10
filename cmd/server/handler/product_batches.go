package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_batches"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductBatches struct {
	productBatchService product_batches.Service
}
type CreateProductBatchesRequest struct {
	ID                 int     `json:"id"`
	BatchNumber        int     `json:"batch_number"`
	CurrentQuantity    int     `json:"current_quantity"`
	CurrentTemperature float64 `json:"current_temperature"`
	DueDate            string  `json:"due_date"`
	InitialQuantity    int     `json:"initial_quantity"`
	ManufacturingDate  string  `json:"manufacturing_date"`
	ManufacturingHour  string  `json:"manufacturing_hour"`
	MinimumTemperature int     `json:"minimum_temperature"`
	ProductID          int     `json:"product_id"`
	SectionID          int     `json:"section_id"`
}

func (s *CreateProductBatchesRequest) ToProductBatches() domain.ProductBatches {

	return domain.ProductBatches{
		BatchNumber:        s.BatchNumber,
		CurrentQuantity:    s.CurrentQuantity,
		CurrentTemperature: s.CurrentTemperature,
		DueDate:            s.DueDate,
		InitialQuantity:    s.InitialQuantity,
		ManufacturingDate:  s.ManufacturingDate,
		ManufacturingHour:  s.ManufacturingHour,
		MinimumTemperature: s.MinimumTemperature,
		ProductID:          s.ProductID,
		SectionID:          s.SectionID,
	}
}

func NewProductBatches(service product_batches.Service) *ProductBatches {
	return &ProductBatches{
		productBatchService: service,
	}
}

// Create godoc
// @Summary Create/ Save a new product batches
// @Description Create a new product batches based on the provided JSON payload
// @Tags ProductBatches
// @Accept json
// @Produce json
// @Param request body CreateProductBatchesRequest true "Product Batches data"
// @Success 201 {object} domain.ProductBatches "Created product batches"
// @Failure 400 {object} web.ErrorResponse "Bad request"
// @Failure 404 {object} web.ErrorResponse "Not found"
// @Failure 422 {object} web.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /product-batches [post]
func (pb *ProductBatches) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := c.MustGet(RequestParamContext).(CreateProductBatchesRequest)

		exists, err := pb.productBatchService.Exists(request.BatchNumber)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		if exists {
			web.Error(c, http.StatusConflict, "product batches already exists")
			return
		}

		created, err := pb.productBatchService.Create(request.ToProductBatches())

		if err != nil {
			if apperr.Is[*apperr.ResourceAlreadyExists](err) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
			if apperr.Is[*apperr.DependentResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
				return
			}
		}

		web.Success(c, http.StatusCreated, created)
	}
}

// Create godoc
// @Summary Get all product batches
// @Description product batches
// @Description If no query param is given, bring the report to all product batches.
// @Tags ProductBatches
// @Accept json
// @Produce json
// @Success 200 {object} []domain.ProductBatches "List of product batches"
// @Failure 400 {object} web.ErrorResponse "Validation error"
// @Failure 404 {object} web.ErrorResponse "Resource not found error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /api/v1/sections/report-product [get]
func (pb *ProductBatches) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := pb.productBatchService.Get()
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		if len(id) == 0 {
			web.Error(c, http.StatusNotFound, "product batches not found")
			return
		}
		web.Success(c, http.StatusOK, id)
	}
}
