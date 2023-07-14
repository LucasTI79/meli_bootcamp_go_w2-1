package handler_test

import (
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/middleware"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/carrier/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var mockedCarrier = domain.Carrier{
	ID:          1,
	CID:         "123",
	CompanyName: "company",
	Address:     "Add",
	Telephone:   "1234678768",
	LocalityID:  1,
}

const (
	ResourceCarriersUri = "/carriers"
)

func TestCreateCarrier(t *testing.T) {
	requestObject := handler.CreateCarrierRequest{
		CID:         &mockedCarrier.CID,
		CompanyName: &mockedCarrier.CompanyName,
		Address:     &mockedCarrier.Address,
		Telephone:   &mockedCarrier.Telephone,
		LocalityID:  &mockedCarrier.LocalityID,
	}

	t.Run("Should return conflict error", func(t *testing.T) {
		server, service, controller := InitCarrierServer(t)

		server.POST(DefinePath(ResourceCarriersUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceCarriersUri), CreateBody(requestObject))

		var serviceReturn *domain.Carrier
		service.On("Create", requestObject.ToCarrier()).Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return conflict error if a dependent resourse was not found", func(t *testing.T) {
		server, service, controller := InitCarrierServer(t)

		server.POST(DefinePath(ResourceCarriersUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceCarriersUri), CreateBody(requestObject))

		var serviceReturn *domain.Carrier
		service.On("Create", requestObject.ToCarrier()).Return(serviceReturn, apperr.NewDependentResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return a created carrier", func(t *testing.T) {
		server, service, controller := InitCarrierServer(t)

		server.POST(DefinePath(ResourceCarriersUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceCarriersUri), CreateBody(requestObject))

		service.On("Create", requestObject.ToCarrier()).Return(&mockedCarrier, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
	})
}

func InitCarrierServer(t *testing.T) (*gin.Engine, *mocks.Service, *handler.Carrier) {
	t.Helper()
	server := CreateServer()
	server.Use(middleware.IdValidation())
	service := new(mocks.Service)
	controller := handler.NewCarrier(service)
	return server, service, controller
}
