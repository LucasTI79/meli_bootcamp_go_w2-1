package seller_test

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/seller/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/stretchr/testify/assert"
)

var (
	p = domain.Seller{
		ID:          1,
		CID:         123,
		CompanyName: "Company Name",
		Address:     "Address",
		Telephone:   "Telephone",
	}
)

func TestServiceCreate(t *testing.T) {
	t.Run("Should return a created seller", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		repository.On("Save", p).Return(id)
		repository.On("Get", id).Return(&p)
		repository.On("Exists", p.CID).Return(false)
		result, err := service.Create(context.TODO(), p)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, p, *result)
	})

	t.Run("Should return a conflict error", func(t *testing.T) {
		service, repository := CreateService(t)

		repository.On("Exists", p.CID).Return(true)
		result, err := service.Create(context.TODO(), p)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})
}

func TestServiceGet(t *testing.T) {
	t.Run("Should return a list of sellers", func(t *testing.T) {
		service, repository := CreateService(t)

		expected := []domain.Seller{p}

		repository.On("GetAll").Return(expected)
		result := service.GetAll(context.TODO())

		assert.NotEmpty(t, result)
		assert.Equal(t, len(result), 1)
		assert.Equal(t, result[0], p)
	})

	t.Run("Should return a seller by specified id", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1

		repository.On("Get", id).Return(&p)
		result, err := service.Get(context.TODO(), id)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, *result, p)
	})

	t.Run("Should return a not found error", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		var respositoryResult *domain.Seller

		repository.On("Get", id).Return(respositoryResult)
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
		cid := 123
		updateSeller := domain.UpdateSeller{
			ID:  &id,
			CID: &cid,
		}

		var respositoryResult *domain.Seller

		repository.On("Get", id).Return(respositoryResult)
		result, err := service.Update(context.TODO(), id, updateSeller)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})

	t.Run("Should return a conflict error", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		cid := 456
		updateSeller := domain.UpdateSeller{
			ID:  &id,
			CID: &cid,
		}

		repository.On("Get", id).Return(&p)
		repository.On("Exists", cid).Return(true)
		result, err := service.Update(context.TODO(), id, updateSeller)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})

	t.Run("Should return an updated seller", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		cid := 123
		companyName := "Company Name 2"
		updateSeller := domain.UpdateSeller{
			ID:          &id,
			CID:         &cid,
			CompanyName: &companyName,
		}
		updatedSeller := p
		updatedSeller.Overlap(updateSeller)

		repository.On("Get", id).Return(&p)
		repository.On("Exists", cid).Return(true)
		repository.On("Update", updatedSeller)
		repository.On("Get", id).Return(&updatedSeller)
		result, err := service.Update(context.TODO(), id, updateSeller)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, companyName, result.CompanyName)
	})
}

func TestServiceDelete(t *testing.T) {
	t.Run("Should return not found error", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		var respositoryResult *domain.Seller

		repository.On("Get", id).Return(respositoryResult)
		err := service.Delete(context.TODO(), id)

		assert.Error(t, err)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})

	t.Run("Should delete a seller with success", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1

		repository.On("Get", id).Return(&p)
		repository.On("Delete", id)
		err := service.Delete(context.TODO(), id)

		assert.NoError(t, err)
	})
}

func CreateService(t *testing.T) (seller.Service, *mocks.Repository) {
	t.Helper()
	repository := new(mocks.Repository)
	service := seller.NewService(repository)
	return service, repository
}
