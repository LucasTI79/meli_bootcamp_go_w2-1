package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/purchase_order"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type PurchaseOrder struct {
	service purchase_order.Service
}

type CreatePurchaseOrderRequest struct {
	OrderNumber     *string `json:"order_number" binding:"required,gt=3"`
	OrderDate       *string `json:"order_date" binding:"required,datetime=2006-01-02 15:04:05"`
	TrackingCode    *string `json:"tracking_code" binding:"required"`
	BuyerID         *int    `json:"buyer_id" binding:"required" `
	CarrierID       *int    `json:"carrier_id" binding:"required"`
	ProductRecordID *int    `json:"product_record_id" binding:"required"`
	OrderStatusID   *int    `json:"order_status_id" binding:"required"`
	WarehouseID     *int    `json:"warehouse_id" binding:"required"`
}

func (r CreatePurchaseOrderRequest) ToPurchaseOrder() domain.PurchaseOrder {

	return domain.PurchaseOrder{
		ID:              0,
		OrderNumber:     *r.OrderNumber,
		OrderDate:       helpers.ToDateTime(*r.OrderDate),
		TrackingCode:    *r.TrackingCode,
		BuyerID:         *r.BuyerID,
		CarrierID:       *r.CarrierID,
		ProductRecordID: *r.ProductRecordID,
		OrderStatusID:   *r.OrderStatusID,
		WarehouseID:     *r.WarehouseID,
	}
}

func NewPurchaseOrder(service purchase_order.Service) *PurchaseOrder {
	return &PurchaseOrder{service}
}

// Create godoc
// @Summary Create a new purchase order
// @Description Create a new purchase order based on the provided JSON payload
// @Tags Localities
// @Accept json
// @Produce json
// @Param request body CreatePurchaseOrderRequest true "Purchase order data"
// @Success 201 {object} domain.PurchaseOrder "Created purchase order"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 422 {object} web.ErrorResponse "Validation error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /purchase-orders [post]
func (po *PurchaseOrder) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := c.MustGet(RequestParamContext).(CreatePurchaseOrderRequest)

		created, err := po.service.Create(request.ToPurchaseOrder())

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
