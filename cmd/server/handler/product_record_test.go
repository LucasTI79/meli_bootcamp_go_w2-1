package handler_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/middleware"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_record/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/helpers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	ResourceProductRecordsUri = "/product-records"
)

var (
	mockedProductRecord = domain.ProductRecord{
		ID:             1,
		LastUpdateDate: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		ProductID:      1,
		PurchasePrice:  1,
		SalePrice:      1,
	}
)

func TestCreateProductRecord(t *testing.T) {
	lastUpdateDate := helpers.ToFormattedDateTime(mockedProductRecord.LastUpdateDate)

	requestObject := handler.CreateProductRecordRequest{
		LastUpdateDate: &lastUpdateDate,
		ProductID:      &mockedProductRecord.ProductID,
		PurchasePrice:  &mockedProductRecord.PurchasePrice,
		SalePrice:      &mockedProductRecord.SalePrice,
	}

	t.Run("Should return conflict error when a product with the same last update date and id already exists", func(t *testing.T) {
		server, service, controller := InitProductRecordServer(t)

		server.POST(DefinePath(ResourceProductRecordsUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceProductRecordsUri), CreateBody(requestObject))

		var serviceReturn *domain.ProductRecord
		service.On("Create", requestObject.ToProductRecord()).Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return conflict error when product id not exists", func(t *testing.T) {
		server, service, controller := InitProductRecordServer(t)

		server.POST(DefinePath(ResourceProductRecordsUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceProductRecordsUri), CreateBody(requestObject))

		var serviceReturn *domain.ProductRecord
		service.On("Create", requestObject.ToProductRecord()).Return(serviceReturn, apperr.NewDependentResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return a created product record", func(t *testing.T) {
		server, service, controller := InitProductRecordServer(t)

		server.POST(DefinePath(ResourceProductRecordsUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceProductRecordsUri), CreateBody(requestObject))

		service.On("Create", requestObject.ToProductRecord()).Return(&mockedProductRecord, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
	})
}

func TestReportProductRecords(t *testing.T) {
	t.Run("Should return records count report of all products", func(t *testing.T) {
		server, service, controller := InitProductRecordServer(t)

		server.GET(DefinePath(ResourceProductRecordsUri), controller.ReportProductRecords())
		request, response := MakeRequest("GET", DefinePath(ResourceProductRecordsUri), "")

		service.On("CountRecordsByAllProducts").Return([]domain.RecordsByProductReport{}, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Should return invalid id error", func(t *testing.T) {
		server, _, controller := InitProductRecordServer(t)

		server.GET(DefinePath(ResourceProductRecordsUri), controller.ReportProductRecords())
		request, response := MakeRequest("GET", DefinePath(ResourceProductRecordsUri)+"?id=abc", "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitProductRecordServer(t)

		server.GET(DefinePath(ResourceProductRecordsUri), controller.ReportProductRecords())
		request, response := MakeRequest("GET", DefinePath(ResourceProductRecordsUri)+"?id=1", "")

		recordId := 1
		var serviceReturn *domain.RecordsByProductReport
		service.On("CountRecordsByProduct", recordId).Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return records count report by product", func(t *testing.T) {
		server, service, controller := InitProductRecordServer(t)

		server.GET(DefinePath(ResourceProductRecordsUri), controller.ReportProductRecords())
		request, response := MakeRequest("GET", DefinePath(ResourceProductRecordsUri)+"?id=1", "")

		recordId := 1
		serviceReturn := domain.RecordsByProductReport{
			ProductID:    1,
			Description:  "Description",
			RecordsCount: 1,
		}
		service.On("CountRecordsByProduct", recordId).Return(&serviceReturn, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func InitProductRecordServer(t *testing.T) (*gin.Engine, *mocks.Service, *handler.ProductRecord) {
	t.Helper()
	server := CreateServer()
	server.Use(middleware.IdValidation())
	service := new(mocks.Service)
	controller := handler.NewProductRecord(service)
	return server, service, controller
}
