package handler_test

import (
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/middleware"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	ResourceBuyerUri = "/buyers"
)

var (
	mockedBuyer = domain.Buyer{
		ID:           1,
		CardNumberID: "123",
		FirstName:    "Teste",
		LastName:     "Teste",
	}
)

func TestCreateBuyer(t *testing.T) {
	requestObject := handler.CreateBuyerRequest{
		CardNumberID: mockedBuyer.CardNumberID,
		FirstName:    mockedBuyer.FirstName,
		LastName:     mockedBuyer.LastName,
	}

	t.Run("Should return conflict error", func(t *testing.T) {
		server, service, controller := InitBuyerServer(t)

		server.POST(DefinePath(ResourceBuyerUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceBuyerUri), CreateBody(requestObject))

		var serviceReturn *domain.Buyer
		service.On("Create", requestObject.ToBuyer()).Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return a created buyer", func(t *testing.T) {
		server, service, controller := InitBuyerServer(t)

		server.POST(DefinePath(ResourceBuyerUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceBuyerUri), CreateBody(requestObject))

		service.On("Create", requestObject.ToBuyer()).Return(&mockedBuyer, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
	})
}

func TestGetBuyer(t *testing.T) {
	t.Run("Should return all buyers", func(t *testing.T) {
		server, service, controller := InitBuyerServer(t)

		server.GET(DefinePath(ResourceBuyerUri), controller.GetAll())
		request, response := MakeRequest("GET", DefinePath(ResourceBuyerUri), "")

		service.On("GetAll").Return([]domain.Buyer{mockedBuyer})

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Should return bad request error when id is invalid", func(t *testing.T) {
		server, _, controller := InitBuyerServer(t)

		server.GET(DefinePath(ResourceBuyerUri)+"/:id", controller.Get())
		request, response := MakeRequest("GET", DefinePath(ResourceBuyerUri)+"/abc", "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitBuyerServer(t)

		id := 1

		server.GET(DefinePath(ResourceBuyerUri)+"/:id", controller.Get())
		request, response := MakeRequest("GET", DefinePathWithId(ResourceBuyerUri, id), "")

		var serviceReturn *domain.Buyer
		service.On("Get", id).Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return the found buyer", func(t *testing.T) {
		server, service, controller := InitBuyerServer(t)

		id := 1

		server.GET(DefinePath(ResourceBuyerUri)+"/:id", controller.Get())
		request, response := MakeRequest("GET", DefinePathWithId(ResourceBuyerUri, id), "")

		service.On("Get", id).Return(&mockedBuyer, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestUpdateBuyer(t *testing.T) {
	requestObject := handler.UpdateBuyerRequest{
		CardNumberID: &mockedBuyer.CardNumberID,
		FirstName:    &mockedBuyer.FirstName,
		LastName:     &mockedBuyer.LastName,
	}

	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitBuyerServer(t)

		id := 1

		server.PATCH(DefinePath(ResourceBuyerUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceBuyerUri, id), CreateBody(requestObject))

		var serviceReturn *domain.Buyer
		service.On(
			"Update", id, requestObject.ToUpdateBuyer()).
			Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return conflict error", func(t *testing.T) {
		server, service, controller := InitBuyerServer(t)

		id := 1

		server.PATCH(DefinePath(ResourceBuyerUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceBuyerUri, id), CreateBody(requestObject))

		var serviceReturn *domain.Buyer
		service.On(
			"Update", id, requestObject.ToUpdateBuyer()).
			Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return updated buyer", func(t *testing.T) {
		server, service, controller := InitBuyerServer(t)

		id := 1

		server.PATCH(DefinePath(ResourceBuyerUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceBuyerUri, id), CreateBody(requestObject))

		service.On(
			"Update", id, requestObject.ToUpdateBuyer()).
			Return(&mockedBuyer, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestDeleteBuyer(t *testing.T) {
	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitBuyerServer(t)

		id := 1

		server.DELETE(DefinePath(ResourceBuyerUri)+"/:id", controller.Delete())
		request, response := MakeRequest("DELETE", DefinePathWithId(ResourceBuyerUri, id), "")

		service.On("Delete", id).Return(apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return success", func(t *testing.T) {
		server, service, controller := InitBuyerServer(t)

		id := 1

		server.DELETE(DefinePath(ResourceBuyerUri)+"/:id", controller.Delete())
		request, response := MakeRequest("DELETE", DefinePathWithId(ResourceBuyerUri, id), "")

		service.On("Delete", id).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})
}

func InitBuyerServer(t *testing.T) (*gin.Engine, *mocks.Service, *handler.Buyer) {
	t.Helper()
	server := CreateServer()
	server.Use(middleware.IdValidation())
	service := new(mocks.Service)
	controller := handler.NewBuyer(service)
	return server, service, controller
}
