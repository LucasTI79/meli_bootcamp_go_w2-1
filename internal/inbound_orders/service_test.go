package inbound_orders_test

import (
	"testing"
	"time"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	eMock "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/employee/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/inbound_orders"
	ioMock "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/inbound_orders/mocks"
	pbMock "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_batch/mocks"
	wMock "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/stretchr/testify/assert"
)

var (
	date = "2022-10-01 00:00:00"
	dateString, _ = time.Parse("2006-01-02", date)
	io = domain.InboundOrder{
		ID: 1,
		OrderDate: dateString,
		OrderNumber: "asdf",
		EmployeeId: 1,
		ProductBatchId: 1,
		WarehouseId: 1,
	}
	e = domain.Employee {
		ID: 1,
		CardNumberID: "123456",
		FirstName: "PrimeiroNome",
		LastName: "Sobrenome",
		WarehouseID: 3,
	}

	pb = domain.ProductBatch{
		ID: 1,
	}

	w = domain.Warehouse{
		ID:                 1,
		Address:            "Address",
		Telephone:          "12345",
		WarehouseCode:      "123",
		MinimumCapacity:    1,
		MinimumTemperature: 1,
	}
)

func TestServiceCreate(t *testing.T) {
	t.Run("should return conflict if order number is already registered", func(t *testing.T) {
		service, ioRepository, _, _, _ := CreateService(t)

		orderNumber := "asdf"

		ioRepository.On("Exists", orderNumber).Return(true)
		result, err := service.Create(io)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})
	t.Run("should return dependente resource not found if employee id doesnt exist", func(t *testing.T) {
		service, ioRepository, eRepository, _, _ := CreateService(t)

		orderNumber := "asdf"
		employeeId := 1
		var emptyEmployee *domain.Employee

		ioRepository.On("Exists", orderNumber).Return(false)
		eRepository.On("Get", employeeId).Return(emptyEmployee)

		result, err := service.Create(io)

		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))		
	})
	t.Run("should return dependente resource not found if product batch id doenst exist", func(t *testing.T) {
		service, ioRepository, eRepository, pbRepository, _ := CreateService(t)

		orderNumber := "asdf"
		employeeId := 1
		productBatchId := 1

		var emptyProductBatch *domain.ProductBatch

		ioRepository.On("Exists", orderNumber).Return(false)
		eRepository.On("Get", employeeId).Return(&e)
		pbRepository.On("Get", productBatchId).Return(emptyProductBatch)

		result, err := service.Create(io)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})
	t.Run("should return dependente resource not found if warehouse id doenst exist", func(t *testing.T) {
		service, ioRepository, eRepository, pbRepository, wRepository := CreateService(t)

		orderNumber := "asdf"
		employeeId := 1
		productBatchId := 1
		warehouseId := 1

		var emptyWarehouse *domain.Warehouse

		ioRepository.On("Exists", orderNumber).Return(false)
		eRepository.On("Get", employeeId).Return(&e)
		pbRepository.On("Get", productBatchId).Return(&pb)
		wRepository.On("Get", warehouseId).Return(emptyWarehouse)

		result, err := service.Create(io)

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})
	t.Run("should return dependente resource not found if warehouse id doenst exist", func(t *testing.T) {
		service, ioRepository, eRepository, pbRepository, wRepository := CreateService(t)

		inboundOrderId := 1
		orderNumber := "asdf"
		employeeId := 1
		productBatchId := 1
		warehouseId := 1

		ioRepository.On("Exists", orderNumber).Return(false)
		eRepository.On("Get", employeeId).Return(&e)
		pbRepository.On("Get", productBatchId).Return(&pb)
		wRepository.On("Get", warehouseId).Return(&w)

		ioRepository.On("Save", io).Return(inboundOrderId)
		ioRepository.On("Get", inboundOrderId).Return(&io)

		result, err := service.Create(io)

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, io.ID, result.ID)
	})
}

func CreateService(t *testing.T) (inbound_orders.Service, *ioMock.Repository, *eMock.Repository, *pbMock.Repository, *wMock.Repository) {
	t.Helper()
	ioRepository := new(ioMock.Repository)
	pbRepository := new(pbMock.Repository)
	eRepository := new(eMock.Repository)
	wRepository := new(wMock.Repository)
	service := inbound_orders.NewService(ioRepository, eRepository, pbRepository, wRepository)

	return service, ioRepository, eRepository, pbRepository, wRepository
}