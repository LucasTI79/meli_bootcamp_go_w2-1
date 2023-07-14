package carrier

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/locality"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	ResourceNotFound      = "transportadora não encontrada com o id %d"
	LocalityNotFound      = "localidade não encontrada com o id %d"
	ResourceAlreadyExists = "uma transportadora com cid '%s' já existe"
)

type Service interface {
	Create(carrier domain.Carrier) (*domain.Carrier, error)
}

type service struct {
	repository         Repository
	localityRepository locality.Repository
}

func NewService(r Repository, localityRepository locality.Repository) Service {
	return &service{
		repository:         r,
		localityRepository: localityRepository,
	}
}

func (s *service) Create(carrier domain.Carrier) (*domain.Carrier, error) {
	if s.repository.Exists(carrier.CID) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, carrier.CID)
	}
	localityById := s.localityRepository.Get(carrier.LocalityID)
	if localityById == nil {
		return nil, apperr.NewDependentResourceNotFound(LocalityNotFound, carrier.LocalityID)
	}
	id := s.repository.Save(carrier)
	c := s.repository.Get(id)
	return c, nil

}
