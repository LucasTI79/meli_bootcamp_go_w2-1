package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_batches"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

const (
	InvalidId = "o id '%s' é inválido"
)

type ProductBatches struct {
	service product_batches.Service
}

type CreateProductBatchesRequest struct {
	ProductID int `json:"product_id" binding:"required"`
	SectionId int `json:"section_id" binding:"required"`
}

func NewProductBatches(service product_batches.Service) *ProductBatches {
	return &ProductBatches{service}
}

func (pb *ProductBatches) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newProductBatch domain.ProductBatches

		if err := c.ShouldBindJSON(&newProductBatch); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		if newProductBatch.BatchNumber == "" || newProductBatch.SectionID == 0 {
			web.Error(c, http.StatusUnprocessableEntity, "Required fields missing")
			return
		}

		for _, existingProductBatch := range productBatches {
			if existingProductBatch.BatchNumber == newProductBatch.BatchNumber {
				web.Error(c, http.StatusConflict, "batch_number already exists")
				return
			}
		}

		if !CheckSectionExists(newProductBatch.SectionID) {
			web.Error(c, http.StatusNotFound, "section_id not found")
			return
		}

		if !CheckProductExists(newProductBatch.ProductID) {
			web.Error(c, http.StatusNotFound, "product_id not found")
			return
		}

		createdProductBatch := CreateNewProductBatch(newProductBatch)

		if createdProductBatch == nil {
			web.Error(c, http.StatusInternalServerError, "Error creating product batch")
			return
		}

		c.JSON(http.StatusCreated, createdProductBatch)
	}
}

func CreateNewProductBatch(newProductBatch domain.ProductBatches) {
	panic("unimplemented")
}
