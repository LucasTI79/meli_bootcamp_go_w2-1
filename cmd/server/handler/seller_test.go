package handler_test

import (
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/middleware"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/seller/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	ResourceSellersUri = "/sellers"
)

var (
	mockedSeller = domain.Seller{
		ID:          1,
		CID:         1,
		CompanyName: "Company Name",
		Address:     "Address",
		Telephone:   "Telephone",
	}
)

func TestCreateSeller(t *testing.T) {
	requestObject := handler.CreateSellerRequest{
		CID:         mockedSeller.CID,
		CompanyName: mockedSeller.CompanyName,
		Address:     mockedSeller.Address,
		Telephone:   mockedSeller.Telephone,
	}

	t.Run("Should return conflict error", func(t *testing.T) {
		server, service, controller := InitSellerServer(t)

		server.POST(DefinePath(ResourceSellersUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceSellersUri), CreateBody(requestObject))

		var serviceReturn *domain.Seller
		service.On("Create", requestObject.ToSeller()).Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return a created seller", func(t *testing.T) {
		server, service, controller := InitSellerServer(t)

		server.POST(DefinePath(ResourceSellersUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceSellersUri), CreateBody(requestObject))

		service.On("Create", requestObject.ToSeller()).Return(&mockedSeller, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
	})
}

func TestGetSeller(t *testing.T) {
	t.Run("Should return all sellers", func(t *testing.T) {
		server, service, controller := InitSellerServer(t)

		server.GET(DefinePath(ResourceSellersUri), controller.GetAll())
		request, response := MakeRequest("GET", DefinePath(ResourceSellersUri), "")

		service.On("GetAll").Return([]domain.Seller{})

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitSellerServer(t)

		id := 1

		server.GET(DefinePath(ResourceSellersUri)+"/:id", controller.Get())
		request, response := MakeRequest("GET", DefinePathWithId(ResourceSellersUri, id), "")

		var serviceReturn *domain.Seller
		service.On("Get", id).Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return the found seller", func(t *testing.T) {
		server, service, controller := InitSellerServer(t)

		id := 1

		server.GET(DefinePath(ResourceSellersUri)+"/:id", controller.Get())
		request, response := MakeRequest("GET", DefinePathWithId(ResourceSellersUri, id), "")

		service.On("Get", id).Return(&mockedSeller, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestUpdateSeller(t *testing.T) {
	requestObject := handler.UpdateSellerRequest{
		CID:         &mockedSeller.CID,
		CompanyName: &mockedSeller.CompanyName,
		Address:     &mockedSeller.Address,
		Telephone:   &mockedSeller.Telephone,
	}

	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitSellerServer(t)

		id := 1

		server.PATCH(DefinePath(ResourceSellersUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceSellersUri, id), CreateBody(requestObject))

		var serviceReturn *domain.Seller
		service.On(
			"Update", id, requestObject.ToUpdateSeller()).
			Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return conflict error", func(t *testing.T) {
		server, service, controller := InitSellerServer(t)

		id := 1

		server.PATCH(DefinePath(ResourceSellersUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceSellersUri, id), CreateBody(requestObject))

		var serviceReturn *domain.Seller
		service.On(
			"Update", id, requestObject.ToUpdateSeller()).
			Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return updated seller", func(t *testing.T) {
		server, service, controller := InitSellerServer(t)

		id := 1

		server.PATCH(DefinePath(ResourceSellersUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceSellersUri, id), CreateBody(requestObject))

		service.On(
			"Update", id, requestObject.ToUpdateSeller()).
			Return(&mockedSeller, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestDeleteSeller(t *testing.T) {
	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitSellerServer(t)

		id := 1

		server.DELETE(DefinePath(ResourceSellersUri)+"/:id", controller.Delete())
		request, response := MakeRequest("DELETE", DefinePathWithId(ResourceSellersUri, id), "")

		service.On("Delete", id).Return(apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return success", func(t *testing.T) {
		server, service, controller := InitSellerServer(t)

		id := 1

		server.DELETE(DefinePath(ResourceSellersUri)+"/:id", controller.Delete())
		request, response := MakeRequest("DELETE", DefinePathWithId(ResourceSellersUri, id), "")

		service.On("Delete", id).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})
}

func InitSellerServer(t *testing.T) (*gin.Engine, *mocks.Service, *handler.Seller) {
	t.Helper()
	server := CreateServer()
	server.Use(middleware.IdValidation())
	service := new(mocks.Service)
	controller := handler.NewSeller(service)
	return server, service, controller
}
