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
}

func CreateService(t *testing.T) (employee.Service, *mocks.Repository) {
	t.Helper()
	repository := new(mocks.Repository)
	service := employee.NewService(repository)

	return service, repository
}