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
	t.Run("Should return a created employee", func(t *testing.T) {
		server, service, controller := InitEmployeeServer(t)

		server.POST(DefinePath(ResourceEmployeesUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceEmployeesUri), CreateBody(requestObject))

		service.On("Create", requestObject.ToEmployee()).Return(&mockedEmployee, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusCreated, response.Code)
	})
}

func TestGetEmployee(t *testing.T) {
	t.Run("Should return a list of all employees", func(t *testing.T) {
		server, service, controller := InitEmployeeServer(t)

		server.GET(DefinePath(ResourceEmployeesUri), controller.GetAll())
		request, response := MakeRequest("GET", DefinePath(ResourceEmployeesUri), "")

		service.On("GetAll").Return([]domain.Employee{})

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Should return bad request error when ID is invalid", func(t *testing.T) {
		server, _, controller := InitEmployeeServer(t)

		server.GET(DefinePath(ResourceEmployeesUri)+"/:id", controller.Get())
		request, response := MakeRequest("GET", DefinePath(ResourceEmployeesUri)+"/teste", "")

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("Should return employee not found error", func(t *testing.T) {
		server, service, controller := InitEmployeeServer(t)

		id := 99

		server.GET(DefinePath(ResourceEmployeesUri)+"/:id", controller.Get())
		request, response := MakeRequest("GET", DefinePathWithId(ResourceEmployeesUri, id), "")

		var serviceReturn *domain.Employee
		service.On("Get", id).Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return an employee based on ID", func(t *testing.T) {
		server, service, controller := InitEmployeeServer(t)

		id := 1

		server.GET(DefinePath(ResourceEmployeesUri)+"/:id", controller.Get())
		request, response := MakeRequest("GET", DefinePathWithId(ResourceEmployeesUri, id), "")

		service.On("Get", id).Return(&mockedEmployee, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func InitEmployeeServer(t *testing.T) (*gin.Engine, *mocks.Service, *handler.Employee) {
	t.Helper()
	server := CreateServer()
	service := new(mocks.Service)
	controller := handler.NewEmployee(service)
	return server, service, controller
}