package seller

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/locality"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	ResourceNotFound      = "vendedor não encontrado com o id %d"
	ResourceAlreadyExists = "um vendedor com o CID '%d' já existe"
)

type Service interface {
	GetAll() []domain.Seller
	Get(id int) (*domain.Seller, error)
	Create(seller domain.Seller) (*domain.Seller, error)
	Update(id int, seller domain.UpdateSeller) (*domain.Seller, error)
	Delete(id int) error
}

type service struct {
	repository         Repository
	localityRepository locality.Repository
}

func NewService(repository Repository, localityRepository locality.Repository) Service {
	return &service{repository, localityRepository}
}

func (s *service) GetAll() []domain.Seller {
	return s.repository.GetAll()
}

func (s *service) Get(id int) (*domain.Seller, error) {
	seller := s.repository.Get(id)

	if seller == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	return seller, nil
}

func (s *service) Create(seller domain.Seller) (*domain.Seller, error) {
	if s.repository.Exists(seller.CID) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, seller.CID)
	}

	localityFound := s.localityRepository.Get(seller.LocalityID)

	if localityFound == nil {
		return nil, apperr.NewDependentResourceNotFound(locality.LocalityNotFound, seller.LocalityID)
	}

	id := s.repository.Save(seller)
	return s.repository.Get(id), nil
}

func (s *service) Update(id int, seller domain.UpdateSeller) (*domain.Seller, error) {
	sellerFound := s.repository.Get(id)

	if sellerFound == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	if seller.CID != nil {
		sellerCID := *seller.CID
		sellerCodeExists := s.repository.Exists(sellerCID)

		if sellerCodeExists && sellerCID != sellerFound.CID {
			return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, sellerCID)
		}
	}

	sellerFound.Overlap(seller)

	localityFound := s.localityRepository.Get(sellerFound.LocalityID)

	if localityFound == nil {
		return nil, apperr.NewDependentResourceNotFound(locality.LocalityNotFound, sellerFound.LocalityID)
	}

	s.repository.Update(*sellerFound)
	return s.repository.Get(id), nil
}

func (s *service) Delete(id int) error {
	seller := s.repository.Get(id)

	if seller == nil {
		return apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	s.repository.Delete(id)
	return nil
}
