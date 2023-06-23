package buyer_test

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/stretchr/testify/assert"
)

var (
	b = domain.Buyer{
		ID:           1,
		CardNumberID: "123",
		FirstName:    "Teste",
		LastName:     "Teste",
	}
)

func TestServiceCreate(t *testing.T) {
	t.Run("Should return a created buyer", func(t *testing.T) {
		service, repository := CreateService(t)
		id := 1
		repository.On("Save", b).Return(id)
		repository.On("Get", id).Return(&b)
		repository.On("Exists", b.CardNumberID).Return(false)
		result, err := service.Create(context.TODO(), b)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, b, *result)

	})

	t.Run("Should return a conflict error", func(t *testing.T) {
		service, repository := CreateService(t)
		repository.On("Exists", b.CardNumberID).Return(true)
		result, err := service.Create(context.TODO(), b)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})
}

func TestServiceGet(t *testing.T) {
	t.Run("Should return a list of buyers", func(t *testing.T) {
		service, repository := CreateService(t)
		expected := []domain.Buyer{b}
		repository.On("GetAll").Return(expected)
		result := service.GetAll(context.TODO())
		assert.NotEmpty(t, result)
		assert.Equal(t, len(result), 1)
		assert.Equal(t, result[0], b)
	})

	t.Run("Shoul returne a buyer by a especified id", func(t *testing.T) {
		service, repository := CreateService(t)
		id := 1
		repository.On("Get", id).Return(&b)
		result, err := service.Get(context.TODO(), id)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, *result, b)
	})

	t.Run("should return a not found error", func(t *testing.T) {
		service, repository := CreateService(t)
		id := 1
		var repositoryResult *domain.Buyer
		repository.On("Get", id).Return(repositoryResult)
		result, err := service.Get(context.TODO(), id)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})
}

func TestServiceUpdate(t *testing.T) {
	t.Run("Should return a not found error", func(t *testing.T) {
		service, repository := CreateService(t)
		id := 2
		cardNumberId := "12"
		updateBuyer := domain.UpdateBuyer{
			ID:           &id,
			CardNumberID: &cardNumberId,
		}
		var repositoryResult *domain.Buyer
		repository.On("Get", id).Return(repositoryResult)
		result, err := service.Update(context.TODO(), id, updateBuyer)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})

	t.Run("Should return a conflict error", func(t *testing.T) {
		service, repository := CreateService(t)
		id := 1
		cardNumberId := "12"
		updateBuyer := domain.UpdateBuyer{
			ID:           &id,
			CardNumberID: &cardNumberId,
		}
		repository.On("Get", id).Return(&b)
		repository.On("Exists", cardNumberId).Return(true)
		result, err := service.Update(context.TODO(), id, updateBuyer)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})

	t.Run("Should return an updated product", func(t *testing.T) {
		service, repository := CreateService(t)
		id := 1
		cardNumberId := "123"
		updateBuyer := domain.UpdateBuyer{
			ID:           &id,
			CardNumberID: &cardNumberId,
		}

		updatedBuyer := b

		updatedBuyer.Overlap(updateBuyer)

		repository.On("Get", id).Return(&b)
		repository.On("Exists", cardNumberId).Return(true)
		repository.On("Update", updatedBuyer).Return(updatedBuyer)
		repository.On("Get", id).Return(&updatedBuyer)
		result, err := service.Update(context.TODO(), id, updateBuyer)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, cardNumberId, result.CardNumberID)

	})
}

func TestServiceDelete(t *testing.T) {
	t.Run("Should return not found error", func(t *testing.T) {
		service, repository := CreateService(t)
		id := 1
		var repositoryResult *domain.Buyer
		repository.On("Get", id).Return(repositoryResult)
		err := service.Delete(context.TODO(), id)
		assert.Error(t, err)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})

	t.Run("Should delete a product with success", func(t *testing.T) {
		service, repository := CreateService(t)
		id := 1
		repository.On("Get", id).Return(&b)
		repository.On("Delete", id)
		err := service.Delete(context.TODO(), id)
		assert.NoError(t, err)
	})
}

func CreateService(t *testing.T) (buyer.IService, *mocks.Repository) {
	t.Helper()
	repository := new(mocks.Repository)
	service := buyer.NewService(repository)
	return service, repository
}
