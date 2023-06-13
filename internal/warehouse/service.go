package warehouse

import (
	"context"
	"errors"
	"regexp"
	"strings"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
)

// Errors
var (
	ErrWarehouseNotFound  = errors.New("warehouse not found")
	ErrWarehouseExists    = errors.New("warehouse already exists")
	ErrInvalidID          = errors.New("invalid ID")
	ErrInvalidPhoneNumber = errors.New("invalid phone number")
	ErrMissingField       = errors.New("missing required field")
	ErrorResponse         = errors.New("Nao localizado no momento")
)

type WarehouseService interface {
	GetAllWarehouses(ctx context.Context) ([]domain.Warehouse, error)
	GetWarehouse(ctx context.Context, id int) (domain.Warehouse, error)
	CreateWarehouse(ctx context.Context, warehouseData domain.Warehouse) (*domain.Warehouse, error)
	UpdateWarehouse(ctx context.Context, warehouseData domain.Warehouse) error
	DeleteWarehouse(ctx context.Context, id int) error
}

type WarehouseData struct {
	ID                 int
	WarehouseCode      string
	Address            string
	Telephone          string
	MinimumCapacity    int
	MinimumTemperature int
}

func formatPhoneNumber(phoneNumber string) string {

	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")
	phoneNumber = "+55" + phoneNumber

	return phoneNumber
}

type warehouseService struct {
	warehouseRepository Repository
}

func NewWarehouseService(repository Repository) WarehouseService {
	return &warehouseService{
		warehouseRepository: repository,
	}
}

func (s *warehouseService) GetAllWarehouses(ctx context.Context) ([]domain.Warehouse, error) {
	return s.warehouseRepository.GetAll(ctx)
}

func (s *warehouseService) GetWarehouse(ctx context.Context, id int) (domain.Warehouse, error) {
	return s.warehouseRepository.Get(ctx, id)
}

func (s *warehouseService) CreateWarehouse(ctx context.Context, warehouseData domain.Warehouse) (*domain.Warehouse, error) {
	if err := validateWarehouseData(warehouseData); err != nil {
		return nil, err
	}

	if s.warehouseRepository.Exists(ctx, warehouseData.WarehouseCode) {
		return nil, ErrWarehouseExists
	}

	warehouse := domain.Warehouse{
		Address:            warehouseData.Address,
		Telephone:          formatPhoneNumber(warehouseData.Telephone),
		WarehouseCode:      warehouseData.WarehouseCode,
		MinimumCapacity:    warehouseData.MinimumCapacity,
		MinimumTemperature: warehouseData.MinimumTemperature,
	}

	id, err := s.warehouseRepository.Save(ctx, warehouse)
	if err != nil {
		return nil, err
	}

	warehouse.ID = id

	return &warehouse, nil
}

func (s *warehouseService) UpdateWarehouse(ctx context.Context, warehouseData domain.Warehouse) error {
	if err := validateWarehouseData(warehouseData); err != nil {
		return err
	}

	existingWarehouse, err := s.warehouseRepository.GetByCode(ctx, warehouseData.WarehouseCode)
	if err != nil && err != ErrNotFound {
		return err
	}

	if existingWarehouse.ID != warehouseData.ID {
		return ErrWarehouseExists
	}
	warehouse, err := s.warehouseRepository.Get(ctx, warehouseData.ID)
	if err != nil {
		if err == ErrNotFound {
			return ErrWarehouseNotFound
		}
		return err
	}

	warehouse.Address = warehouseData.Address
	warehouse.Telephone = formatPhoneNumber(warehouseData.Telephone)
	warehouse.WarehouseCode = warehouseData.WarehouseCode
	warehouse.MinimumCapacity = warehouseData.MinimumCapacity
	warehouse.MinimumTemperature = warehouseData.MinimumTemperature

	return s.warehouseRepository.Update(ctx, warehouse)
}

func (s *warehouseService) DeleteWarehouse(ctx context.Context, id int) error {
	err := s.warehouseRepository.Delete(ctx, id)
	if err != nil {
		if err == ErrNotFound {
			return ErrWarehouseNotFound
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
	//validar o tel fornecido
	if !isValidPhoneNumber(warehouseData.Telephone) {
		return ErrInvalidPhoneNumber
	}

	return nil
}

func isValidPhoneNumber(phoneNumber string) bool {
	// Padrão Regex para validar o formato do número de telefone: DDI + DDD + number (String)
	pattern := `^\+\d{1,3}\s?\d{1,3}\s?\d+$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(phoneNumber)
}
