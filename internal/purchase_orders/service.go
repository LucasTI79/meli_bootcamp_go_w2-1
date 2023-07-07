package purchase_orders

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)
const (
	BuyerNotFound = "comprador não encontrado com o id %d"
	ResourceAlreadyExists = "uma ordem de compra com o número '%s' já existe"
)

type Service interface {
	Create(locality domain.PurchaseOrders) (*domain.PurchaseOrders, error)
}

type service struct {
	repository         	Repository
	buyerRepository 	buyer.Repository
}

func NewService(repository Repository, buyerRepository buyer.Repository) Service {
	return &service{repository, buyerRepository}
}

func (s *service) Create(po domain.PurchaseOrders) (*domain.PurchaseOrders, error) {
	if s.repository.Exists(po.OrderNumber) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, po.OrderNumber)
	}

	buyerFound := s.buyerRepository.Get(po.BuyerID)

	if buyerFound == nil {
		return nil, apperr.NewDependentResourceNotFound(BuyerNotFound, po.BuyerID)
	}

	id := s.repository.Save(po)
	return s.repository.Get(id), nil
}
