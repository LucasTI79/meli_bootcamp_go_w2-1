package product_batch_test

import (
	"testing"
	"time"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	product_mocks "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_batch"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_batch/mocks"
	section_mocks "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/section/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/stretchr/testify/assert"
)

var (
	productBatch = domain.ProductBatch{
		ID:                 1,
		BatchNumber:        1,
		CurrentQuantity:    1,
		CurrentTemperature: 2,
		DueDate:            time.Date(2021, 01, 01, 10, 10, 10, 10, time.UTC),
		InitialQuantity:    10,
		ManufacturingDate:  time.Date(2021, 01, 01, 10, 10, 10, 10, time.UTC),
		ManufacturingHour:  10,
		MinimumTemperature: 0,
		ProductID:          1,
		SectionID:          1,
	}
)

func TestServiceCreate(t *testing.T) {
	t.Run("Should return a created product batches", func(t *testing.T) {
		service, repository, productRepository, sectionRepository := CreateService(t)

		id := 1
		repository.On("Save", productBatch).Return(id)
		repository.On("Get", id).Return(&productBatch)
		repository.On("Exists", productBatch.BatchNumber).Return(false)
		productRepository.On("Get", productBatch.ProductID).Return(&domain.Product{})
		sectionRepository.On("Get", productBatch.SectionID).Return(&domain.Section{})
		result, err := service.Create(productBatch)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, productBatch, *result)
	})
	t.Run("Should return a conflict error when product not found", func(t *testing.T) {
		service, repository, productRepository, _ := CreateService(t)

		mockedProductBatch := productBatch
		var productGetResult *domain.Product

		repository.On("Exists", mockedProductBatch.BatchNumber).Return(false)
		productRepository.On("Get", mockedProductBatch.ProductID).Return(productGetResult)
		result, err := service.Create(mockedProductBatch)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})
	t.Run("Should return a conflict error when section not found", func(t *testing.T) {
		service, repository, productRepository, sectionRepository := CreateService(t)

		mockedProductBatch := productBatch
		var sectionGetResult *domain.Section

		repository.On("Exists", mockedProductBatch.BatchNumber).Return(false)
		productRepository.On("Get", mockedProductBatch.ProductID).Return(&domain.Product{})
		sectionRepository.On("Get", mockedProductBatch.SectionID).Return(sectionGetResult)
		result, err := service.Create(mockedProductBatch)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})
	t.Run("Should return a conflict error when batch number already exists", func(t *testing.T) {
		service, repository, _, _ := CreateService(t)

		mockedProductBatch := productBatch

		repository.On("Exists", mockedProductBatch.BatchNumber).Return(true)
		result, err := service.Create(mockedProductBatch)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))

	})
}

func CreateService(t *testing.T) (product_batch.Service, *mocks.Repository, *product_mocks.Repository, *section_mocks.Repository) {
	t.Helper()
	repository := new(mocks.Repository)
	productRepository := new(product_mocks.Repository)
	sectionRepository := new(section_mocks.Repository)
	service := product_batch.NewService(repository, productRepository, sectionRepository)

	return service, repository, productRepository, sectionRepository
}
