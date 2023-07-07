package purchase_orders

import (

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/order_status"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)
const (
	BuyerNotFound = "comprador não encontrado com o id %d"
	OrderStatusNotFound = "status da ordem não encontrado com o id %d"
	WarehouseNotFound = "armazém não encontrado com o id %d"
	ResourceAlreadyExists = "uma ordem de compra com o número '%s' já existe"
)

type Service interface {
	Create(locality domain.PurchaseOrders) (*domain.PurchaseOrders, error)
}

type service struct {
	repository         		Repository
	buyerRepository 		buyer.Repository
	orderStatusrepository 	order_status.Repository
	warehouserepository 	warehouse.Repository
}

func NewService(repository Repository, buyerRepository buyer.Repository, orderStatusrepository 	order_status.Repository, warehouserepository warehouse.Repository) Service {
	return &service{repository, buyerRepository, orderStatusrepository, warehouserepository}
}

func (s *service) Create(po domain.PurchaseOrders) (*domain.PurchaseOrders, error) {
	if s.repository.Exists(po.OrderNumber) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, po.OrderNumber)
	}

	buyerFound := s.buyerRepository.Get(po.BuyerID)
	if buyerFound == nil {
		return nil, apperr.NewDependentResourceNotFound(BuyerNotFound, po.BuyerID)
	}

	orderStatusFound := s.orderStatusrepository.Get(po.OrderStatusID)
	if orderStatusFound == nil {
		return nil, apperr.NewDependentResourceNotFound(OrderStatusNotFound, po.OrderStatusID)
	}

	warehouseFound := s.warehouserepository.Get(po.WarehouseID)
	if warehouseFound == nil {
		return nil, apperr.NewDependentResourceNotFound(WarehouseNotFound, po.WarehouseID)
	}

	id := s.repository.Save(po)
	return s.repository.Get(id), nil
}
