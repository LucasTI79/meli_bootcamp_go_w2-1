package handler_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/middleware"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/purchase_order/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	ResourcePurchaseOrdersUri = "/purchase-orders"
)

var (
	mockedPurchaseOrder = domain.PurchaseOrder{
		ID:              1,
		OrderNumber:     "order#123",
		OrderDate:       time.Date(2023, 07, 10, 0, 0, 0, 0, time.UTC),
		TrackingCode:    "TRACK007",
		BuyerID:         1,
		CarrierID:       1,
		ProductRecordID: 1,
		OrderStatusID:   1,
		WarehouseID:     1,
	}
)

func TestCreatePurchaseOrder(t *testing.T) {
	orderDate := helpers.ToFormattedDateTime(mockedPurchaseOrder.OrderDate)

	requestObject := handler.CreatePurchaseOrderRequest{
		OrderNumber:     &mockedPurchaseOrder.OrderNumber,
		OrderDate:       &orderDate,
		TrackingCode:    &mockedPurchaseOrder.TrackingCode,
		BuyerID:         &mockedPurchaseOrder.BuyerID,
		CarrierID:       &mockedPurchaseOrder.CarrierID,
		ProductRecordID: &mockedPurchaseOrder.ProductRecordID,
		OrderStatusID:   &mockedPurchaseOrder.OrderStatusID,
		WarehouseID:     &mockedPurchaseOrder.WarehouseID,
	}

	t.Run("Should return conflict error when order number already exists", func(t *testing.T) {
		server, service, controller := InitPurchaseOrderServer(t)

		server.POST(DefinePath(ResourcePurchaseOrdersUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourcePurchaseOrdersUri), CreateBody(requestObject))

		var serviceReturn *domain.PurchaseOrder
		service.On("Create", requestObject.ToPurchaseOrder()).Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return conflict error when buyer id, order status id, warehouse id, product record id or carrier id does not exists", func(t *testing.T) {
		server, service, controller := InitPurchaseOrderServer(t)

		server.POST(DefinePath(ResourcePurchaseOrdersUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourcePurchaseOrdersUri), CreateBody(requestObject))

		var serviceReturn *domain.PurchaseOrder
		service.On("Create", requestObject.ToPurchaseOrder()).Return(serviceReturn, apperr.NewDependentResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return a created purchase order", func(t *testing.T) {
		server, service, controller := InitPurchaseOrderServer(t)

		server.POST(DefinePath(ResourcePurchaseOrdersUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourcePurchaseOrdersUri), CreateBody(requestObject))

		service.On("Create", requestObject.ToPurchaseOrder()).Return(&mockedPurchaseOrder, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
	})
}

func InitPurchaseOrderServer(t *testing.T) (*gin.Engine, *mocks.Service, *handler.PurchaseOrder) {
	t.Helper()
	server := CreateServer()
	server.Use(middleware.IdValidation())
	service := new(mocks.Service)
	controller := handler.NewPurchaseOrder(service)
	return server, service, controller
}
