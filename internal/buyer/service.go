package buyer

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	ResourceNotFound      = "comprador não encontrado com o id %d"
	ResourceAlreadyExists = "um comprador com o número de cartão '%s' já existe"
)

type IService interface {
	GetAll(c context.Context) []domain.Buyer
	Get(c context.Context, id int) (*domain.Buyer, error)
	Create(c context.Context, b domain.Buyer) (*domain.Buyer, error)
	Update(c context.Context, id int, b domain.UpdateBuyer) (*domain.Buyer, error)
	Delete(c context.Context, id int) error
}

type service struct {
	repository IRepository
}

func NewService(r IRepository) IService {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(c context.Context) []domain.Buyer {
	return s.repository.GetAll(c)
}

func (s *service) Get(c context.Context, id int) (*domain.Buyer, error) {
	buyer := s.repository.Get(c, id)

	if buyer == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	return buyer, nil
}

func (s *service) Create(c context.Context, b domain.Buyer) (*domain.Buyer, error) {
	exists := s.repository.Exists(c, b.CardNumberID)

	if !exists {
		id := s.repository.Save(c, b)
		return s.repository.Get(c, id), nil
	}

	return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, b.CardNumberID)
}

func (s *service) Update(c context.Context, id int, buyer domain.UpdateBuyer) (*domain.Buyer, error) {
	buyerFound := s.repository.Get(c, id)

	if buyerFound == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	if buyer.CardNumberID != nil {
		cardNumberID := *buyer.CardNumberID
		cardNumberIDExists := s.repository.Exists(c, cardNumberID)

		if cardNumberIDExists && cardNumberID != buyerFound.CardNumberID {
			return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, cardNumberID)
		}
	}

	buyerFound.Overlap(buyer)

	s.repository.Update(c, *buyerFound)
	updated := s.repository.Get(c, id)

	return updated, nil
}

func (s *service) Delete(c context.Context, id int) error {
	buyer := s.repository.Get(c, id)

	if buyer == nil {
		return apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	s.repository.Delete(c, id)
	return nil
}
