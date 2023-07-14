package warehouse_test

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	localityMock "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/locality/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/stretchr/testify/assert"
)

var (
	w = domain.Warehouse{
		ID:                 1,
		Address:            "Address",
		Telephone:          "12345",
		WarehouseCode:      "123",
		MinimumCapacity:    1,
		MinimumTemperature: 1,
		LocalityID: 1,
	}

	l = domain.Locality{
		ID: 1,
	}
)

func TestServiceCreate(t *testing.T) {
	t.Run("Should return a created warehous", func(t *testing.T) {
		service, repository, localityRepository := CreateService(t)

		id := 1
		repository.On("Exists", w.WarehouseCode).Return(false)
		localityRepository.On("Get", w.LocalityID).Return(&l)
		repository.On("Save", w).Return(id)
		repository.On("Get", id).Return(&w)

		result, err := service.Create(w)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, w, *result)
	})

	t.Run("Should return a conflict error if warehouse already exists", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		repository.On("Exists", w.WarehouseCode).Return(true)
		result, err := service.Create(w)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})

	t.Run("should return a conflict error if locality id doesn't exist", func(t *testing.T) {
		service, repository, localityRepository := CreateService(t)

		var emptyLocality *domain.Locality

		repository.On("Exists", w.WarehouseCode).Return(false)
		localityRepository.On("Get", w.LocalityID).Return(emptyLocality)

		result, err := service.Create(w)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})
}

func TestServiceGet(t *testing.T) {
	t.Run("Should return a list of warehouses", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		expected := []domain.Warehouse{w}

		repository.On("GetAll").Return(expected)
		result := service.GetAll()

		assert.NotEmpty(t, result)
		assert.Equal(t, len(result), 1)
		assert.Equal(t, result[0], w)
	})

	t.Run("Should return a warehouse by specified id", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		id := 1

		repository.On("Get", id).Return(&w)
		result, err := service.Get(id)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, *result, w)
	})

	t.Run("Should return a not found error", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		id := 1
		var respositoryResult *domain.Warehouse

		repository.On("Get", id).Return(respositoryResult)
		result, err := service.Get(id)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})
}

func TestServiceUpdate(t *testing.T) {
	t.Run("Should return a not found error", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		id := 2
		warehouseCode := "153"
		updateWarehouse := domain.UpdateWarehouse{
			ID:            &id,
			WarehouseCode: &warehouseCode,
		}
		var respositoryResult *domain.Warehouse

		repository.On("Get", id).Return(respositoryResult)
		result, err := service.Update(id, updateWarehouse)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})

	t.Run("Should return a conflict error", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		id := 1
		warehouseCode := "496"
		updateWarehouse := domain.UpdateWarehouse{
			ID:            &id,
			WarehouseCode: &warehouseCode,
		}

		repository.On("Get", id).Return(&w)
		repository.On("Exists", warehouseCode).Return(true)
		result, err := service.Update(id, updateWarehouse)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})

	t.Run("Should return an updated warehouse", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		id := 1
		warehouseCode := "123"
		address := "Address 3"
		updateWarehouse := domain.UpdateWarehouse{
			ID:            &id,
			Address:       &address,
			WarehouseCode: &warehouseCode,
		}
		updatedWarehouse := w
		updatedWarehouse.Overlap(updateWarehouse)

		repository.On("Get", id).Return(&w)
		repository.On("Exists", warehouseCode).Return(true)
		repository.On("Update", updatedWarehouse)
		repository.On("Get", id).Return(&updatedWarehouse)
		result, err := service.Update(id, updateWarehouse)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, address, result.Address)
	})
}

func TestServiceDelete(t *testing.T) {
	t.Run("Should return not found error", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		id := 1
		var respositoryResult *domain.Warehouse

		repository.On("Get", id).Return(respositoryResult)
		err := service.Delete(id)

		assert.Error(t, err)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})

	t.Run("Should delete a warehouse with success", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		id := 1

		repository.On("Get", 1).Return(&w)
		repository.On("Delete", id)
		err := service.Delete(id)

		assert.NoError(t, err)
	})
}

func CreateService(t *testing.T) (warehouse.Service, *mocks.Repository, *localityMock.Repository) {
	t.Helper()
	repository := new(mocks.Repository)
	localityRepository := new(localityMock.Repository)
	service := warehouse.NewService(repository, localityRepository)
	return service, repository, localityRepository
}
