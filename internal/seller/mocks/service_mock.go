package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (s *Service) GetAll() []domain.Seller {
	args := s.Called()
	return args.Get(0).([]domain.Seller)
}

func (s *Service) Get(id int) (*domain.Seller, error) {
	args := s.Called(id)
	return args.Get(0).(*domain.Seller), args.Error(1)
}

func (s *Service) Create(p domain.Seller) (*domain.Seller, error) {
	args := s.Called(p)
	return args.Get(0).(*domain.Seller), args.Error(1)
}

func (s *Service) Update(id int, p domain.UpdateSeller) (*domain.Seller, error) {
	args := s.Called(id, p)
	return args.Get(0).(*domain.Seller), args.Error(1)
}

func (s *Service) Delete(id int) error {
	args := s.Called(id)
	return args.Error(0)
}
