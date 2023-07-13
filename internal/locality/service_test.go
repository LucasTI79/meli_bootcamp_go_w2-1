package locality_test

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/locality"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/locality/mocks"
	provinceMocks "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/province/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/stretchr/testify/assert"
)

var (
	mockedProvinceTemplate = domain.Province{
		ID: 1,
	}
)

func TestServiceCreate(t *testing.T) {
	t.Run("Should return a created locality", func(t *testing.T) {
		service, repository, provinceRepository := CreateService(t)

		mockedLocality := mockedLocalityTemplate
		mockedProvince := mockedProvinceTemplate
		id := 1
		localityId := 1

		repository.On("Exists", mockedLocality.LocalityName).Return(false)
		provinceRepository.On("Get", localityId).Return(&mockedProvince)
		repository.On("Save", mockedLocality).Return(id)
		repository.On("Get", id).Return(&mockedLocality)
		result, err := service.Create(mockedLocality)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, mockedLocality, *result)
	})

	t.Run("Should return a conflict error when locality name already exists", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		mockedLocality := mockedLocalityTemplate
		repository.On("Exists", mockedLocality.LocalityName).Return(true)
		result, err := service.Create(mockedLocality)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})

	t.Run("Should return a conflict error when locality id not exists", func(t *testing.T) {
		service, repository, provinceRepository := CreateService(t)

		mockedLocality := mockedLocalityTemplate
		mockedProvince := mockedProvinceTemplate
		var provinceRepositoryGetResult *domain.Province

		repository.On("Exists", mockedLocality.LocalityName).Return(false)
		provinceRepository.On("Get", mockedProvince.ID).Return(provinceRepositoryGetResult)
		result, err := service.Create(mockedLocality)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})
}

func TestServiceCountSellersByAllLocalities(t *testing.T) {
	t.Run("Should return sellers count report of all localities", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		mockedSellersByLocalityReport := domain.SellersByLocalityReport{
			ID:           1,
			LocalityName: "Locality Name",
			SellersCount: 1,
		}
		mockedSellersByLocalitiesReport := []domain.SellersByLocalityReport{mockedSellersByLocalityReport}

		repository.On("CountSellersByAllLocalities").Return(mockedSellersByLocalitiesReport)

		result := service.CountSellersByAllLocalities()

		assert.Equal(t, 1, len(result))
		assert.Equal(t, result[0], mockedSellersByLocalityReport)
	})
}

func TestServiceCountSellersByLocality(t *testing.T) {
	t.Run("Should return sellers count report by specified locality id", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		localityId := 1
		mockedLocality := mockedLocalityTemplate
		mockedSellersByLocalityReport := domain.SellersByLocalityReport{
			ID:           1,
			LocalityName: "Locality Name",
			SellersCount: 1,
		}

		repository.On("Get", localityId).Return(&mockedLocality)
		repository.On("CountSellersByLocality", localityId).Return(&mockedSellersByLocalityReport)

		result, err := service.CountSellersByLocality(localityId)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, *result, mockedSellersByLocalityReport)
	})

	t.Run("Should return not found when locality id not exists", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		localityId := 1

		var localityRepositoryGetResult *domain.Locality
		repository.On("Get", localityId).Return(localityRepositoryGetResult)

		result, err := service.CountSellersByLocality(localityId)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})
}

func TestServiceCountCarriersByAllLocalities(t *testing.T) {
	t.Run("Should return carriers count of all localities", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		mockedCarriersByLocalityReport := domain.CarriersByLocalityReport{
			ID:            1,
			LocalityName:  "Locality Name",
			CarriersCount: 1,
		}
		mockedSellersByLocalitiesReport := []domain.CarriersByLocalityReport{mockedCarriersByLocalityReport}

		repository.On("CountCarriersByAllLocalities").Return(mockedSellersByLocalitiesReport)

		result := service.CountCarriersByAllLocalities()

		assert.Equal(t, 1, len(result))
		assert.Equal(t, result[0], mockedCarriersByLocalityReport)
	})
}

func TestServiceCountCarriersByLocality(t *testing.T) {
	t.Run("Should return carriers count by specified locality id", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		localityId := 1
		mockedLocality := mockedLocalityTemplate
		mockedSellersByLocalityReport := domain.CarriersByLocalityReport{
			ID:            1,
			LocalityName:  "Locality Name",
			CarriersCount: 1,
		}

		repository.On("Get", localityId).Return(&mockedLocality)
		repository.On("CountCarriersByLocality", localityId).Return(&mockedSellersByLocalityReport)

		result, err := service.CountCarriersByLocality(localityId)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, *result, mockedSellersByLocalityReport)
	})

	t.Run("Should return not found when locality id not exists", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		localityId := 1

		var localityRepositoryGetResult *domain.Locality
		repository.On("Get", localityId).Return(localityRepositoryGetResult)

		result, err := service.CountCarriersByLocality(localityId)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})
}

func CreateService(t *testing.T) (locality.Service, *mocks.Repository, *provinceMocks.Repository) {
	t.Helper()
	repository := new(mocks.Repository)
	provinceRepository := new(provinceMocks.Repository)
	service := locality.NewService(repository, provinceRepository)
	return service, repository, provinceRepository
}
