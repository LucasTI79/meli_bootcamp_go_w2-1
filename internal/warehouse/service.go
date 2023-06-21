package warehouse

import (
	"context"
	"errors"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w2-1/pkg/apperr"
)

const (
	ResourceNotFound      = "armazém não encontrado com o id '%d'"
	ResourceAlreadyExists = "já existe um armazém com o código '%s'"
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
	GetAll(ctx context.Context) []domain.Warehouse
	Get(ctx context.Context, id int) (*domain.Warehouse, error)
	Create(ctx context.Context, warehouse domain.Warehouse) (*domain.Warehouse, error)
	Update(ctx context.Context, id int, warehouse domain.UpdateWarehouse) (*domain.Warehouse, error)
	Delete(ctx context.Context, id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll(ctx context.Context) []domain.Warehouse {
	return s.repository.GetAll(ctx)
}

func (s *service) Get(ctx context.Context, id int) (*domain.Warehouse, error) {
	warehouse := s.repository.Get(ctx, id)

	if warehouse == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	return warehouse, nil
}

func (s *service) Create(ctx context.Context, warehouse domain.Warehouse) (*domain.Warehouse, error) {
	if s.repository.Exists(ctx, warehouse.WarehouseCode) {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, warehouse.WarehouseCode)
	}

	warehouseId := s.repository.Save(ctx, warehouse)
	return s.repository.Get(ctx, warehouseId), nil
}

func (s *service) Update(ctx context.Context, id int, warehouse domain.UpdateWarehouse) (*domain.Warehouse, error) {
	warehouseFound := s.repository.Get(ctx, id)

	if warehouseFound == nil {
		return nil, apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	warehouseCode := *warehouse.WarehouseCode
	warehouseExists := s.repository.Exists(ctx, warehouseCode)

	if warehouseExists && warehouseCode != warehouseFound.WarehouseCode {
		return nil, apperr.NewResourceAlreadyExists(ResourceAlreadyExists, warehouseCode)
	}

	warehouseFound.Overlap(warehouse)

	s.repository.Update(ctx, *warehouseFound)

	updated := s.repository.Get(ctx, id)
	return updated, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	warehouse := s.repository.Get(ctx, id)

	if warehouse == nil {
		return apperr.NewResourceNotFound(ResourceNotFound, id)
	}

	s.repository.Delete(ctx, id)
	return nil
}
