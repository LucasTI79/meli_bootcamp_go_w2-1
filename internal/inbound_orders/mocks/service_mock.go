package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (s *Service) Create(inboundOrder domain.InboundOrder) (*domain.InboundOrder, error) {
		args := s.Called(inboundOrder)
		return args.Get(0).(*domain.InboundOrder), args.Error(1)
	}
