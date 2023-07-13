package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (s *Service) Create(pb domain.ProductBatch) (*domain.ProductBatch, error) {
	args := s.Called(pb)
	return args.Get(0).(*domain.ProductBatch), args.Error(1)
}
