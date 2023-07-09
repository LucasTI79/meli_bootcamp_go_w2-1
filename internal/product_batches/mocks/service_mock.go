package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (s *Service) Exists(BatchNumber string) bool {
	args := s.Called(BatchNumber)
	return args.Get(0).(bool)
}
func (s *Service) CheckSectionExists(id int) bool {
	args := s.Called(id)
	return args.Get(0).(bool)
}
func (s *Service) CheckProductExists(id int) bool {
	args := s.Called(id)
	return args.Get(0).(bool)
}
func (s *Service) Create(pb domain.ProductBatches) int {
	args := s.Called(pb)
	return args.Get(0).(int)
}
func (s *Service) CountProductBatchesBySection() []domain.CountProductBatchesBySection {
	args := s.Called()
	return args.Get(0).([]domain.CountProductBatchesBySection)
}
