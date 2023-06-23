package employee_test

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/employee/mocks"
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
		repository.On("Save", e).Return(id)
		repository.On("Get", id).Return(&e)
		repository.On("Exists", e.CardNumberID).Return(false)
		result, err := service.Create(context.TODO(), e)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, e, *result)
	})
}


func CreateService(t *testing.T) (employee.Service, *mocks.Repository) {
	t.Helper()
	repository := new(mocks.Repository)
	service := employee.NewService(repository)

	return service, repository
}