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

const (
	resourceProductsBatchesUri = "/products-batches"
)

var (
	productBatch = domain.ProductBatches{
		BatchNumber:        1,
		CurrentQuantity:    1,
		CurrentTemperature: 2,
		DueDate:            "2021-01-01",
		InitialQuantity:    10,
		ManufacturingDate:  "2021-01-01",
		ManufacturingHour:  10,
		MinimumTemperature: 0,
		ProductID:          1,
		SectionID:          1,
	}
)

func TestCreateProductBatches(t *testing.T) {
	requestObject := handler.CreateProductBatchesRequest{
		BatchNumber:        &productBatch.BatchNumber,
		CurrentQuantity:    &productBatch.CurrentQuantity,
		CurrentTemperature: &productBatch.CurrentTemperature,
		DueDate:            &productBatch.DueDate,
		InitialQuantity:    &productBatch.InitialQuantity,
		ManufacturingDate:  &productBatch.ManufacturingDate,
		ManufacturingHour:  &productBatch.ManufacturingHour,
		MinimumTemperature: &productBatch.MinimumTemperature,
		ProductID:          &productBatch.ProductID,
		SectionID:          &productBatch.SectionID,
	}
	t.Run("Should return conflict error when product batch already exists", func(t *testing.T) {
		server, service, controller := InitProductBatchesServer(t)

		server.POST(DefinePath(resourceProductsBatchesUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(resourceProductsBatchesUri), CreateBody(requestObject))

		var productBatchReturn *domain.ProductBatches
		service.On("Create", requestObject.ToProductBatches()).Return(productBatchReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})
	t.Run("Should return conflict error when product batch id not exists", func(t *testing.T) {
		server, service, controller := InitProductBatchesServer(t)

		server.POST(DefinePath(resourceProductsBatchesUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(resourceProductsBatchesUri), CreateBody(requestObject))

		var productBatchReturn *domain.ProductBatches
		service.On("Create", requestObject.ToProductBatches()).Return(productBatchReturn, apperr.NewDependentResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})
	t.Run("Should return created product batch", func(t *testing.T) {
		server, service, controller := InitProductBatchesServer(t)

		server.POST(DefinePath(resourceProductsBatchesUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(resourceProductsBatchesUri), CreateBody(requestObject))

		service.On("Create", requestObject.ToProductBatches()).Return(&productBatch, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
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
