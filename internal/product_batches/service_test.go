package product_batches_test

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	product_mocks "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_batches"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_batches/mocks"
	section_mocks "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/section/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
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
		ManufacturingHour:  10,
		MinimumTemperature: 0,
		ProductID:          1,
		SectionID:          1,
	}
)

func TestServiceCreate(t *testing.T) {
	t.Run("Should return a created product batches", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		repository.On("Save", productBatch).Return(id)
		repository.On("Get", id).Return(&productBatch)
		repository.On("Exists", productBatch.BatchNumber).Return(false)
		result, err := service.Create(productBatch)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, productBatch, result)
	})
	t.Run("Should return a conflict error when product batches number already exists", func(t *testing.T) {
		service, repository := CreateService(t)

		repository.On("Exists", pb.BatchNumber).Return(true)
		result, err := service.Create(pb)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})
}

func CreateService(t *testing.T) (product_batches.Service, *mocks.Repository) {
	t.Helper()
	repository := new(mocks.Repository)
	productRepository := new(product_mocks.Repository)
	sectionRepository := new(section_mocks.Repository)
	service := product_batches.NewService(repository, productRepository, sectionRepository)

	return service, repository
}
