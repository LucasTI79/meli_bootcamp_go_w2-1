package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_batch"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductBatch struct {
	productBatchService product_batch.Service
}
type CreateProductBatchRequest struct {
	BatchNumber        *int     `json:"batch_number" binding:"required"`
	CurrentQuantity    *int     `json:"current_quantity" binding:"required"`
	CurrentTemperature *float32 `json:"current_temperature" binding:"required"`
	DueDate            *string  `json:"due_date" binding:"required,datetime=2006-01-02 15:04:05"`
	InitialQuantity    *int     `json:"initial_quantity" binding:"required"`
	ManufacturingDate  *string  `json:"manufacturing_date" binding:"required,datetime=2006-01-02 15:04:05"`
	ManufacturingHour  *int     `json:"manufacturing_hour" binding:"required"`
	MinimumTemperature *float32 `json:"minimum_temperature" binding:"required"`
	ProductID          *int     `json:"product_id" binding:"required"`
	SectionID          *int     `json:"section_id" binding:"required"`
}

func (pb *CreateProductBatchRequest) ToProductBatches() domain.ProductBatch {

	return domain.ProductBatch{
		BatchNumber:        *pb.BatchNumber,
		CurrentQuantity:    *pb.CurrentQuantity,
		CurrentTemperature: *pb.CurrentTemperature,
		DueDate:            helpers.ToDateTime(*pb.DueDate),
		InitialQuantity:    *pb.InitialQuantity,
		ManufacturingDate:  helpers.ToDateTime(*pb.ManufacturingDate),
		ManufacturingHour:  *pb.ManufacturingHour,
		MinimumTemperature: *pb.MinimumTemperature,
		ProductID:          *pb.ProductID,
		SectionID:          *pb.SectionID,
	}
}

func NewProductBatches(service product_batch.Service) *ProductBatch {
	return &ProductBatch{
		productBatchService: service,
	}
}

// Create godoc
// @Summary Create/ Save a new product batch
// @Description Create a new product batches based on the provided JSON payload
// @Tags Product Batch
// @Accept json
// @Produce json
// @Param request body CreateProductBatchRequest true "Product Batches data"
// @Success 201 {object} domain.ProductBatch "Created product batches"
// @Failure 400 {object} web.ErrorResponse "Bad request"
// @Failure 404 {object} web.ErrorResponse "Not found"
// @Failure 422 {object} web.ErrorResponse "Unprocessable Entity"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /product-batches [post]
func (pb *ProductBatch) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := c.MustGet(RequestParamContext).(CreateProductBatchRequest)

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
