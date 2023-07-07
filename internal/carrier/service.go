package carrier

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	ResourceNotFound = "transportadora não encontrada com o id %d"
)

type Service interface {
	Get(id int) (*domain.Carrier, error)
	Create(domain.Employee) (*domain.Carrier, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Get(id int) (*domain.Carrier, error) {
	carrier := s.repository.Get(id)
	if carrier == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}
	return carrier, nil
}

func (s *service) Create(employee domain.Employee) (*domain.Carrier, error) {

	//TODO implement me
	panic("implement me")
}
