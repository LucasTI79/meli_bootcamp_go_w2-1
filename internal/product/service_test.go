package product_test

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/stretchr/testify/assert"
)

var (
	p = domain.Product{
		ID:             1,
		Description:    "Description",
		ExpirationRate: 1,
		FreezingRate:   1,
		Height:         1,
		Length:         1,
		Netweight:      1,
		ProductCode:    "123",
		RecomFreezTemp: 1,
		Width:          1,
		ProductTypeID:  1,
		SellerID:       1,
	}
)

func TestServiceCreate(t *testing.T) {
	t.Run("Should return a created product", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		repository.On("Save", p).Return(id)
		repository.On("Get", id).Return(&p)
		repository.On("Exists", p.ProductCode).Return(false)
		result, err := service.Create(p)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, p, *result)
	})

	t.Run("Should return a conflict error", func(t *testing.T) {
		service, repository := CreateService(t)

		repository.On("Exists", p.ProductCode).Return(true)
		result, err := service.Create(p)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})
}

func TestServiceGet(t *testing.T) {
	t.Run("Should return a list of products", func(t *testing.T) {
		service, repository := CreateService(t)

		expected := []domain.Product{p}

		repository.On("GetAll").Return(expected)
		result := service.GetAll()

		assert.NotEmpty(t, result)
		assert.Equal(t, len(result), 1)
		assert.Equal(t, result[0], p)
	})

	t.Run("Should return a product by specified id", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1

		repository.On("Get", id).Return(&p)
		result, err := service.Get(id)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, *result, p)
	})

	t.Run("Should return a not found error", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		var respositoryResult *domain.Product

		repository.On("Get", id).Return(respositoryResult)
		result, err := service.Get(id)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})
}

func TestServiceUpdate(t *testing.T) {
	t.Run("Should return a not found error", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 2
		productCode := "123"
		updateProduct := domain.UpdateProduct{
			ID:          &id,
			ProductCode: &productCode,
		}

		var respositoryResult *domain.Product

		repository.On("Get", id).Return(respositoryResult)
		result, err := service.Update(id, updateProduct)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})

	t.Run("Should return a conflict error", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		productCode := "456"
		updateProduct := domain.UpdateProduct{
			ID:          &id,
			ProductCode: &productCode,
		}

		repository.On("Get", id).Return(&p)
		repository.On("Exists", productCode).Return(true)
		result, err := service.Update(id, updateProduct)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})

	t.Run("Should return an updated product", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		productCode := "123"
		description := "Description 2"
		updateProduct := domain.UpdateProduct{
			ID:          &id,
			Description: &description,
			ProductCode: &productCode,
		}
		updatedProduct := p
		updatedProduct.Overlap(updateProduct)

		repository.On("Get", id).Return(&p)
		repository.On("Exists", productCode).Return(true)
		repository.On("Update", updatedProduct)
		repository.On("Get", id).Return(&updatedProduct)
		result, err := service.Update(id, updateProduct)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, description, result.Description)
	})
}

func TestServiceDelete(t *testing.T) {
	t.Run("Should return not found error", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		var respositoryResult *domain.Product

		repository.On("Get", id).Return(respositoryResult)
		err := service.Delete(id)

		assert.Error(t, err)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})

	t.Run("Should delete a product with success", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1

		repository.On("Get", id).Return(&p)
		repository.On("Delete", id)
		err := service.Delete(id)

		assert.NoError(t, err)
	})
}

func TestServiceCountRecordsByAllProducts(t *testing.T) {
	t.Run("Should return records count report of all product", func(t *testing.T) {
		service, repository := CreateService(t)

		mockedRecordsByProductReport := domain.RecordsByProductReport{
			ProductID:    1,
			Description:  "Description",
			RecordsCount: 1,
		}
		mockedRecordsByProductsReport := []domain.RecordsByProductReport{mockedRecordsByProductReport}

		repository.On("CountRecordsByAllProducts").Return(mockedRecordsByProductsReport)

		result := service.CountRecordsByAllProducts()

		assert.Equal(t, 1, len(result))
		assert.Equal(t, result[0], mockedRecordsByProductReport)
	})
}

func TestServiceCountRecordsByProduct(t *testing.T) {
	t.Run("Should return records count report by specified product id", func(t *testing.T) {
		service, repository := CreateService(t)

		recordId := 1
		mockedProduct := mockedProductTemplate
		mockedRecordsByProductReport := domain.RecordsByProductReport{
			ProductID:    1,
			Description:  "Description",
			RecordsCount: 1,
		}

		repository.On("Get", recordId).Return(&mockedProduct)
		repository.On("CountRecordsByProduct", recordId).Return(&mockedRecordsByProductReport)

		result, err := service.CountRecordsByProduct(recordId)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, *result, mockedRecordsByProductReport)
	})

	t.Run("Should return not found when product id not exists", func(t *testing.T) {
		service, repository := CreateService(t)

		recordId := 1

		var productRepositoryGetResult *domain.Product
		repository.On("Get", recordId).Return(productRepositoryGetResult)

		result, err := service.CountRecordsByProduct(recordId)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})
}

func CreateService(t *testing.T) (product.Service, *mocks.Repository) {
	t.Helper()
	repository := new(mocks.Repository)
	service := product.NewService(repository)
	return service, repository
}
