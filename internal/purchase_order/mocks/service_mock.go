package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (s *Service) Create(po domain.PurchaseOrder) (*domain.PurchaseOrder, error) {
	args := s.Called(po)
	return args.Get(0).(*domain.PurchaseOrder), args.Error(1)
}
