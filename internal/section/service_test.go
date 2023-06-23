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
	s = domain.Section{
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
		repository.On("Save", s).Return(id)
		repository.On("Get", id).Return(&s)
		repository.On("Exists", s.SectionNumber).Return(false)
		result, err := service.Create(context.TODO(), s)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, s, *result)
	})

	t.Run("Should return a conflict error", func(t *testing.T) {
		service, repository := CreateService(t)

		repository.On("Exists", s.SectionNumber).Return(true)
		result, err := service.Create(context.TODO(), s)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})
}

func CreateService(t *testing.T) (section.Service, *mocks.Repository) {
	t.Helper()
	repository := new(mocks.Repository)
	service := section.NewService(repository)
	return service, repository
}
