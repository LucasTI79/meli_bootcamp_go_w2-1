package employee_test

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/employee/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/stretchr/testify/assert"
)

var (
	e = domain.Employee {
		ID: 1,
		CardNumberID: "123456",
		FirstName: "PrimeiroNome",
		LastName: "Sobrenome",
		WarehouseID: 3,
	}
)

func TestServiceCreate(t *testing.T) {
	t.Run("Should return a created employee", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		repository.On("Exists", e.CardNumberID).Return(false)
		repository.On("Save", e).Return(id)
		repository.On("Get", id).Return(&e)
		result, err := service.Create(context.TODO(), e)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, e, *result)
	})

	t.Run("Should return a conflict error", func(t *testing.T) {
		service, repository := CreateService(t)

		repository.On("Exists", e.CardNumberID).Return(true)
		result, err := service.Create(context.TODO(), e)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})
}

func TestServiceGet(t *testing.T) {
	t.Run("Should return a list of employees", func(t *testing.T) {
		service, repository := CreateService(t)

		expectedResult := []domain.Employee{e}

		repository.On("GetAll").Return(expectedResult)
		result := service.GetAll(context.TODO())

		assert.NotEmpty(t, result)
		assert.True(t, len(result) >= 1)
		assert.Equal(t, result[0].CardNumberID, e.CardNumberID)
		assert.Equal(t, result[0].FirstName, e.FirstName)
		assert.Equal(t, result[0].LastName, e.LastName)
		assert.Equal(t, result[0].WarehouseID, e.WarehouseID)
	})

	t.Run("Should return a employee by id", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1

		repository.On("Get", id).Return(&e)
		result, err := service.Get(context.TODO(), id)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, result.ID, e.ID)
		assert.Equal(t, *result, e)
	})

	t.Run("Should return a resource not found error", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 99
		var emptyEmployee *domain.Employee

		repository.On("Get", id).Return(emptyEmployee)
		result, err := service.Get(context.TODO(), id)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})
}

func TestServiceUpdate(t *testing.T) {
	t.Run("Should return a not found error", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 99
		cardNumberID := "444555"
		updateEmployee := domain.UpdateEmployee {
			ID: &id,
			CardNumberID: &cardNumberID,
		}

		var emptyEmployee *domain.Employee

		repository.On("Get", id).Return(emptyEmployee)
		result, err := service.Update(context.TODO(), id, updateEmployee)

		assert.Nil(t, result)
		assert.Error(t, err)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})

	t.Run("Should return a conflict error", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		cardNumberID := "555555"
		updateEmployee := domain.UpdateEmployee{
			ID: &id,
			CardNumberID: &cardNumberID,
		}

		repository.On("Get", id).Return(&e)
		repository.On("Exists", cardNumberID).Return(true)

		result, err := service.Update(context.TODO(), id, updateEmployee)

		assert.Nil(t, result)
		assert.Error(t, err)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})

	t.Run("Should return an updated employee", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		cardNumberID := "123456"
		firstName := "Cleber"
		updateEmployee := domain.UpdateEmployee{
			ID: &id,
			CardNumberID: &cardNumberID,
			FirstName: &firstName,
		}

		updatedEmployee := e
		
		repository.On("Get", id).Return(&e)
		repository.On("Exists", cardNumberID).Return(true)
		updatedEmployee.Overlap(updateEmployee)
		repository.On("Update", updatedEmployee)
		repository.On("Get", id).Return(&updatedEmployee)

		result, err := service.Update(context.TODO(), id, updateEmployee)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, firstName, result.FirstName)
	})
}

func TestServiceDelete(t *testing.T) {
	t.Run("Should return not found error", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 99
		var emptyEmployee *domain.Employee

		repository.On("Get", id).Return(emptyEmployee)
		err := service.Delete(context.TODO(), id)

		assert.Error(t, err)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})
}

func CreateService(t *testing.T) (employee.Service, *mocks.Repository) {
	t.Helper()
	repository := new(mocks.Repository)
	service := employee.NewService(repository)

	return service, repository
}