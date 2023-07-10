package inbound_orders

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/employee"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/product_batch"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/warehouse"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	EmployeeNotFound        = "funcionario não encontrado com o id %d"
	ProductBatchNotFound    = "lote de produto não encontrado com o id %d"
	WarehouseNotFound       = "armazem não encontrado com o id %d"
	ResourceAlreadyExists   = "ordem de entrada com o numero '%d' já existe"
)

type Service interface {
	Create(inboundOrder domain.InboundOrder) (*domain.InboundOrder, error)
}

type service struct {
	repository Repository
	employeeRepository employee.Repository
	productBatchRepository product_batch.Repository
	warehouseRepository warehouse.Repository
}

func NewService(repository Repository, employeeRepository employee.Repository, productBatchRepository product_batch.Repository, warehouseRepository warehouse.Repository) Service {
	return &service{repository, employeeRepository, productBatchRepository, warehouseRepository}
} 

func (s *service) Create(inboundOrder domain.InboundOrder) (*domain.InboundOrder, error) {
	if s.repository.Exists(inboundOrder.OrderNumber) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, inboundOrder.OrderNumber)
	}
	
	employee := s.employeeRepository.Get(inboundOrder.EmployeeId)
	if employee == nil {
		return nil, apperr.NewDependentResourceNotFound(EmployeeNotFound, inboundOrder.EmployeeId)
	}
	productBatch := s.productBatchRepository.Get(inboundOrder.ProductBatchId)
	if productBatch == nil {
		return nil, apperr.NewDependentResourceNotFound(ProductBatchNotFound, inboundOrder.ProductBatchId)
	}
	
	warehouse := s.warehouseRepository.Get(inboundOrder.WarehouseId)
	if warehouse == nil {
		return nil, apperr.NewDependentResourceNotFound(WarehouseNotFound, inboundOrder.WarehouseId)
	}
	
	id := s.repository.Save(inboundOrder)
	test := s.repository.Get(id)
	return test, nil
}