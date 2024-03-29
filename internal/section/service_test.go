package section_test

import (
	"testing"

	product_type_mocks "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_type/mocks"
	warehouse_mocks "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse/mocks"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/section/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/stretchr/testify/assert"
)

var (
	mockedSection = domain.Section{
		ID:                 1,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseID:        1,
		ProductTypeID:      1,
	}

	mockedProductsBySection = domain.ProductsBySectionReport{
		SectionID:     1,
		SectionNumber: 1,
		ProductsCount: 1,
	}
)

func TestServiceCreate(t *testing.T) {
	t.Run("Should return a created section", func(t *testing.T) {
		service, repository, warehouseRepository, productTypeRepository := CreateService(t)

		id := 1
		repository.On("Save", mockedSection).Return(id)
		repository.On("Get", id).Return(&mockedSection)
		repository.On("Exists", mockedSection.SectionNumber).Return(false)
		productTypeRepository.On("Get", mockedSection.ProductTypeID).Return(&domain.ProductType{})
		warehouseRepository.On("Get", mockedSection.WarehouseID).Return(&domain.Warehouse{})
		result, err := service.Create(mockedSection)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, mockedSection, *result)
	})

	t.Run("Should return a conflict error", func(t *testing.T) {
		service, repository, _, _ := CreateService(t)

		repository.On("Exists", mockedSection.SectionNumber).Return(true)
		result, err := service.Create(mockedSection)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})

	t.Run("Should return a product type dependent resource not found error", func(t *testing.T) {
		service, repository, _, productTypeRepository := CreateService(t)

		var productTypeResult *domain.ProductType

		repository.On("Exists", mockedSection.SectionNumber).Return(false)
		productTypeRepository.On("Get", mockedSection.ProductTypeID).Return(productTypeResult)
		result, err := service.Create(mockedSection)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})

	t.Run("Should return a warehouse dependent resource not found error", func(t *testing.T) {
		service, repository, warehouseRepository, productTypeRepository := CreateService(t)

		var warehouseResult *domain.Warehouse

		repository.On("Exists", mockedSection.SectionNumber).Return(false)
		productTypeRepository.On("Get", mockedSection.ProductTypeID).Return(&domain.ProductType{})
		warehouseRepository.On("Get", mockedSection.WarehouseID).Return(warehouseResult)
		result, err := service.Create(mockedSection)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})
}

func TestServiceGet(t *testing.T) {
	t.Run("Should return a list of sections", func(t *testing.T) {
		service, repository, _, _ := CreateService(t)

		expected := []domain.Section{mockedSection}

		repository.On("GetAll").Return(expected)
		result := service.GetAll()

		assert.NotEmpty(t, result)
		assert.Equal(t, len(result), 1)
		assert.Equal(t, result[0], mockedSection)
	})

	t.Run("Should return a section by specified id", func(t *testing.T) {
		service, repository, _, _ := CreateService(t)

		id := 1

		repository.On("Get", id).Return(&mockedSection)
		result, err := service.Get(id)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, *result, mockedSection)
	})

	t.Run("Should return a not found error", func(t *testing.T) {
		service, repository, _, _ := CreateService(t)

		id := 1
		var respositoryResult *domain.Section

		repository.On("Get", id).Return(respositoryResult)
		result, err := service.Get(id)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})
}

func TestServiceUpdate(t *testing.T) {
	t.Run("Should return a not found error", func(t *testing.T) {
		service, repository, _, _ := CreateService(t)

		id := 2
		sectionNumber := 123
		UpdateSection := domain.UpdateSection{
			ID:            &id,
			SectionNumber: &sectionNumber,
		}

		var respositoryResult *domain.Section

		repository.On("Get", id).Return(respositoryResult)
		result, err := service.Update(id, UpdateSection)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})

	t.Run("Should return a conflict error", func(t *testing.T) {
		service, repository, _, _ := CreateService(t)

		id := 1
		sectionNumber := 456
		UpdateSection := domain.UpdateSection{
			ID:            &id,
			SectionNumber: &sectionNumber,
		}

		repository.On("Get", id).Return(&mockedSection)
		repository.On("Exists", sectionNumber).Return(true)
		result, err := service.Update(id, UpdateSection)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})

	t.Run("Should return a product type dependent resource not found error", func(t *testing.T) {
		service, repository, _, productTypeRepository := CreateService(t)

		id := 1
		sectionNumber := 456
		UpdateSection := domain.UpdateSection{
			ID:            &id,
			SectionNumber: &sectionNumber,
		}
		var productTypeResult *domain.ProductType

		repository.On("Get", id).Return(&mockedSection)
		repository.On("Exists", sectionNumber).Return(false)
		productTypeRepository.On("Get", mockedSection.ProductTypeID).Return(productTypeResult)
		result, err := service.Update(id, UpdateSection)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})

	t.Run("Should return a warehouse dependent resource not found error", func(t *testing.T) {
		service, repository, warehouseRepository, productTypeRepository := CreateService(t)

		id := 1
		sectionNumber := 456
		UpdateSection := domain.UpdateSection{
			ID:            &id,
			SectionNumber: &sectionNumber,
		}
		var warehouseResult *domain.Warehouse

		repository.On("Get", id).Return(&mockedSection)
		repository.On("Exists", sectionNumber).Return(false)
		productTypeRepository.On("Get", mockedSection.ProductTypeID).Return(&domain.ProductType{})
		warehouseRepository.On("Get", mockedSection.WarehouseID).Return(warehouseResult)
		result, err := service.Update(id, UpdateSection)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})

	t.Run("Should return an updated section", func(t *testing.T) {
		service, repository, warehouseRepository, productTypeRepository := CreateService(t)

		id := 1
		sectionNumber := 123
		currentTemperature := float32(2)
		UpdateSection := domain.UpdateSection{
			ID:                 &id,
			CurrentTemperature: &currentTemperature,
			SectionNumber:      &sectionNumber,
		}
		updatedSection := mockedSection
		updatedSection.Overlap(UpdateSection)

		repository.On("Get", id).Return(&mockedSection)
		repository.On("Exists", sectionNumber).Return(false)
		productTypeRepository.On("Get", mockedSection.ProductTypeID).Return(&domain.ProductType{})
		warehouseRepository.On("Get", mockedSection.WarehouseID).Return(&domain.Warehouse{})
		repository.On("Update", updatedSection)
		repository.On("Get", id).Return(&updatedSection)
		result, err := service.Update(id, UpdateSection)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, currentTemperature, result.CurrentTemperature)
	})
}

func TestServiceDelete(t *testing.T) {
	t.Run("Should return not found error", func(t *testing.T) {
		service, repository, _, _ := CreateService(t)

		id := 1
		var respositoryResult *domain.Section

		repository.On("Get", id).Return(respositoryResult)
		err := service.Delete(id)

		assert.Error(t, err)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})

	t.Run("Should delete a section with success", func(t *testing.T) {
		service, repository, _, _ := CreateService(t)

		id := 1

		repository.On("Get", id).Return(&mockedSection)
		repository.On("Delete", id)

		err := service.Delete(id)

		assert.NoError(t, err)
	})
}

func TestCountProductsByAllSections(t *testing.T) {
	t.Run("Should return a list of sections", func(t *testing.T) {
		service, repository, _, _ := CreateService(t)

		expected := []domain.ProductsBySectionReport{mockedProductsBySection}

		repository.On("CountProductsByAllSections").Return(expected)
		result := service.CountProductsByAllSections()

		assert.NotEmpty(t, result)
		assert.True(t, len(result) == 1)
		assert.Equal(t, result[0], mockedProductsBySection)
	})
}

// ok
func TestCountProductsBySection(t *testing.T) {
	t.Run("Should return amount of products by section by a specified id", func(t *testing.T) {
		service, repository, _, _ := CreateService(t)

		mockedProductsBySec := mockedProductsBySection
		id := 1

		repository.On("CountProductsBySection", id).Return(&mockedProductsBySec)
		result, err := service.CountProductsBySection(id)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, *result, mockedProductsBySec)
	})
	t.Run("Should return a not found error", func(t *testing.T) {
		service, repository, _, _ := CreateService(t)

		id := 1
		var respositoryResult *domain.ProductsBySectionReport

		repository.On("CountProductsBySection", id).Return(respositoryResult)
		result, err := service.CountProductsBySection(id)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})
}

func CreateService(t *testing.T) (section.Service, *mocks.Repository, *warehouse_mocks.Repository, *product_type_mocks.Repository) {
	t.Helper()
	repository := new(mocks.Repository)
	productTypeRepository := new(product_type_mocks.Repository)
	warehouseRepository := new(warehouse_mocks.Repository)
	service := section.NewService(repository, warehouseRepository, productTypeRepository)
	return service, repository, warehouseRepository, productTypeRepository
}
