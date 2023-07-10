package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (s *Service) GetAll() []domain.Buyer {
	args := s.Called()
	return args.Get(0).([]domain.Buyer)
}

func (s *Service) Get(id int) (*domain.Buyer, error) {
	args := s.Called(id)
	return args.Get(0).(*domain.Buyer), args.Error(1)
}

func (s *Service) Create(p domain.Buyer) (*domain.Buyer, error) {
	args := s.Called(p)
	return args.Get(0).(*domain.Buyer), args.Error(1)
}

func (s *Service) Update(id int, p domain.UpdateBuyer) (*domain.Buyer, error) {
	args := s.Called(id, p)
	return args.Get(0).(*domain.Buyer), args.Error(1)
}

func (s *Service) Delete(id int) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *Service) CountPurchasesByAllBuyers() []domain.PurchasesByBuyerReport {
	args := s.Called()
	return args.Get(0).([]domain.PurchasesByBuyerReport)
}

func (s *Service) CountPurchasesByBuyer(id int) (*domain.PurchasesByBuyerReport, error) {
	args := s.Called(id)
	return args.Get(0).(*domain.PurchasesByBuyerReport), args.Error(1)
}
