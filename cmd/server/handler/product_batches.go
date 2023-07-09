package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_batches"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductBatches struct {
	productBatchService product_batches.Service
	sectionService      section.Service
	productService      product.Service
}

type CountProductBatchesBySection struct {
	SectionID     int `json:"section_id"`
	SectionNumber int `json:"section_number"`
	ProductsCount int `json:"products_count"`
}

func NewProductBatches(pb product_batches.Service, sectionService section.Service, productService product.Service) *ProductBatches {
	return &ProductBatches{
		productBatchService: pb,
		sectionService:      sectionService,
		productService:      productService,
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
func (s *ProductBatches) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var request domain.ProductBatches
		if err := c.ShouldBindJSON(&request); err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		if err := request.Validate(); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}
		err := s.productService.Exists(request.ProductID)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}
		err = s.sectionService.ExistsSectionID(request.SectionID)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}
		productBatchesID, err := s.productBatchService.Save(request)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, err.Error())
			return
		}
		request.ID = productBatchesID
		web.Success(c, http.StatusCreated, request)
	}
}
