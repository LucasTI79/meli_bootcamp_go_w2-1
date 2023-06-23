package handler_test

import (
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	ResourceUri           = "/products"
	ResourceAlreadyExists = "resource already exists"
	ResourceNotFound      = "resource not found"
)

var (
	p = domain.Product{
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
		Description:    &p.Description,
		ExpirationRate: &p.ExpirationRate,
		FreezingRate:   &p.FreezingRate,
		Height:         &p.Height,
		Length:         &p.Length,
		Netweight:      &p.Netweight,
		ProductCode:    &p.ProductCode,
		RecomFreezTemp: &p.RecomFreezTemp,
		Width:          &p.Width,
		ProductTypeID:  &p.ProductTypeID,
		SellerID:       &p.SellerID,
	}

	t.Run("Should return conflict error", func(t *testing.T) {
		server, service, controller := InitServer(t)

		server.POST(DefinePath(ResourceUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceUri), CreateBody(requestObject))

		var serviceReturn *domain.Product
		service.On("Create", requestObject.ToProduct()).Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return a created product", func(t *testing.T) {
		server, service, controller := InitServer(t)

		server.POST(DefinePath(ResourceUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceUri), CreateBody(requestObject))

		service.On("Create", requestObject.ToProduct()).Return(&p, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
	})
}

func TestGetProduct(t *testing.T) {
	t.Run("Should return all products", func(t *testing.T) {
		server, service, controller := InitServer(t)

		server.GET(DefinePath(ResourceUri), controller.GetAll())
		request, response := MakeRequest("GET", DefinePath(ResourceUri), "")

		service.On("GetAll").Return([]domain.Product{})

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Should return bad request error when id is invalid", func(t *testing.T) {
		server, _, controller := InitServer(t)

		server.GET(DefinePath(ResourceUri)+"/:id", controller.Get())
		request, response := MakeRequest("GET", DefinePath(ResourceUri)+"/abc", "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitServer(t)

		id := 1

		server.GET(DefinePath(ResourceUri)+"/:id", controller.Get())
		request, response := MakeRequest("GET", DefinePathWithId(ResourceUri, id), "")

		var serviceReturn *domain.Product
		service.On("Get", id).Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return the found product", func(t *testing.T) {
		server, service, controller := InitServer(t)

		id := 1

		server.GET(DefinePath(ResourceUri)+"/:id", controller.Get())
		request, response := MakeRequest("GET", DefinePathWithId(ResourceUri, id), "")

		service.On("Get", id).Return(&p, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestUpdateProduct(t *testing.T) {
	requestObject := handler.UpdateProductRequest{
		Description:    &p.Description,
		ExpirationRate: &p.ExpirationRate,
		FreezingRate:   &p.FreezingRate,
		Height:         &p.Height,
		Length:         &p.Length,
		Netweight:      &p.Netweight,
		ProductCode:    &p.ProductCode,
		RecomFreezTemp: &p.RecomFreezTemp,
		Width:          &p.Width,
		ProductTypeID:  &p.ProductTypeID,
		SellerID:       &p.SellerID,
	}
	t.Run("Should return bad request error when id is invalid", func(t *testing.T) {
		server, _, controller := InitServer(t)

		server.PATCH(DefinePath(ResourceUri)+"/:id", controller.Update())
		request, response := MakeRequest("PATCH", DefinePath(ResourceUri)+"/abc", CreateBody(requestObject))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return bad request error when request is blank", func(t *testing.T) {
		server, _, controller := InitServer(t)

		var requestObject handler.UpdateProductRequest
		server.PATCH(DefinePath(ResourceUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceUri, 1), CreateBody(requestObject))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitServer(t)

		id := 1

		server.PATCH(DefinePath(ResourceUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceUri, id), CreateBody(requestObject))

		var serviceReturn *domain.Product
		service.On(
			"Update", id, requestObject.ToUpdateProduct()).
			Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("Should return conflict error", func(t *testing.T) {
		server, service, controller := InitServer(t)

		id := 1

		server.PATCH(DefinePath(ResourceUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceUri, id), CreateBody(requestObject))

		var serviceReturn *domain.Product
		service.On(
			"Update", id, requestObject.ToUpdateProduct()).
			Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return updated product", func(t *testing.T) {
		server, service, controller := InitServer(t)

		id := 1

		server.PATCH(DefinePath(ResourceUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceUri, id), CreateBody(requestObject))

		service.On(
			"Update", id, requestObject.ToUpdateProduct()).
			Return(&p, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestDeleteProduct(t *testing.T) {

	t.Run("Should return bad request error when id is invalid", func(t *testing.T) {
		server, _, controller := InitServer(t)

		server.DELETE(DefinePath(ResourceUri)+"/:id", controller.Delete())
		request, response := MakeRequest("DELETE", DefinePath(ResourceUri)+"/abc", "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitServer(t)

		id := 1

		server.DELETE(DefinePath(ResourceUri)+"/:id", controller.Delete())
		request, response := MakeRequest("DELETE", DefinePathWithId(ResourceUri, id), "")

		service.On("Delete", id).Return(apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return success", func(t *testing.T) {
		server, service, controller := InitServer(t)

		id := 1

		server.DELETE(DefinePath(ResourceUri)+"/:id", controller.Delete())
		request, response := MakeRequest("DELETE", DefinePathWithId(ResourceUri, id), "")

		service.On("Delete", id).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})
}

func InitServer(t *testing.T) (*gin.Engine, *mocks.Service, *handler.Product) {
	t.Helper()
	server := CreateServer()
	service := new(mocks.Service)
	controller := handler.NewProduct(service)
	return server, service, controller
}
