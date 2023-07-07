package seller_test

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	localityMocks "github.com/extmatperez/meli_bootcamp_go_w2-1/internal/locality/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/seller"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/seller/mocks"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
	"github.com/stretchr/testify/assert"
)

var (
	mockedLocalityTemplate = domain.Locality{
		ID:           1,
		LocalityName: "Locality",
		ProvinceID:   1,
	}

	mockedSellerTemplate = domain.Seller{
		ID:          1,
		CID:         123,
		CompanyName: "Company Name",
		Address:     "Address",
		Telephone:   "Telephone",
		LocalityID:  mockedLocalityTemplate.ID,
	}
)

func TestServiceCreate(t *testing.T) {
	t.Run("Should return a created seller", func(t *testing.T) {
		service, repository, localityRepository := CreateService(t)

		mockedSeller := mockedSellerTemplate
		mockedLocality := mockedLocalityTemplate
		id := 1
		localityId := 1

		repository.On("Exists", mockedSeller.CID).Return(false)
		localityRepository.On("Get", localityId).Return(&mockedLocality)
		repository.On("Save", mockedSeller).Return(id)
		repository.On("Get", id).Return(&mockedSeller)
		result, err := service.Create(mockedSeller)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, mockedSeller, *result)
	})

	t.Run("Should return a conflict error when cid already exists", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		mockedSeller := mockedSellerTemplate
		repository.On("Exists", mockedSeller.CID).Return(true)
		result, err := service.Create(mockedSeller)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})

	t.Run("Should return a conflict error when locality id not exists", func(t *testing.T) {
		service, repository, localityRepository := CreateService(t)

		mockedSeller := mockedSellerTemplate
		mockedLocality := mockedLocalityTemplate
		var localityRepositoryGetResult *domain.Locality

		repository.On("Exists", mockedSeller.CID).Return(false)
		localityRepository.On("Get", mockedLocality.ID).Return(localityRepositoryGetResult)
		result, err := service.Create(mockedSeller)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})
}

func TestServiceGet(t *testing.T) {
	t.Run("Should return a list of sellers", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		mockedSeller := mockedSellerTemplate
		expected := []domain.Seller{mockedSeller}

		repository.On("GetAll").Return(expected)
		result := service.GetAll()

		assert.NotEmpty(t, result)
		assert.Equal(t, len(result), 1)
		assert.Equal(t, result[0], mockedSeller)
	})

	t.Run("Should return a seller by specified id", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		mockedSeller := mockedSellerTemplate
		id := 1

		repository.On("Get", id).Return(&mockedSeller)
		result, err := service.Get(id)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, *result, mockedSeller)
	})

	t.Run("Should return a not found error", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		id := 1
		var respositoryResult *domain.Seller

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
		cid := 123
		updateSeller := domain.UpdateSeller{
			ID:  &id,
			CID: &cid,
		}

		var respositoryResult *domain.Seller

		repository.On("Get", id).Return(respositoryResult)
		result, err := service.Update(id, updateSeller)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})

	t.Run("Should return a conflict error when cid already exists", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		mockedSeller := mockedSellerTemplate
		id := 1
		cid := 456
		updateSeller := domain.UpdateSeller{
			ID:  &id,
			CID: &cid,
		}

		repository.On("Get", id).Return(&mockedSeller)
		repository.On("Exists", cid).Return(true)
		result, err := service.Update(id, updateSeller)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.ResourceAlreadyExists](err))
	})

	t.Run("Should return a conflict error when locality id not exists", func(t *testing.T) {
		service, repository, localityRepository := CreateService(t)

		mockedSeller := mockedSellerTemplate
		id := 1
		cid := 456
		localityId := 1
		updateSeller := domain.UpdateSeller{
			ID:         &id,
			CID:        &cid,
			LocalityID: &localityId,
		}

		var localityRepositoryGetResult *domain.Locality

		repository.On("Get", id).Return(&mockedSeller)
		repository.On("Exists", cid).Return(false)
		localityRepository.On("Get", localityId).Return(localityRepositoryGetResult)
		result, err := service.Update(id, updateSeller)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.True(t, apperr.Is[*apperr.DependentResourceNotFound](err))
	})

	t.Run("Should return an updated seller", func(t *testing.T) {
		service, repository, localityRepository := CreateService(t)

		mockedSeller := mockedSellerTemplate
		mockedLocality := mockedLocalityTemplate
		id := 1
		cid := 123
		companyName := "Company Name 2"
		localityId := 1
		updateSeller := domain.UpdateSeller{
			ID:          &id,
			CID:         &cid,
			CompanyName: &companyName,
			LocalityID:  &localityId,
		}
		updatedSeller := mockedSeller
		updatedSeller.Overlap(updateSeller)

		repository.On("Get", id).Return(&mockedSeller)
		repository.On("Exists", cid).Return(true)
		localityRepository.On("Get", localityId).Return(&mockedLocality)
		repository.On("Update", updatedSeller)
		repository.On("Get", id).Return(&updatedSeller)
		result, err := service.Update(id, updateSeller)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, companyName, result.CompanyName)
	})
}

func TestServiceDelete(t *testing.T) {
	t.Run("Should return not found error", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		id := 1
		var respositoryResult *domain.Seller

		repository.On("Get", id).Return(respositoryResult)
		err := service.Delete(id)

		assert.Error(t, err)
		assert.True(t, apperr.Is[*apperr.ResourceNotFound](err))
	})

	t.Run("Should delete a seller with success", func(t *testing.T) {
		service, repository, _ := CreateService(t)

		mockedSeller := mockedSellerTemplate
		id := 1

		repository.On("Get", id).Return(&mockedSeller)
		repository.On("Delete", id)
		err := service.Delete(id)

		assert.NoError(t, err)
	})
}

func CreateService(t *testing.T) (seller.Service, *mocks.Repository, *localityMocks.Repository) {
	t.Helper()
	repository := new(mocks.Repository)
	localityRepository := new(localityMocks.Repository)
	service := seller.NewService(repository, localityRepository)
	return service, repository, localityRepository
}
