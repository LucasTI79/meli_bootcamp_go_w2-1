package buyer

import (
	"errors"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"context"
) 

// Errors
var (
	ErrNotFound = errors.New("Comprador não encontrado.")
	ErrAlreadyExists = errors.New("Não é possível cadastrar um comprador com Card Number repetido.")
)

type IService interface{
	GetAll(c context.Context) ([]domain.Buyer, error)
	Get(c context.Context, id int) (domain.Buyer, error)
	Save(c context.Context, b domain.Request) (int, error)
	Update(c context.Context, b domain.Buyer) error
	Delete(c context.Context, id int) error
	Exists(c context.Context, cardNumberID string) bool
}

type service struct{
	repository IRepository
}

func NewService(r IRepository) IService {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(c context.Context) ([]domain.Buyer, error) {
	buyers, err := s.repository.GetAll(c)
	if err != nil {
		return nil, err
	}

	return buyers, nil
}

func (s *service) Get(c context.Context, id int) (domain.Buyer, error) {
	buyer, err := s.repository.Get(c, id)
	if err != nil {
		return domain.Buyer{}, ErrNotFound
	}

	return buyer, nil
}

func (s *service) Save(c context.Context, b domain.Request) (int, error) {
	exists := s.repository.Exists(c, b.CardNumberID)

	if !exists {
		id, err := s.repository.Save(c, b)
		if err != nil {
			return 0, err
		}
		return id, nil 
	} else {
		return 0, ErrAlreadyExists
	}

}

func (s *service) Update(c context.Context, b domain.Buyer) error {

	return s.repository.Update(c, b)
	
}

func (s *service) Delete(c context.Context, id int) error {
	
	return s.repository.Delete(c, id)

}

func (s *service) Exists(c context.Context, cardNumberID string) bool {
	
	return s.repository.Exists(c, cardNumberID)
}