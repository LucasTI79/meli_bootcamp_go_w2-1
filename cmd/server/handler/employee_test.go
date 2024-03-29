package handler_test

import (
	"net/http"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/handler"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/cmd/server/middleware"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/employee/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	ResourceEmployeesUri = "/employees"
	ResourceReportInboundOrdersUri = "/employees/report-inbound-orders"
)

var (
	mockedEmployee = domain.Employee{
		ID:           1,
		CardNumberID: "123456",
		FirstName:    "PrimeiroNome",
		LastName:     "Sobrenome",
		WarehouseID:  2,
	}

	mockedEmployeeInboundOrder = domain.InboundOrdersByEmployee{
		ID:           1,
		CardNumberID: "123456",
		FirstName:    "PrimeiroNome",
		LastName:     "Sobrenome",
		WarehouseID:  2,
		InboundOrdersCount: 1,
	}
)

func TestCreateEmployee(t *testing.T) {
	requestObject := handler.CreateEmployeeRequest{
		CardNumberID: &mockedEmployee.CardNumberID,
		FirstName:    &mockedEmployee.FirstName,
		LastName:     &mockedEmployee.LastName,
		WarehouseID:  &mockedEmployee.WarehouseID,
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

	t.Run("Should return conflict error", func(t *testing.T) {
		server, service, controller := InitEmployeeServer(t)

		server.POST(DefinePath(ResourceEmployeesUri), ValidationMiddleware(requestObject), controller.Create())
		request, response := MakeRequest("POST", DefinePath(ResourceEmployeesUri), CreateBody(requestObject))

		var serviceReturn *domain.Employee
		service.On("Create", requestObject.ToEmployee()).Return(serviceReturn, apperr.NewDependentResourceNotFound(ResourceNotFound))

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

func TestUpdateEmployee(t *testing.T) {
	requestObject := handler.UpdateEmployeeRequest{
		CardNumberID: &mockedEmployee.CardNumberID,
		FirstName:    &mockedEmployee.FirstName,
		LastName:     &mockedEmployee.LastName,
		WarehouseID:  &mockedEmployee.WarehouseID,
	}

	t.Run("Should return employee not found error", func(t *testing.T) {
		server, service, controller := InitEmployeeServer(t)

		id := 99
		var serviceReturn *domain.Employee

		server.PATCH(DefinePath(ResourceEmployeesUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceEmployeesUri, id), CreateBody(requestObject))

		service.On("Update", id, requestObject.ToUpdateEmployee()).Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return conflict error", func(t *testing.T) {
		server, service, controller := InitEmployeeServer(t)

		id := 1
		var serviceReturn *domain.Employee

		server.PATCH(DefinePath(ResourceEmployeesUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceEmployeesUri, id), CreateBody(requestObject))

		service.On("Update", id, requestObject.ToUpdateEmployee()).Return(serviceReturn, apperr.NewResourceAlreadyExists(ResourceAlreadyExists))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return conflict error", func(t *testing.T) {
		server, service, controller := InitEmployeeServer(t)

		id := 1
		var serviceReturn *domain.Employee

		server.PATCH(DefinePath(ResourceEmployeesUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceEmployeesUri, id), CreateBody(requestObject))

		service.On("Update", id, requestObject.ToUpdateEmployee()).Return(serviceReturn, apperr.NewDependentResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusConflict, response.Code)
	})

	t.Run("Should return updated product", func(t *testing.T) {
		server, service, controller := InitEmployeeServer(t)

		id := 1

		server.PATCH(DefinePath(ResourceEmployeesUri)+"/:id", ValidationMiddleware(requestObject), controller.Update())
		request, response := MakeRequest("PATCH", DefinePathWithId(ResourceEmployeesUri, id), CreateBody(requestObject))

		service.On("Update", id, requestObject.ToUpdateEmployee()).Return(&mockedEmployee, nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func TestDeleteEmployee(t *testing.T) {
	t.Run("Should return not found error", func(t *testing.T) {
		server, service, controller := InitEmployeeServer(t)

		id := 99

		server.DELETE(DefinePath(ResourceEmployeesUri)+"/:id", controller.Delete())
		request, response := MakeRequest("DELETE", DefinePathWithId(ResourceEmployeesUri, id), "")

		service.On("Delete", id).Return(apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})

	t.Run("Should return success with no content", func(t *testing.T) {
		server, service, controller := InitEmployeeServer(t)

		id := 1

		server.DELETE(DefinePath(ResourceEmployeesUri)+"/:id", controller.Delete())
		request, response := MakeRequest("DELETE", DefinePathWithId(ResourceEmployeesUri, id), "")

		service.On("Delete", id).Return(nil)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNoContent, response.Code)
	})
}

func TestReportInboundOrders(t *testing.T) {
	t.Run("Should return sucess with all inbound orders by employee if no id was found", func(t *testing.T) {
		server, service, controller := InitEmployeeServer(t)

		server.GET(DefinePath(ResourceReportInboundOrdersUri), controller.ReportInboundOrders())
		request, response := MakeRequest("GET", DefinePath(ResourceReportInboundOrdersUri), "")
		
		service.On("CountInboundOrdersByAllEmployees").Return([]domain.InboundOrdersByEmployee{})

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})
	t.Run("Should return status not found if no employee was found with given id", func(t *testing.T) {
		server, service, controller := InitEmployeeServer(t)

		queryId := "?id=10"
		id := 10
		serviceReturn := &domain.InboundOrdersByEmployee{}

		server.GET(DefinePath(ResourceReportInboundOrdersUri), controller.ReportInboundOrders())
		request, response := MakeRequest("GET", DefinePath(ResourceReportInboundOrdersUri)+queryId, "")

		service.On("CountInboundOrdersByEmployee", id).Return(serviceReturn, apperr.NewResourceNotFound(ResourceNotFound))

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusNotFound, response.Code)
	})
	t.Run("should return error if id format is invalid", func(t *testing.T) {
		server, _, controller := InitEmployeeServer(t)

		queryId := "?id=asdasd"

		server.GET(DefinePath(ResourceReportInboundOrdersUri), controller.ReportInboundOrders())
		request, response := MakeRequest("GET", DefinePath(ResourceReportInboundOrdersUri)+queryId, "")
		server.ServeHTTP(response, request)
		
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
	t.Run("should return employee with inbound orders count if given id is valid", func(t *testing.T) {
		server, service, controller := InitEmployeeServer(t)

		queryId := "?id=1"
		id := 1
		serviceReturn := &domain.InboundOrdersByEmployee{}

		server.GET(DefinePath(ResourceReportInboundOrdersUri), controller.ReportInboundOrders())
		request, response := MakeRequest("GET", DefinePath(ResourceReportInboundOrdersUri)+queryId, "")
		service.On("CountInboundOrdersByEmployee", id).Return(serviceReturn, nil)
		
		server.ServeHTTP(response, request)


		assert.Equal(t, http.StatusOK, response.Code)
	})
}

func InitEmployeeServer(t *testing.T) (*gin.Engine, *mocks.Service, *handler.Employee) {
	t.Helper()
	server := CreateServer()
	server.Use(middleware.IdValidation())
	service := new(mocks.Service)
	controller := handler.NewEmployee(service)
	return server, service, controller
}
