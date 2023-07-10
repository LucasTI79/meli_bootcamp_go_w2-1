package handler

import (
	"net/http"
	"time"

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

func DataConvert(pb CreateProductBatchesRequest) (product_batches.ProductBatches, error) {
	Duedate, err := time.Parse("2023-07-07", pb.DueDate)
	if err != nil {
		return product_batches.ProductBatches{}, err
	}
	ManufacturingDate, err := time.Parse("2023-07-07", pb.ManufacturingDate)
	if err != nil {
		return product_batches.ProductBatches{}, err
	}
	return product_batches.ProductBatches{
		BatchNumber:        pb.BatchNumber,
		CurrentQuantity:    pb.CurrentQuantity,
		CurrentTemperature: pb.CurrentTemperature,
		DueDate:            Duedate,
		InitialQuantity:    pb.InitialQuantity,
		ManufacturingDate:  ManufacturingDate,
		ManufacturingHour:  pb.ManufacturingHour,
		MinimumTemperature: pb.MinimumTemperature,
		ProductID:          pb.ProductID,
		SectionID:          pb.SectionID,
	}, nil
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

func (pb *ProductBatches) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		productBatches, err := pb.productBatchService.Get(id)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		web.Success(c, http.StatusOK, productBatches)
	}
}
