package product_batches_test

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_batches"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_batches/mocks"
	"github.com/stretchr/testify/assert"
)

var (
	productBatch = domain.ProductBatches{
		ID:                 1,
		BatchNumber:        1,
		CurrentQuantity:    1,
		CurrentTemperature: 2,
		DueDate:            "2021-01-01",
		InitialQuantity:    10,
		ManufacturingDate:  "2021-01-01",
		ManufacturingHour:  "10:00",
		MinimumTemperature: 0,
		ProductID:          1,
		SectionID:          1,
	}
)

var (
	productBySection = domain.ProductsBySectionReport{
		SectionID:     1,
		SectionNumber: 1,
		ProductsCount: 1,
	}
)

func TestServiceCreate(t *testing.T) {
	t.Run("Should return a created product batches", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		repository.On("Exists", productBatch.BatchNumber).Return(false)
		repository.On("Save", productBatch).Return(id)
		repository.On("Get", id).Return(&productBatch)

		result, err := service.Create(productBatch)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, productBatch, *result)
	})
	t.Run("Should return a conflict error when product batches number already exists", func(t *testing.T) {
		service, repository := CreateService(t)

		repository.On("Exists", pb.BatchNumber).Return(true)
		result, err := service.Create(pb)

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestServiceExists(t *testing.T) {
	t.Run("Should return true when product batches exists", func(t *testing.T) {
		service, repository := CreateService(t)

		repository.On("Exists", pb.BatchNumber).Return(true)
		result, _ := service.Exists(pb.BatchNumber)

		assert.True(t, result)
	})

	t.Run("Should return false when product batches does not exists", func(t *testing.T) {
		service, repository := CreateService(t)

		repository.On("Exists", pb.BatchNumber).Return(false)
		result, _ := service.Exists(pb.BatchNumber)

		assert.False(t, result)
	})
}

func TestServiceGet(t *testing.T) {
	t.Run("Should return all product batches", func(t *testing.T) {
		service, repository := CreateService(t)

		repository.On("CountProductsByAllSections").Return([]domain.ProductsBySectionReport{productBySection}, nil)
		result, err := service.CountProductsByAllSections()

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, []domain.ProductsBySectionReport{productBySection}, result)
	})
	t.Run("Should return all products counts by specified id", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		repository.On("CountProductsBySection", id).Return([]domain.ProductsBySectionReport{productBySection}, nil)
		result, err := service.CountProductsBySection(id)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, []domain.ProductsBySectionReport{productBySection}, result)
	})
}

func CreateService(t *testing.T) (product_batches.Service, *mocks.Repository) {
	t.Helper()
	repository := new(mocks.Repository)
	service := product_batches.NewService(repository)

	return service, repository
}
