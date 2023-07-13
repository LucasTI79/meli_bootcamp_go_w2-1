package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	record "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_record"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type ProductRecord struct {
	service record.Service
}

type CreateProductRecordRequest struct {
	LastUpdateDate *string  `json:"last_update_date" binding:"required,datetime=2006-01-02 15:04:05"`
	PurchasePrice  *float32 `json:"purchase_price" binding:"required"`
	SalePrice      *float32 `json:"sale_price" binding:"required"`
	ProductID      *int     `json:"product_id" binding:"required"`
}

func (s CreateProductRecordRequest) ToProductRecord() domain.ProductRecord {
	return domain.ProductRecord{
		ID:             0,
		LastUpdateDate: helpers.ToDateTime(*s.LastUpdateDate),
		PurchasePrice:  *s.PurchasePrice,
		SalePrice:      *s.SalePrice,
		ProductID:      *s.ProductID,
	}
}

func NewProductRecord(service record.Service) *ProductRecord {
	return &ProductRecord{service}
}

// Create godoc
// @Summary Create a product record
// @Description Create a new product record based on the provided JSON payload.
// @Tags Product Records
// @Accept json
// @Produce json
// @Param request body CreateProductRecordRequest true "Product Record to be created"
// @Success 201 {object} domain.ProductRecord "Created product record"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 422 {object} web.ErrorResponse "Validation error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /product-records [post]
func (pr *ProductRecord) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := c.MustGet(RequestParamContext).(CreateProductRecordRequest)

		created, err := pr.service.Create(request.ToProductRecord())

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
