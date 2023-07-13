package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/inbound_order"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type InboundOrder struct {
	service inbound_order.Service
}

type CreateInboundOrderRequest struct {
	OrderDate      *string `json:"order_date" binding:"required,datetime=2006-01-02 15:04:05"`
	OrderNumber    *string `json:"order_number" binding:"required,gt=3"`
	EmployeeId     *int    `json:"employee_id" binding:"required"`
	ProductBatchId *int    `json:"product_batch_id" binding:"required"`
	WarehouseId    *int    `json:"warehouse_id" binding:"required"`
}

func (i CreateInboundOrderRequest) ToInboundOrder() domain.InboundOrder {
	return domain.InboundOrder{
		ID:             0,
		OrderDate:      helpers.ToDateTime(*i.OrderDate),
		OrderNumber:    *i.OrderNumber,
		EmployeeId:     *i.EmployeeId,
		ProductBatchId: *i.ProductBatchId,
		WarehouseId:    *i.WarehouseId,
	}
}

func NewInboundOrder(service inbound_order.Service) *InboundOrder {
	return &InboundOrder{service}
}

// Create godoc
// @Summary Create a new inbound_order
// @Description Create a new inbound_order based on the provided JSON payload
// @Tags Inbound Orders
// @Accept json
// @Produce json
// @Param request body CreateInboundOrderRequest true "InboundOrder data"
// @Success 201 {object} domain.InboundOrder "Created inbound_order"
// @Failure 422 {object} web.ErrorResponse "Validation error"
// @Failure 409 {object} web.ErrorResponse "Conflict error"
// @Failure 500 {object} web.ErrorResponse "Internal server error"
// @Router /inbound-orders [post]
func (i *InboundOrder) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		request := c.MustGet(RequestParamContext).(CreateInboundOrderRequest)

		created, err := i.service.Create(request.ToInboundOrder())

		if err != nil {
			if apperr.Is[*apperr.DependentResourceNotFound](err) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
			if apperr.Is[*apperr.ResourceAlreadyExists](err) {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
		}

		web.Success(c, http.StatusCreated, created)

	}
}
