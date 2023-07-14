package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (s *Service) Create(p domain.Locality) (*domain.Locality, error) {
	args := s.Called(p)
	return args.Get(0).(*domain.Locality), args.Error(1)
}

func (s *Service) CountSellersByAllLocalities() []domain.SellersByLocalityReport {
	args := s.Called()
	return args.Get(0).([]domain.SellersByLocalityReport)
}

func (s *Service) CountSellersByLocality(id int) (*domain.SellersByLocalityReport, error) {
	args := s.Called(id)
	return args.Get(0).(*domain.SellersByLocalityReport), args.Error(1)
}

func (s *Service) CountCarriersByAllLocalities() []domain.CarriersByLocalityReport {
	args := s.Called()
	return args.Get(0).([]domain.CarriersByLocalityReport)
}

func (s *Service) CountCarriersByLocality(id int) (*domain.CarriersByLocalityReport, error) {
	args := s.Called(id)
	return args.Get(0).(*domain.CarriersByLocalityReport), args.Error(1)
}
