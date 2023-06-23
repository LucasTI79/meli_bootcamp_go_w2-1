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

func initSectionServer(t *testing.T) (*gin.Engine, *mocks.Service, *handler.Section) {
	t.Helper()
	server := CreateServer()
	service := new(mocks.Service)
	controller := handler.NewSection(service)
	return server, service, controller
}