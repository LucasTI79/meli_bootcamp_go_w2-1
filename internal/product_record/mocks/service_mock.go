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
