package locality

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/province"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	LocalityNotFound      = "localização não encontrada com o id %d"
	ProvinceNotFound      = "estado não encontrado com o id %d"
	ResourceAlreadyExists = "uma localização com o nome '%s' já existe"
)

type Service interface {
	CountSellersByAllLocalities() []domain.SellersByLocalityReport
	CountSellersByLocality(id int) (*domain.SellersByLocalityReport, error)
	Create(locality domain.Locality) (*domain.Locality, error)
}

type service struct {
	repository         Repository
	provinceRepository province.Repository
}

func NewService(repository Repository, provinceRepository province.Repository) Service {
	return &service{repository, provinceRepository}
}

func (s *service) CountSellersByAllLocalities() []domain.SellersByLocalityReport {
	return s.repository.CountSellersByAllLocalities()
}

func (s *service) CountSellersByLocality(id int) (*domain.SellersByLocalityReport, error) {
	locality := s.repository.Get(id)

	if locality == nil {
		return nil, apperr.NewResourceNotFound(LocalityNotFound, id)
	}

	return s.repository.CountSellersByLocality(id), nil
}

func (s *service) Get(id int) (*domain.Locality, error) {
	locality := s.repository.Get(id)

	if locality == nil {
		return nil, apperr.NewResourceNotFound(LocalityNotFound, id)
	}

	return locality, nil
}

func (s *service) Create(locality domain.Locality) (*domain.Locality, error) {
	if s.repository.Exists(locality.LocalityName) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, locality.LocalityName)
	}

	provinceFound := s.provinceRepository.Get(locality.ProvinceID)

	if provinceFound == nil {
		return nil, apperr.NewDependentResourceNotFound(ProvinceNotFound, locality.ProvinceID)
	}

	id := s.repository.Save(locality)
	return s.repository.Get(id), nil
}
