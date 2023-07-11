package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (s *Service) Create(pb domain.ProductBatches) (*domain.ProductBatches, error) {
	args := s.Called(pb)
	return args.Get(0).(*domain.ProductBatches), args.Error(1)
}

func (s *Service) Exists(batchNumber int) (bool, error) {
	args := s.Called(batchNumber)
	return args.Bool(0), args.Error(1)
}

func (s *Service) Get() ([]domain.ProductsBySectionReport, error) {
	args := s.Called()
	return args.Get(0).([]domain.ProductsBySectionReport), args.Error(1)
}

func (s *Service) CountProductsBySection(id int) ([]domain.ProductsBySectionReport, error) {
	args := s.Called(id)
	return args.Get(0).([]domain.ProductsBySectionReport), args.Error(1)
}

func (s *Service) CountProductsByAllSections() ([]domain.ProductsBySectionReport, error) {
	args := s.Called()
	return args.Get(0).([]domain.ProductsBySectionReport), args.Error(1)
}
