package buyer

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	ResourceNotFound      = "comprador não encontrado com o id %d"
	ResourceAlreadyExists = "um comprador com o número de cartão '%s' já existe"
)

type Service interface {
	GetAll() []domain.Buyer
	Get(id int) (*domain.Buyer, error)
	Create(b domain.Buyer) (*domain.Buyer, error)
	Update(id int, b domain.UpdateBuyer) (*domain.Buyer, error)
	Delete(id int) error
	CountPurchasesByAllBuyers() []domain.PurchasesByBuyerReport
	CountPurchasesByBuyer(id int) (*domain.PurchasesByBuyerReport, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CountPurchasesByAllBuyers() []domain.PurchasesByBuyerReport {
	return s.repository.CountPurchasesByAllBuyers()
}

func (s *service) CountPurchasesByBuyer(id int) (*domain.PurchasesByBuyerReport, error) {
	buyer := s.repository.Get(id)

	if buyer == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	return s.repository.CountPurchasesByBuyer(id), nil
}

func (s *service) GetAll() []domain.Buyer {
	return s.repository.GetAll()
}

func (s *service) Get(id int) (*domain.Buyer, error) {
	buyer := s.repository.Get(id)

	if buyer == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	return buyer, nil
}

func (s *service) Create(b domain.Buyer) (*domain.Buyer, error) {
	exists := s.repository.Exists(b.CardNumberID)

	if !exists {
		id := s.repository.Save(b)
		return s.repository.Get(id), nil
	}

	return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, b.CardNumberID)
}

func (s *service) Update(id int, buyer domain.UpdateBuyer) (*domain.Buyer, error) {
	buyerFound := s.repository.Get(id)

	if buyerFound == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	if buyer.CardNumberID != nil {
		cardNumberID := *buyer.CardNumberID
		cardNumberIDExists := s.repository.Exists(cardNumberID)

		if cardNumberIDExists && cardNumberID != buyerFound.CardNumberID {
			return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, cardNumberID)
		}
	}

	buyerFound.Overlap(buyer)

	s.repository.Update(*buyerFound)
	updated := s.repository.Get(id)

	return updated, nil
}

func (s *service) Delete(id int) error {
	buyer := s.repository.Get(id)

	if buyer == nil {
		return apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	s.repository.Delete(id)
	return nil
}
