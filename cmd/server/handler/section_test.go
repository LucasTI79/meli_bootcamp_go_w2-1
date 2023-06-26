package handler_test

import (
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/section/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	resourceSectionUri = "/sections"
)

var (
	s = domain.Section{
		ID: 1,
		SectionNumber: 1,      
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity: 1,
		MinimumCapacity: 1,
		MaximumCapacity: 1,
		WarehouseID: 1,
		ProductTypeID: 1,      
	}
)

func TestCreateSection(t *testing.T) {
	requestObject := handler.CreateSectionRequest{
		SectionNumber: &s.SectionNumber,      
		CurrentTemperature: &s.CurrentTemperature,
		MinimumTemperature: &s.MinimumTemperature,
		CurrentCapacity: &s.CurrentCapacity,
		MinimumCapacity: &s.MinimumCapacity,
		MaximumCapacity: &s.MaximumCapacity,
		WarehouseID: &s.WarehouseID,
		ProductTypeID: &s.ProductTypeID,
	}

	t.Run("Should return conflict error", func(t *testing.T) {
		server, service, controller := initSectionServer(t)

		server.POST(DefinePath(resourceSectionUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(resourceSectionUri), CreateBody(requestObject))

		var serviceReturn *domain.Section
		service.On("Create", requestObject.ToSection()).Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return a created section", func(t *testing.T) {
		server, service, controller := initSectionServer(t)

		server.POST(DefinePath(resourceSectionUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(resourceSectionUri), CreateBody(requestObject))

		service.On("Create", requestObject.ToSection()).Return(&s, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
	})
}

func TestGetSection(t *testing.T) {
	t.Run("Should return all sections", func(t *testing.T) {
		server, service, controller := initSectionServer(t)

		server.GET(DefinePath(resourceSectionUri), controller.GetAll())
		request, response := MakeRequest("GET", DefinePath(resourceSectionUri), "")

		service.On("GetAll").Return([]domain.Section{})

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Should return bad request error when id is invalid", func(t *testing.T) {
		server, _, controller := initSectionServer(t)

		server.GET(DefinePath(resourceSectionUri)+"/:id", controller.Get())
		request, response := MakeRequest("GET", DefinePath(resourceSectionUri)+"/abc", "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := initSectionServer(t)

		id := 1

		server.GET(DefinePath(resourceSectionUri)+"/:id", controller.Get())
		request, response := MakeRequest("GET", DefinePathWithId(resourceSectionUri, id), "")

		var serviceReturn *domain.Section
		service.On("Get", id).Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return the found section", func(t *testing.T) {
		server, service, controller := initSectionServer(t)

		id := 1

		server.GET(DefinePath(resourceSectionUri)+"/:id", controller.Get())
		request, response := MakeRequest("GET", DefinePathWithId(resourceSectionUri, id), "")

		service.On("Get", id).Return(&s, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestUpdateSection(t *testing.T) {
	requestObject := handler.UpdateSectionRequest{
		SectionNumber: &s.SectionNumber,      
		CurrentTemperature: &s.CurrentTemperature,
		MinimumTemperature: &s.MinimumTemperature,
		CurrentCapacity: &s.CurrentCapacity,
		MinimumCapacity: &s.MinimumCapacity,
		MaximumCapacity: &s.MaximumCapacity,
		WarehouseID: &s.WarehouseID,
		ProductTypeID: &s.ProductTypeID, 
	}
	t.Run("Should return bad request error when id is invalid", func(t *testing.T) {
		server, _, controller := initSectionServer(t)

		server.PATCH(DefinePath(resourceSectionUri)+"/:id", controller.Update())
		request, response := MakeRequest("PATCH", DefinePath(resourceSectionUri)+"/abc", CreateBody(requestObject))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return bad request error when request is blank", func(t *testing.T) {
		server, _, controller := initSectionServer(t)

		var requestObject handler.UpdateSectionRequest
		server.PATCH(DefinePath(resourceSectionUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(resourceSectionUri, 1), CreateBody(requestObject))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := initSectionServer(t)

		id := 2

		server.PATCH(DefinePath(resourceSectionUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(resourceSectionUri, id), CreateBody(requestObject))

		var serviceReturn *domain.Section
		service.On(
			"Update", id, requestObject.ToUpdateSection()).
			Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("Should return conflict error", func(t *testing.T) {
		server, service, controller := initSectionServer(t)

		id := 1

		server.PATCH(DefinePath(resourceSectionUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(resourceSectionUri, id), CreateBody(requestObject))

		var serviceReturn *domain.Section
		service.On(
			"Update", id, requestObject.ToUpdateSection()).
			Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return updated section", func(t *testing.T) {
		server, service, controller := initSectionServer(t)

		id := 1

		server.PATCH(DefinePath(resourceSectionUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(resourceSectionUri, id), CreateBody(requestObject))

		service.On(
			"Update", id, requestObject.ToUpdateSection()).
			Return(&s, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestDeleteSection(t *testing.T) {

	t.Run("Should return bad request error when id is invalid", func(t *testing.T) {
		server, _, controller := initSectionServer(t)

		server.DELETE(DefinePath(resourceSectionUri)+"/:id", controller.Delete())
		request, response := MakeRequest("DELETE", DefinePath(resourceSectionUri)+"/abc", "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := initSectionServer(t)

		id := 2

		server.DELETE(DefinePath(resourceSectionUri)+"/:id", controller.Delete())
		request, response := MakeRequest("DELETE", DefinePathWithId(resourceSectionUri, id), "")

		service.On("Delete", id).Return(apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return success", func(t *testing.T) {
		server, service, controller := initSectionServer(t)

		id := 1

		server.DELETE(DefinePath(resourceSectionUri)+"/:id", controller.Delete())
		request, response := MakeRequest("DELETE", DefinePathWithId(resourceSectionUri, id), "")

		service.On("Delete", id).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})
}

func initSectionServer(t *testing.T) (*gin.Engine, *mocks.Service, *handler.Section) {
	t.Helper()
	server := CreateServer()
	service := new(mocks.Service)
	controller := handler.NewSection(service)
	return server, service, controller
}