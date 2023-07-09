package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (s *Service) GetAll() []domain.Product {
	args := s.Called()
	return args.Get(0).([]domain.Product)
}

func (s *Service) Get(id int) (*domain.Product, error) {
	args := s.Called(id)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (s *Service) Create(p domain.Product) (*domain.Product, error) {
	args := s.Called(p)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (s *Service) Update(id int, p domain.UpdateProduct) (*domain.Product, error) {
	args := s.Called(id, p)
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (s *Service) Delete(id int) error {
	args := s.Called(id)
	return args.Error(0)
}
