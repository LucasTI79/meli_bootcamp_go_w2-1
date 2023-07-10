package handler

import (
	"net/http"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/inbound_orders"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/web"
	"github.com/gin-gonic/gin"
)

type InboundOrder struct {
	service inbound_orders.Service
}

type CreateInboundOrderRequest struct {
	OrderDate string `json:"order_date" binding:"required,datetime=2006-01-02 15:04:05"`
	OrderNumber int `json:"order_number" binding:"required"`
	EmployeeId int `json:"employee_id" binding:"required"`
	ProductBatchId int `json:"product_batch_id" binding:"required"`
	WarehouseId int `json:"warehouse_id" binding:"required"`
}

func (i CreateInboundOrderRequest) ToInboundOrder() domain.InboundOrder {
	return domain.InboundOrder{
		ID: 0,
		OrderDate: helpers.ToDateTime(i.OrderDate),
		OrderNumber: i.OrderNumber,
		EmployeeId: i.EmployeeId,
		ProductBatchId: i.ProductBatchId,
		WarehouseId: i.WarehouseId,
	}
}

func NewInboundOrder(service inbound_orders.Service) *InboundOrder {
	return &InboundOrder{service}
}

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