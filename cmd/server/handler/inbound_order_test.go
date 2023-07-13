package handler_test

import (
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/middleware"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/inbound_order/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	ResourceInboundOrdersUri = "/inbound-orders"
)

var (
	date               = "2022-10-01 00:00:00"
	dateString         = helpers.ToDateTime(date)
	mockedInboundOrder = domain.InboundOrder{
		ID:             1,
		OrderDate:      dateString,
		OrderNumber:    "asdf",
		EmployeeId:     1,
		ProductBatchId: 1,
		WarehouseId:    1,
	}
)

func TestCreateInboundOrder(t *testing.T) {
	requestObject := handler.CreateInboundOrderRequest{
		OrderDate:      &date,
		OrderNumber:    &mockedInboundOrder.OrderNumber,
		EmployeeId:     &mockedInboundOrder.EmployeeId,
		ProductBatchId: &mockedInboundOrder.ProductBatchId,
		WarehouseId:    &mockedEmployeeInboundOrder.WarehouseID,
	}

	t.Run("should return 409 conflict when order number already exists", func(t *testing.T) {
		server, service, controller := InitInboundOrdersServer(t)

		server.POST(DefinePath(ResourceInboundOrdersUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceInboundOrdersUri), CreateBody(requestObject))

		var serviceReturn *domain.InboundOrder
		service.On("Create", requestObject.ToInboundOrder()).Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("should return status conflict when a dependent resource does not exist", func(t *testing.T) {
		server, service, controller := InitInboundOrdersServer(t)

		server.POST(DefinePath(ResourceInboundOrdersUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceInboundOrdersUri), CreateBody(requestObject))

		var serviceReturn *domain.InboundOrder
		service.On("Create", requestObject.ToInboundOrder()).Return(serviceReturn, apperr.NewDependentResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("should return status 201 success", func(t *testing.T) {
		server, service, controller := InitInboundOrdersServer(t)

		server.POST(DefinePath(ResourceInboundOrdersUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceInboundOrdersUri), CreateBody(requestObject))

		var serviceReturn *domain.InboundOrder = &mockedInboundOrder
		service.On("Create", requestObject.ToInboundOrder()).Return(serviceReturn, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
	})
}

func InitInboundOrdersServer(t *testing.T) (*gin.Engine, *mocks.Service, *handler.InboundOrder) {
	t.Helper()
	server := CreateServer()
	server.Use(middleware.IdValidation())
	service := new(mocks.Service)
	controller := handler.NewInboundOrder(service)
	return server, service, controller
}
