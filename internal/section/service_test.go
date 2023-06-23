package section_test

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/section"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/section/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/stretchr/testify/assert"
)

var (
	mockedSection = domain.Section{
		ID: 1,
		SectionNumber: 1,      
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity: 1,
		MinimumCapacity: 1,
		MaximumCapacity: 1,
		WarehouseID: 1,
		ProductTypeID: 1,      
	}
)

func TestServiceCreate(t *testing.T) {
	t.Run("Should return a created section", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		repository.On("Save", mockedSection).Return(id)
		repository.On("Get", id).Return(&mockedSection)
		repository.On("Exists", mockedSection.SectionNumber).Return(false)
		result, err := service.Create(context.TODO(), mockedSection)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, mockedSection, *result)
	})

	t.Run("Should return a conflict error", func(t *testing.T) {
		service, repository := CreateService(t)

		repository.On("Exists", mockedSection.SectionNumber).Return(true)
		result, err := service.Create(context.TODO(), mockedSection)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})
}

func TestServiceGet(t *testing.T) {
	t.Run("Should return a list o sections", func(t *testing.T) {
		service, repository := CreateService(t)

		expected := []domain.Section{mockedSection}

		repository.On("GetAll").Return(expected)
		result := service.GetAll(context.TODO())

		assert.NotEmpty(t, result)
		assert.Equal(t, len(result), 1)
		assert.Equal(t, result[0], mockedSection)
	})

	t.Run("Should return a section by specified id", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1

		repository.On("Get", id).Return(&mockedSection)
		result, err := service.Get(context.TODO(), id)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, *result, mockedSection)
	})

	t.Run("Should return a not found error", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		var respositoryResult *domain.Section

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
		sectionNumber := 123
		UpdateSection := domain.UpdateSection{
			ID:          &id,
			SectionNumber: &sectionNumber,
		}

		var respositoryResult *domain.Section

		repository.On("Get", id).Return(respositoryResult)
		result, err := service.Update(context.TODO(), id, UpdateSection)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})

	t.Run("Should return a conflict error", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		sectionNumber := 456
		UpdateSection := domain.UpdateSection{
			ID:          &id,
			SectionNumber: &sectionNumber,
		}

		repository.On("Get", id).Return(&mockedSection)
		repository.On("Exists", sectionNumber).Return(true)
		result, err := service.Update(context.TODO(), id, UpdateSection)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})

	t.Run("Should return an updated section", func(t *testing.T) {
		service, repository := CreateService(t)

		id := 1
		sectionNumber := 123
		currentTemperature := 2
		UpdateSection := domain.UpdateSection{
			ID:          &id,
			CurrentTemperature: &currentTemperature,
			SectionNumber: &sectionNumber,
		}
		updatedSection := mockedSection
		updatedSection.Overlap(UpdateSection)

		repository.On("Get", id).Return(&mockedSection)
		repository.On("Exists", sectionNumber).Return(false)
		repository.On("Update", updatedSection)
		repository.On("Get", id).Return(&updatedSection)
		result, err := service.Update(context.TODO(), id, UpdateSection)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, currentTemperature, result.CurrentTemperature)
	})
}



func CreateService(t *testing.T) (section.Service, *mocks.Repository) {
	t.Helper()
	repository := new(mocks.Repository)
	service := section.NewService(repository)
	return service, repository
}
