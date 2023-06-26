package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (s *Service) GetAll(c context.Context) []domain.Section {
	args := s.Called()
	return args.Get(0).([]domain.Section)
}

func (s *Service) Get(c context.Context, id int) (*domain.Section, error) {
	args := s.Called(id)
	return args.Get(0).(*domain.Section), args.Error(1)
}

func (s *Service) Create(c context.Context, sc domain.Section) (*domain.Section, error) {
	args := s.Called(sc)
	return args.Get(0).(*domain.Section), args.Error(1)
}

func (s *Service) Update(c context.Context, id int, sc domain.UpdateSection) (*domain.Section, error) {
	args := s.Called(id, sc)
	return args.Get(0).(*domain.Section), args.Error(1)
}

func (s *Service) Delete(c context.Context, id int) error {
	args := s.Called(id)
	return args.Error(0)
}
