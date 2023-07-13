package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (s Service) Get(id int) (*domain.Carrier, error) {
	args := s.Called(id)
	return args.Get(0).(*domain.Carrier), args.Error(1)
}

func (s Service) Create(carrier domain.Carrier) (*domain.Carrier, error) {
	args := s.Called(carrier)
	return args.Get(0).(*domain.Carrier), args.Error(1)
}
