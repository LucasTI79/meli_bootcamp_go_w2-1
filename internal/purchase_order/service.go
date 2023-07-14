package purchase_order

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/buyer"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/carrier"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/order_status"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_record"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	BuyerNotFound         = "comprador não encontrado com o id %d"
	OrderStatusNotFound   = "status da ordem não encontrado com o id %d"
	WarehouseNotFound     = "armazém não encontrado com o id %d"
	CarrierNotFound       = "transportadora não encontrada com o id %d"
	ProductRecordNotFound = "registro de produto não encontrado com o id %d"
	ResourceAlreadyExists = "uma ordem de compra com o número '%s' já existe"
)

type Service interface {
	Create(locality domain.PurchaseOrder) (*domain.PurchaseOrder, error)
}

type service struct {
	repository              Repository
	buyerRepository         buyer.Repository
	orderStatusrepository   order_status.Repository
	warehouseRepository     warehouse.Repository
	carrierRepository       carrier.Repository
	productRecordRepository product_record.Repository
}

func NewService(repository Repository, buyerRepository buyer.Repository, orderStatusrepository order_status.Repository, warehouseRepository warehouse.Repository, carrierRepository carrier.Repository, productRecordRepository product_record.Repository) Service {

	return &service{repository, buyerRepository, orderStatusrepository, warehouseRepository, carrierRepository, productRecordRepository}
}

func (s *service) Create(po domain.PurchaseOrder) (*domain.PurchaseOrder, error) {
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

	warehouseFound := s.warehouseRepository.Get(po.WarehouseID)
	if warehouseFound == nil {
		return nil, apperr.NewDependentResourceNotFound(WarehouseNotFound, po.WarehouseID)
	}

	productRecordFound := s.productRecordRepository.Get(po.ProductRecordID)
	if productRecordFound == nil {
		return nil, apperr.NewDependentResourceNotFound(ProductRecordNotFound, po.ProductRecordID)
	}

	carrierFound := s.carrierRepository.Get(po.CarrierID)
	if carrierFound == nil {
		return nil, apperr.NewDependentResourceNotFound(CarrierNotFound, po.CarrierID)
	}

	id := s.repository.Save(po)
	return s.repository.Get(id), nil
}
