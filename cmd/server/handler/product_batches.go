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
	BatchNumber        *int     `json:"batch_number" binding:"required"`
	CurrentQuantity    *int     `json:"current_quantity" binding:"required"`
	CurrentTemperature *float32 `json:"current_temperature" binding:"required"`
	DueDate            *string  `json:"due_date" binding:"required"`
	InitialQuantity    *int     `json:"initial_quantity" binding:"required"`
	ManufacturingDate  *string  `json:"manufacturing_date" binding:"required"`
	ManufacturingHour  *int     `json:"manufacturing_hour" binding:"required"`
	MinimumTemperature *float32 `json:"minimum_temperature" binding:"required"`
	ProductID          *int     `json:"product_id" binding:"required"`
	SectionID          *int     `json:"section_id" binding:"required"`
}

func (s *CreateProductBatchesRequest) ToProductBatches() domain.ProductBatches {

	return domain.ProductBatches{
		BatchNumber:        *s.BatchNumber,
		CurrentQuantity:    *s.CurrentQuantity,
		CurrentTemperature: *s.CurrentTemperature,
		DueDate:            *s.DueDate,
		InitialQuantity:    *s.InitialQuantity,
		ManufacturingDate:  *s.ManufacturingDate,
		ManufacturingHour:  *s.ManufacturingHour,
		MinimumTemperature: *s.MinimumTemperature,
		ProductID:          *s.ProductID,
		SectionID:          *s.SectionID,
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

		created, err := pb.productBatchService.Create(request.ToProductBatches())

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
