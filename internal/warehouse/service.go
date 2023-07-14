package warehouse

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/locality"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	ResourceNotFound      = "armazém não encontrado com o id '%d'"
	ResourceAlreadyExists = "já existe um armazém com o código '%s'"
	LocalityNotFound      = "localidade não encontrada com o id '%d'"
)

type Service interface {
	GetAll() []domain.Warehouse
	Get(id int) (*domain.Warehouse, error)
	Create(warehouse domain.Warehouse) (*domain.Warehouse, error)
	Update(id int, warehouse domain.UpdateWarehouse) (*domain.Warehouse, error)
	Delete(id int) error
}

type service struct {
	repository         Repository
	localityRepository locality.Repository
}

func NewService(repository Repository, localityRepository locality.Repository) Service {
	return &service{repository, localityRepository}
}

func (s *service) GetAll() []domain.Warehouse {
	return s.repository.GetAll()
}

func (s *service) Get(id int) (*domain.Warehouse, error) {
	warehouse := s.repository.Get(id)

	if warehouse == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	return warehouse, nil
}

func (s *service) Create(warehouse domain.Warehouse) (*domain.Warehouse, error) {
	if s.repository.Exists(warehouse.WarehouseCode) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, warehouse.WarehouseCode)
	}

	locality := s.localityRepository.Get(warehouse.LocalityID)

	if locality == nil {
		return nil, apperr.NewDependentResourceNotFound(LocalityNotFound, warehouse.LocalityID)
	}

	warehouseId := s.repository.Save(warehouse)
	return s.repository.Get(warehouseId), nil
}

func (s *service) Update(id int, warehouse domain.UpdateWarehouse) (*domain.Warehouse, error) {
	warehouseFound := s.repository.Get(id)

	if warehouseFound == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	if warehouse.WarehouseCode != nil {
		warehouseCode := *warehouse.WarehouseCode
		warehouseExists := s.repository.Exists(warehouseCode)

		if warehouseExists && warehouseCode != warehouseFound.WarehouseCode {
			return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, warehouseCode)
		}
	}

	warehouseFound.Overlap(warehouse)

	locality := s.localityRepository.Get(warehouseFound.LocalityID)

	if locality == nil {
		return nil, apperr.NewDependentResourceNotFound(LocalityNotFound, warehouseFound.LocalityID)
	}

	s.repository.Update(*warehouseFound)

	updated := s.repository.Get(id)
	return updated, nil
}

func (s *service) Delete(id int) error {
	warehouse := s.repository.Get(id)

	if warehouse == nil {
		return apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	s.repository.Delete(id)
	return nil
}
