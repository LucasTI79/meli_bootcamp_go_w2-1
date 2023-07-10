package product_batches_test

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/stretchr/testify/assert"
)

func TestServiceCreate(t *testing.T) {
	t.Run("should create a product batch", func(t *testing.T) {
		service, repository := CreateService(t)

		productBatch := domain.ProductBatches{
			BatchNumber:       1,
			ProductID:         1,
			DueDate:           "2021-01-01",
			ManufacturingDate: "2021-01-01",
		}
		repository.On("Create", productBatch).Return(productBatch, nil)
		repository.On("GetByBatchNumber", productBatch.BatchNumber).Return(domain.ProductBatches{}, apperr.ErrConflict)
		repository.On("GetByProductID", productBatch.ProductID).Return(domain.ProductBatches{}, apperr.ErrConflict)

		result, err := service.Create(productBatch)

		assert.NoError(t, err)
		assert.Equal(t, productBatch, result)
		assert.NotNil(t, result)
	})

}
