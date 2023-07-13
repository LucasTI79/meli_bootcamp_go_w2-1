package handler_test

import (
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/middleware"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	ResourceProductsUri = "/products"
)

var (
	mockedProduct = domain.Product{
		ID:             1,
		Description:    "Description",
		ExpirationRate: 1,
		FreezingRate:   1,
		Height:         1,
		Length:         1,
		Netweight:      1,
		ProductCode:    "123",
		RecomFreezTemp: 1,
		Width:          1,
		ProductTypeID:  1,
		SellerID:       1,
	}
)

func TestCreateProduct(t *testing.T) {
	requestObject := handler.CreateProductRequest{
		Description:    &mockedProduct.Description,
		ExpirationRate: &mockedProduct.ExpirationRate,
		FreezingRate:   &mockedProduct.FreezingRate,
		Height:         &mockedProduct.Height,
		Length:         &mockedProduct.Length,
		Netweight:      &mockedProduct.Netweight,
		ProductCode:    &mockedProduct.ProductCode,
		RecomFreezTemp: &mockedProduct.RecomFreezTemp,
		Width:          &mockedProduct.Width,
		ProductTypeID:  &mockedProduct.ProductTypeID,
		SellerID:       &mockedProduct.SellerID,
	}

	t.Run("Should return conflict error", func(t *testing.T) {
		server, service, controller := InitProductServer(t)

		server.POST(DefinePath(ResourceProductsUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceProductsUri), CreateBody(requestObject))

		var serviceReturn *domain.Product
		service.On("Create", requestObject.ToProduct()).Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return dependent resource not found error", func(t *testing.T) {
		server, service, controller := InitProductServer(t)

		server.POST(DefinePath(ResourceProductsUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceProductsUri), CreateBody(requestObject))

		var serviceReturn *domain.Product
		service.On("Create", requestObject.ToProduct()).Return(serviceReturn, apperr.NewDependentResourceNotFound(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return a created product", func(t *testing.T) {
		server, service, controller := InitProductServer(t)

		server.POST(DefinePath(ResourceProductsUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceProductsUri), CreateBody(requestObject))

		service.On("Create", requestObject.ToProduct()).Return(&mockedProduct, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
	})
}

func TestGetProduct(t *testing.T) {
	t.Run("Should return all products", func(t *testing.T) {
		server, service, controller := InitProductServer(t)

		server.GET(DefinePath(ResourceProductsUri), controller.GetAll())
		request, response := MakeRequest("GET", DefinePath(ResourceProductsUri), "")

		service.On("GetAll").Return([]domain.Product{})

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitProductServer(t)

		id := 1

		server.GET(DefinePath(ResourceProductsUri)+"/:id", controller.Get())
		request, response := MakeRequest("GET", DefinePathWithId(ResourceProductsUri, id), "")

		var serviceReturn *domain.Product
		service.On("Get", id).Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return the found product", func(t *testing.T) {
		server, service, controller := InitProductServer(t)

		id := 1

		server.GET(DefinePath(ResourceProductsUri)+"/:id", controller.Get())
		request, response := MakeRequest("GET", DefinePathWithId(ResourceProductsUri, id), "")

		service.On("Get", id).Return(&mockedProduct, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestUpdateProduct(t *testing.T) {
	requestObject := handler.UpdateProductRequest{
		Description:    &mockedProduct.Description,
		ExpirationRate: &mockedProduct.ExpirationRate,
		FreezingRate:   &mockedProduct.FreezingRate,
		Height:         &mockedProduct.Height,
		Length:         &mockedProduct.Length,
		Netweight:      &mockedProduct.Netweight,
		ProductCode:    &mockedProduct.ProductCode,
		RecomFreezTemp: &mockedProduct.RecomFreezTemp,
		Width:          &mockedProduct.Width,
		ProductTypeID:  &mockedProduct.ProductTypeID,
		SellerID:       &mockedProduct.SellerID,
	}

	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitProductServer(t)

		id := 1

		server.PATCH(DefinePath(ResourceProductsUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceProductsUri, id), CreateBody(requestObject))

		var serviceReturn *domain.Product
		service.On(
			"Update", id, requestObject.ToUpdateProduct()).
			Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return conflict error", func(t *testing.T) {
		server, service, controller := InitProductServer(t)

		id := 1

		server.PATCH(DefinePath(ResourceProductsUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceProductsUri, id), CreateBody(requestObject))

		var serviceReturn *domain.Product
		service.On(
			"Update", id, requestObject.ToUpdateProduct()).
			Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return dependent resource not found error", func(t *testing.T) {
		server, service, controller := InitProductServer(t)

		id := 1

		server.PATCH(DefinePath(ResourceProductsUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceProductsUri, id), CreateBody(requestObject))

		var serviceReturn *domain.Product
		service.On(
			"Update", id, requestObject.ToUpdateProduct()).
			Return(serviceReturn, apperr.NewDependentResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return updated product", func(t *testing.T) {
		server, service, controller := InitProductServer(t)

		id := 1

		server.PATCH(DefinePath(ResourceProductsUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceProductsUri, id), CreateBody(requestObject))

		service.On(
			"Update", id, requestObject.ToUpdateProduct()).
			Return(&mockedProduct, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitProductServer(t)

		id := 1

		server.DELETE(DefinePath(ResourceProductsUri)+"/:id", controller.Delete())
		request, response := MakeRequest("DELETE", DefinePathWithId(ResourceProductsUri, id), "")

		service.On("Delete", id).Return(apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return success", func(t *testing.T) {
		server, service, controller := InitProductServer(t)

		id := 1

		server.DELETE(DefinePath(ResourceProductsUri)+"/:id", controller.Delete())
		request, response := MakeRequest("DELETE", DefinePathWithId(ResourceProductsUri, id), "")

		service.On("Delete", id).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})
}

func TestReportProductRecords(t *testing.T) {
	t.Run("Should return records count report of all products", func(t *testing.T) {
		server, service, controller := InitProductServer(t)

		server.GET(DefinePath(ResourceProductRecordsUri), controller.ReportRecords())
		request, response := MakeRequest("GET", DefinePath(ResourceProductRecordsUri), "")

		service.On("CountRecordsByAllProducts").Return([]domain.RecordsByProductReport{}, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Should return invalid id error", func(t *testing.T) {
		server, _, controller := InitProductServer(t)

		server.GET(DefinePath(ResourceProductRecordsUri), controller.ReportRecords())
		request, response := MakeRequest("GET", DefinePath(ResourceProductRecordsUri)+"?id=abc", "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitProductServer(t)

		server.GET(DefinePath(ResourceProductRecordsUri), controller.ReportRecords())
		request, response := MakeRequest("GET", DefinePath(ResourceProductRecordsUri)+"?id=1", "")

		recordId := 1
		var serviceReturn *domain.RecordsByProductReport
		service.On("CountRecordsByProduct", recordId).Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return records count report by product", func(t *testing.T) {
		server, service, controller := InitProductServer(t)

		server.GET(DefinePath(ResourceProductRecordsUri), controller.ReportRecords())
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

func InitProductServer(t *testing.T) (*gin.Engine, *mocks.Service, *handler.Product) {
	t.Helper()
	server := CreateServer()
	server.Use(middleware.IdValidation())
	service := new(mocks.Service)
	controller := handler.NewProduct(service)
	return server, service, controller
}
