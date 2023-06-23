package handler_test

import (
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/employee/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	ResourceEmployeesUri = "/employees"
)

var (
	mockedEmployee = domain.Employee{
		ID: 1,
		CardNumberID: "123456",
		FirstName: "PrimeiroNome",
		LastName: "Sobrenome",
		WarehouseID: 2,
	}
)

func TestCreateEmployee(t *testing.T) {
	requestObject := handler.CreateEmployeeRequest{
		CardNumberID: &mockedEmployee.CardNumberID,
		FirstName: &mockedEmployee.FirstName,
		LastName: &mockedEmployee.LastName,
		WarehouseID: &mockedEmployee.WarehouseID,
	}

	t.Run("Should return conflict error", func(t *testing.T) {
		server, service, controller := InitEmployeeServer(t)

		server.POST(DefinePath(ResourceEmployeesUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceEmployeesUri), CreateBody(requestObject))

		var serviceReturn *domain.Employee
		service.On("Create", requestObject.ToEmployee()).Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})
}

func InitEmployeeServer(t *testing.T) (*gin.Engine, *mocks.Service, *handler.Employee) {
	t.Helper()
	server := CreateServer()
	service := new(mocks.Service)
	controller := handler.NewEmployee(service)
	return server, service, controller
}