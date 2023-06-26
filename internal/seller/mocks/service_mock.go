package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (s *Service) GetAll(c context.Context) []domain.Seller {
	args := s.Called()
	return args.Get(0).([]domain.Seller)
}

func (s *Service) Get(c context.Context, id int) (*domain.Seller, error) {
	args := s.Called(id)
	return args.Get(0).(*domain.Seller), args.Error(1)
}

func (s *Service) Create(c context.Context, p domain.Seller) (*domain.Seller, error) {
	args := s.Called(p)
	return args.Get(0).(*domain.Seller), args.Error(1)
}

func (s *Service) Update(c context.Context, id int, p domain.UpdateSeller) (*domain.Seller, error) {
	args := s.Called(id, p)
	return args.Get(0).(*domain.Seller), args.Error(1)
}

func (s *Service) Delete(c context.Context, id int) error {
	args := s.Called(id)
	return args.Error(0)
}
