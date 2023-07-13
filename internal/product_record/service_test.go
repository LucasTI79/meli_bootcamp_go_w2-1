package product_record_test

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	productMocks "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product/mocks"
	record "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_record"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_record/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/stretchr/testify/assert"
)

var (
	mockedProductTemplate = domain.Product{
		ID: 1,
	}
)

func TestServiceCreate(t *testing.T) {
	t.Run("Should return a created record", func(t *testing.T) {
		service, repository, productRepository := CreateService(t)

		mockedProductRecord := mockedProductRecordTemplate
		mockedProduct := mockedProductTemplate
		id := 1
		recordId := 1

		repository.On("Exists", mockedProductRecord.ProductID, mockedProductRecord.LastUpdateDate).Return(false)
		productRepository.On("Get", recordId).Return(&mockedProduct)
		repository.On("Save", mockedProductRecord).Return(id)
		repository.On("Get", id).Return(&mockedProductRecord)
		result, err := service.Create(mockedProductRecord)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, mockedProductRecord, *result)
	})

	t.Run("Should return a conflict error when record name already exists", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		mockedProductRecord := mockedProductRecordTemplate
		repository.On("Exists", mockedProductRecord.ProductID, mockedProductRecord.LastUpdateDate).Return(true)
		result, err := service.Create(mockedProductRecord)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})

	t.Run("Should return a conflict error when record id not exists", func(t *testing.T) {
		service, repository, productRepository := CreateService(t)

		mockedProductRecord := mockedProductRecordTemplate
		mockedProduct := mockedProductTemplate
		var productRepositoryGetResult *domain.Product

		repository.On("Exists", mockedProductRecord.ProductID, mockedProductRecord.LastUpdateDate).Return(false)
		productRepository.On("Get", mockedProduct.ID).Return(productRepositoryGetResult)
		result, err := service.Create(mockedProductRecord)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})
}

func CreateService(t *testing.T) (record.Service, *mocks.Repository, *productMocks.Repository) {
	t.Helper()
	repository := new(mocks.Repository)
	productRepository := new(productMocks.Repository)
	service := record.NewService(repository, productRepository)
	return service, repository, productRepository
}
