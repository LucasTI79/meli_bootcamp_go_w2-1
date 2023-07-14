package carrier_test

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/carrier"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/carrier/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	localityMocks "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/locality/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/stretchr/testify/assert"
)

var (
	c = domain.Carrier{
		ID:          1,
		CID:         "123456",
		CompanyName: "CompaniName",
		Address:     "Address",
		Telephone:   "+554312343212",
		LocalityID:  1,
	}

	mockedLocalityTemplate = domain.Locality{
		ID:           1,
		LocalityName: "Locality",
		ProvinceID:   1,
	}
)

func TestServiceCreate(t *testing.T) {
	t.Run("Should return a created carrier", func(t *testing.T) {
		service, repository, localityRepo := CreateService(t)

		id := 1
		repository.On("Exists", c.CID).Return(false)
		localityRepo.On("Get", c.LocalityID).Return(&mockedLocalityTemplate)
		repository.On("Save", c).Return(id)
		repository.On("Get", id).Return(&c)
		result, err := service.Create(c)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, c, *result)
	})

	t.Run("Should return a conflict error", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		repository.On("Exists", c.CID).Return(true)
		result, err := service.Create(c)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})

	t.Run("should return dependent resource not found error", func(t *testing.T) {
		service, repository, localityRepo := CreateService(t)

		var emptyLocality *domain.Locality
		id := 1
		repository.On("Exists", c.CID).Return(false)
		localityRepo.On("Get", c.LocalityID).Return(emptyLocality)
		repository.On("Save", c).Return(id)
		repository.On("Get", id).Return(&c)
		result, err := service.Create(c)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})
}

func CreateService(t *testing.T) (carrier.Service, *mocks.Repository, *localityMocks.Repository) {
	t.Helper()
	repository := new(mocks.Repository)
	localityRepo := new(localityMocks.Repository)
	service := carrier.NewService(repository, localityRepo)

	return service, repository, localityRepo
}
