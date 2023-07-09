package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (s *Service) Create(p domain.ProductRecord) (*domain.ProductRecord, error) {
	args := s.Called(p)
	return args.Get(0).(*domain.ProductRecord), args.Error(1)
}

func (s *Service) CountRecordsByAllProducts() []domain.RecordsByProductReport {
	args := s.Called()
	return args.Get(0).([]domain.RecordsByProductReport)
}

func (s *Service) CountRecordsByProduct(id int) (*domain.RecordsByProductReport, error) {
	args := s.Called(id)
	return args.Get(0).(*domain.RecordsByProductReport), args.Error(1)
}
