package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (s *Service) GetAll(c context.Context) []domain.Employee {
	return nil
}

func (s *Service) Get(c context.Context, id int) (*domain.Employee, error) {
	return nil, nil
}

func (s *Service) Create(c context.Context, p domain.Employee) (*domain.Employee, error) {
	args := s.Called(p)
	return args.Get(0).(*domain.Employee), args.Error(1)
}

func (s *Service) Update(c context.Context, id int, p domain.UpdateEmployee) (*domain.Employee, error) {
	return nil, nil
}

func (s *Service) Delete(c context.Context, id int) error {
	return nil
}