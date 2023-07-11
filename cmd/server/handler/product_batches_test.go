package handler_test

import (
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/middleware"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_batches/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var pb = domain.ProductBatches{
	ID:                 1,
	BatchNumber:        1,
	CurrentQuantity:    1,
	CurrentTemperature: 1,
	DueDate:            "2021-01-01",
	InitialQuantity:    1,
	ManufacturingDate:  "2021-01-01",
	ManufacturingHour:  "00:00:00",
	MinimumTemperature: 1,
	ProductID:          1,
	SectionID:          1,
}

const (
	ResourceProductBatchesUri = "/product-batches"
)

func TestProductBatchesCreate(t *testing.T) {
	requestObject := handler.CreateProductBatchesRequest{
		BatchNumber:        pb.BatchNumber,
		CurrentQuantity:    pb.CurrentQuantity,
		CurrentTemperature: pb.CurrentTemperature,
		DueDate:            pb.DueDate,
		InitialQuantity:    pb.InitialQuantity,
		ManufacturingDate:  pb.ManufacturingDate,
		ManufacturingHour:  pb.ManufacturingHour,
		MinimumTemperature: pb.MinimumTemperature,
		ProductID:          pb.ProductID,
		SectionID:          pb.SectionID,
	}
	t.Run("Should return conflict error", func(t *testing.T) {
		server, service, controller := InitProductBatchesServer(t)

		server.POST(DefinePath(ResourceProductBatchesUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceProductBatchesUri), CreateBody(requestObject))

		var serviceReturn *domain.ProductBatches
		service.On("Create", requestObject.ToProductBatches()).Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})
	t.Run("Should return a created product bacthes", func(t *testing.T) {
		server, service, controller := InitProductBatchesServer(t)

		server.POST(DefinePath(ResourceProductBatchesUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceProductBatchesUri), CreateBody(requestObject))

		service.On("Create", requestObject.ToProductBatches()).Return(&pb, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
	})
}

func TestProductBatchesGet(t *testing.T) {
	t.Run("Should return all products by section", func(t *testing.T) {
		server, service, controller := InitProductBatchesServer(t)

		server.GET(DefinePath(ResourceProductBatchesUri), controller.Get())
		request, response := MakeRequest("GET", DefinePath(ResourceProductBatchesUri), "")

		service.On("Get").Return([]domain.ProductBatches{})

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitProductBatchesServer(t)

		id := 1

		server.GET(DefinePath(ResourceProductBatchesUri)+"/sections/report-products", controller.Get())
		request, response := MakeRequest("GET", DefinePathWithId(ResourceProductBatchesUri, id), "")

		var serviceReturn *domain.ProductBatches
		service.On("Get", id).Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return the found product batches", func(t *testing.T) {
		server, service, controller := InitProductBatchesServer(t)

		server.GET(DefinePath(ResourceProductBatchesUri), controller.Get())
		request, response := MakeRequest("GET", DefinePath(ResourceProductBatchesUri)+"/sections/report-products", "")

		productBatchID := 1
		serviceReturn := domain.ProductsBySectionReport{
			SectionID:     1,
			ProductsCount: 1,
			SectionNumber: 1,
		}
		service.On("Get", productBatchID).Return(&serviceReturn, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func InitProductBatchesServer(t *testing.T) (*gin.Engine, *mocks.Service, *handler.ProductBatches) {
	t.Helper()
	server := CreateServer()
	server.Use(middleware.IdValidation())
	service := new(mocks.Service)
	controller := handler.NewProductBatches(service)
	return server, service, controller
}
