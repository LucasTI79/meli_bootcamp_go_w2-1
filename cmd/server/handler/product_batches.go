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
	service product_batches.Service
}

type CreateProductBatchesRequest struct {
	BatchNumber        string  `json:"batch_number"`
	CurrentQuantity    int     `json:"current_quantity"`
	CurrentTemperature float64 `json:"current_temperature"`
	DueDate            string  `json:"due_date"`
	InitialQuantity    int     `json:"initial_quantity"`
	ManufacturingDate  string  `json:"manufacturing_date"`
	ManufacturingHour  string  `json:"manufacturing_date"`
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
	return &ProductBatches{service}
}

func (pb *ProductBatches) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := c.MustGet(RequestParamContext).(CreateProductBatchesRequest)

		created, err := pb.service.Create(request.ToProductBatches())

		if err != nil {
			if apperr.Is[*apperr.ResourceAlreadyExists](err) {
				web.Error(c, http.StatusConflict, err.Error())
			}
			if apperr.Is[*apperr.DependentResourceNotFound](err) {
				web.Error(c, http.StatusNotFound, err.Error())
			}
		}
		web.Success(c, http.StatusCreated, created)
	}
}

func (pb *ProductBatches) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		//id := c.Param("id")

		productBatches := pb.service.CountProductBatchesBySection()

		web.Success(c, http.StatusOK, productBatches)
	}
}
