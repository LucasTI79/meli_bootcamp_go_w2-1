package warehouse

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

// Errors
var (
	ErrWarehouseNotFound  = errors.New("warehouse not found")
	ErrWarehouseExists    = errors.New("warehouse already exists")
	ErrInvalidID          = errors.New("invalid ID")
	ErrInvalidPhoneNumber = errors.New("invalid phone number")
	ErrMissingField       = errors.New("missing required field")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Warehouse, error)
	Get(ctx context.Context, id int) (domain.Warehouse, error)
	Create(ctx context.Context, warehouseData domain.Warehouse) (*domain.Warehouse, error)
	Save(ctx context.Context, warehouse domain.Warehouse) (int, error)
	Update(ctx context.Context, warehouseData domain.Warehouse) (domain.Warehouse, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	return s.repository.Get(ctx, id)
}

func (s *service) Create(ctx context.Context, warehouseData domain.Warehouse) (*domain.Warehouse, error) {
	if err := validateWarehouseData(warehouseData); err != nil {
		return nil, err
	}

	if s.repository.Exists(ctx, warehouseData.WarehouseCode) {
		return nil, ErrWarehouseExists
	}

	warehouse := domain.Warehouse{
		Address:            warehouseData.Address,
		Telephone:          warehouseData.Telephone,
		WarehouseCode:      warehouseData.WarehouseCode,
		MinimumCapacity:    warehouseData.MinimumCapacity,
		MinimumTemperature: warehouseData.MinimumTemperature,
	}

	id, err := s.repository.Save(ctx, warehouse)
	if err != nil {
		return nil, err
	}

	warehouse.ID = id

	return &warehouse, nil
}

func (s *service) Save(ctx context.Context, warehouse domain.Warehouse) (int, error) {
	if s.repository.Exists(ctx, warehouse.WarehouseCode) {
		return 0, ErrWarehouseExists
	}

	wCode, err := s.repository.Save(ctx, warehouse)
	if err != nil {
		return 0, err
	}
	return wCode, nil
}

func (s *service) Update(ctx context.Context, warehouseData domain.Warehouse) (domain.Warehouse, error) {

	if err := validateWarehouseData(warehouseData); err != nil {
		return domain.Warehouse{}, err
	}

	existingWarehouse, err := s.repository.GetByCode(ctx, warehouseData.WarehouseCode)
	if err != nil && err != ErrWarehouseNotFound {
		return domain.Warehouse{}, err
	}

	if existingWarehouse.ID != warehouseData.ID {
		return domain.Warehouse{}, ErrWarehouseExists
	}

	warehouse, err := s.repository.Get(ctx, warehouseData.ID)
	if err != nil {
		if err == ErrWarehouseNotFound {
			return domain.Warehouse{}, err
		}
		return domain.Warehouse{}, err
	}

	originalWarehouseCode := warehouse.WarehouseCode

	warehouse.Address = warehouseData.Address
	warehouse.Telephone = warehouseData.Telephone
	warehouse.WarehouseCode = warehouseData.WarehouseCode
	warehouse.MinimumCapacity = warehouseData.MinimumCapacity
	warehouse.MinimumTemperature = warehouseData.MinimumTemperature

	warehouse.WarehouseCode = originalWarehouseCode

	err = s.repository.Update(ctx, warehouse)
	if err != nil {
		return domain.Warehouse{}, err
	}

	return warehouse, err
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		if err == ErrWarehouseNotFound {
			return err
		}
		return err
	}

	return nil
}

func validateWarehouseData(warehouseData domain.Warehouse) error {
	if warehouseData.WarehouseCode == "" ||
		warehouseData.Address == "" ||
		warehouseData.Telephone == "" ||
		warehouseData.MinimumCapacity == 0 ||
		warehouseData.MinimumTemperature == 0 {
		//indica que um campo obrigatorio esta faltando
		return ErrMissingField
	}

	return nil
}
