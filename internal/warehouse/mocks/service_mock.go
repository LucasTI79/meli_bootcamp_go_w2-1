package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (s *Service) GetAll() []domain.Warehouse {
	args := s.Called()
	return args.Get(0).([]domain.Warehouse)
}

func (s *Service) Get(id int) (*domain.Warehouse, error) {
	args := s.Called(id)
	return args.Get(0).(*domain.Warehouse), args.Error(1)
}

func (s *Service) Create(w domain.Warehouse) (*domain.Warehouse, error) {
	args := s.Called(w)
	return args.Get(0).(*domain.Warehouse), args.Error(1)
}

func (s *Service) Update(id int, w domain.UpdateWarehouse) (*domain.Warehouse, error) {
	args := s.Called(id, w)
	return args.Get(0).(*domain.Warehouse), args.Error(1)
}

func (s *Service) Delete(id int) error {
	args := s.Called(id)
	return args.Error(0)
}
