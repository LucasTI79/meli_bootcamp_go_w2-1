package handler_test

import (
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/middleware"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	ResourceWarehouseUri = "/warehouses"
)

var (
	mockedWarehouse = domain.Warehouse{
		ID:                 1,
		Address:            "Address",
		Telephone:          "234566",
		WarehouseCode:      "123",
		MinimumCapacity:    1,
		MinimumTemperature: 1,
	}
)

func TestCreateWarehouse(t *testing.T) {
	requestObject := handler.CreateWarehouseRequest{
		Address:            &mockedWarehouse.Address,
		Telephone:          &mockedWarehouse.Telephone,
		WarehouseCode:      &mockedWarehouse.WarehouseCode,
		MinimumCapacity:    &mockedWarehouse.MinimumCapacity,
		MinimumTemperature: &mockedWarehouse.MinimumTemperature,
		LocalityID:         &mockedWarehouse.LocalityID,
	}

	t.Run("Should return conflict error", func(t *testing.T) {
		server, service, controller := InitWarehouseServer(t)

		server.POST(DefinePath(ResourceWarehouseUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceWarehouseUri), CreateBody(requestObject))

		var serviceReturn *domain.Warehouse
		service.On("Create", requestObject.ToWarehouse()).Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return a created warehouse", func(t *testing.T) {
		server, service, controller := InitWarehouseServer(t)

		server.POST(DefinePath(ResourceWarehouseUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceWarehouseUri), CreateBody(requestObject))

		service.On("Create", requestObject.ToWarehouse()).Return(&mockedWarehouse, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
	})
}

func TestGetWarehouse(t *testing.T) {
	t.Run("Should return all warehouses", func(t *testing.T) {
		server, service, controller := InitWarehouseServer(t)

		server.GET(DefinePath(ResourceWarehouseUri), controller.GetAll())
		request, response := MakeRequest("GET", DefinePath(ResourceWarehouseUri), "")

		service.On("GetAll").Return([]domain.Warehouse{})

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitWarehouseServer(t)

		id := 1

		server.GET(DefinePath(ResourceWarehouseUri)+"/:id", controller.Get())
		request, response := MakeRequest("GET", DefinePathWithId(ResourceWarehouseUri, id), "")

		var serviceReturn *domain.Warehouse
		service.On("Get", id).Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return the found warehouse", func(t *testing.T) {
		server, service, controller := InitWarehouseServer(t)

		id := 1

		server.GET(DefinePath(ResourceWarehouseUri)+"/:id", controller.Get())
		request, response := MakeRequest("GET", DefinePathWithId(ResourceWarehouseUri, id), "")

		service.On("Get", id).Return(&mockedWarehouse, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestUpdateWarehouse(t *testing.T) {
	requestObject := handler.UpdateWarehouseRequest{
		Address:            &mockedWarehouse.Address,
		Telephone:          &mockedWarehouse.Telephone,
		WarehouseCode:      &mockedWarehouse.WarehouseCode,
		MinimumCapacity:    &mockedWarehouse.MinimumCapacity,
		MinimumTemperature: &mockedWarehouse.MinimumTemperature,
	}

	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitWarehouseServer(t)

		id := 1

		server.PATCH(DefinePath(ResourceWarehouseUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceWarehouseUri, id), CreateBody(requestObject))

		var serviceReturn *domain.Warehouse
		service.On(
			"Update", id, requestObject.ToUpdateWarehouse()).
			Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return conflict error", func(t *testing.T) {
		server, service, controller := InitWarehouseServer(t)

		id := 1

		server.PATCH(DefinePath(ResourceWarehouseUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceWarehouseUri, id), CreateBody(requestObject))

		var serviceReturn *domain.Warehouse
		service.On(
			"Update", id, requestObject.ToUpdateWarehouse()).
			Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return updated product", func(t *testing.T) {
		server, service, controller := InitWarehouseServer(t)

		id := 1

		server.PATCH(DefinePath(ResourceWarehouseUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceWarehouseUri, id), CreateBody(requestObject))

		service.On(
			"Update", id, requestObject.ToUpdateWarehouse()).
			Return(&mockedWarehouse, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestDeleteWarehouse(t *testing.T) {
	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitWarehouseServer(t)

		id := 1

		server.DELETE(DefinePath(ResourceWarehouseUri)+"/:id", controller.Delete())
		request, response := MakeRequest("DELETE", DefinePathWithId(ResourceWarehouseUri, id), "")

		service.On("Delete", id).Return(apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return success", func(t *testing.T) {
		server, service, controller := InitWarehouseServer(t)

		id := 1

		server.DELETE(DefinePath(ResourceWarehouseUri)+"/:id", controller.Delete())
		request, response := MakeRequest("DELETE", DefinePathWithId(ResourceWarehouseUri, id), "")

		service.On("Delete", id).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})
}

func InitWarehouseServer(t *testing.T) (*gin.Engine, *mocks.Service, *handler.Warehouse) {
	t.Helper()
	server := CreateServer()
	server.Use(middleware.IdValidation())
	service := new(mocks.Service)
	controller := handler.NewWarehouse(service)
	return server, service, controller
}
