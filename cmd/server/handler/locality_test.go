package handler_test

import (
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/middleware"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/locality/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	ResourceLocalitiesUri = "/localities"
)

var (
	mockedLocality = domain.Locality{
		ID: 1,
	}
)

func TestCreateLocality(t *testing.T) {
	requestObject := handler.CreateLocalityRequest{}

	t.Run("Should return conflict error when locality name already exists", func(t *testing.T) {
		server, service, controller := InitLocalityServer(t)

		server.POST(DefinePath(ResourceLocalitiesUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceLocalitiesUri), CreateBody(requestObject))

		var serviceReturn *domain.Locality
		service.On("Create", requestObject.ToLocality()).Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return conflict error when locality id not exists", func(t *testing.T) {
		server, service, controller := InitLocalityServer(t)

		server.POST(DefinePath(ResourceLocalitiesUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceLocalitiesUri), CreateBody(requestObject))

		var serviceReturn *domain.Locality
		service.On("Create", requestObject.ToLocality()).Return(serviceReturn, apperr.NewDependentResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return a created locality", func(t *testing.T) {
		server, service, controller := InitLocalityServer(t)

		server.POST(DefinePath(ResourceLocalitiesUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceLocalitiesUri), CreateBody(requestObject))

		service.On("Create", requestObject.ToLocality()).Return(&mockedLocality, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
	})
}

func TestReportSellers(t *testing.T) {
	t.Run("Should return sellers count report of all localities", func(t *testing.T) {
		server, service, controller := InitLocalityServer(t)

		server.GET(DefinePath(ResourceLocalitiesUri), controller.ReportSellers())
		request, response := MakeRequest("GET", DefinePath(ResourceLocalitiesUri), "")

		service.On("CountSellersByAllLocalities").Return([]domain.SellersByLocalityReport{}, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Should return invalid id error", func(t *testing.T) {
		server, _, controller := InitLocalityServer(t)

		server.GET(DefinePath(ResourceLocalitiesUri), controller.ReportSellers())
		request, response := MakeRequest("GET", DefinePath(ResourceLocalitiesUri)+"?id=abc", "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitLocalityServer(t)

		server.GET(DefinePath(ResourceLocalitiesUri), controller.ReportSellers())
		request, response := MakeRequest("GET", DefinePath(ResourceLocalitiesUri)+"?id=1", "")

		localityId := 1
		var serviceReturn *domain.SellersByLocalityReport
		service.On("CountSellersByLocality", localityId).Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return sellers count report by locality", func(t *testing.T) {
		server, service, controller := InitLocalityServer(t)

		server.GET(DefinePath(ResourceLocalitiesUri), controller.ReportSellers())
		request, response := MakeRequest("GET", DefinePath(ResourceLocalitiesUri)+"?id=1", "")

		localityId := 1
		serviceReturn := domain.SellersByLocalityReport{
			ID:           1,
			LocalityName: "Locality",
			SellersCount: 1,
		}
		service.On("CountSellersByLocality", localityId).Return(&serviceReturn, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestReportCarriers(t *testing.T) {
	t.Run("Should return carriers count report of all localities", func(t *testing.T) {
		server, service, controller := InitLocalityServer(t)

		server.GET(DefinePath(ResourceLocalitiesUri)+"report-carriers", controller.ReportCarriers())
		request, response := MakeRequest("GET", DefinePath(ResourceLocalitiesUri)+"report-carriers", "")

		service.On("CountCarriersByAllLocalities").Return([]domain.CarriersByLocalityReport{}, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Should return invalid id error", func(t *testing.T) {
		server, _, controller := InitLocalityServer(t)

		server.GET(DefinePath(ResourceLocalitiesUri)+"report-carriers", controller.ReportCarriers())
		request, response := MakeRequest("GET", DefinePath(ResourceLocalitiesUri)+"report-carriers?id=abc", "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitLocalityServer(t)

		server.GET(DefinePath(ResourceLocalitiesUri)+"report-carriers", controller.ReportCarriers())
		request, response := MakeRequest("GET", DefinePath(ResourceLocalitiesUri)+"report-carriers?id=1", "")

		localityId := 1
		var serviceReturn *domain.CarriersByLocalityReport
		service.On("CountCarriersByLocality", localityId).Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return carriers count report by locality", func(t *testing.T) {
		server, service, controller := InitLocalityServer(t)

		server.GET(DefinePath(ResourceLocalitiesUri)+"report-carriers", controller.ReportCarriers())
		request, response := MakeRequest("GET", DefinePath(ResourceLocalitiesUri)+"report-carriers?id=1", "")

		localityId := 1
		serviceReturn := domain.CarriersByLocalityReport{
			ID:            1,
			LocalityName:  "Locality",
			CarriersCount: 1,
		}
		service.On("CountCarriersByLocality", localityId).Return(&serviceReturn, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func InitLocalityServer(t *testing.T) (*gin.Engine, *mocks.Service, *handler.Locality) {
	t.Helper()
	server := CreateServer()
	server.Use(middleware.IdValidation())
	service := new(mocks.Service)
	controller := handler.NewLocality(service)
	return server, service, controller
}
