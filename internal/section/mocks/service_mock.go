package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (s *Service) GetAll() []domain.Section {
	args := s.Called()
	return args.Get(0).([]domain.Section)
}

func (s *Service) Get(id int) (*domain.Section, error) {
	args := s.Called(id)
	return args.Get(0).(*domain.Section), args.Error(1)
}

func (s *Service) Create(sc domain.Section) (*domain.Section, error) {
	args := s.Called(sc)
	return args.Get(0).(*domain.Section), args.Error(1)
}

func (s *Service) Update(id int, sc domain.UpdateSection) (*domain.Section, error) {
	args := s.Called(id, sc)
	return args.Get(0).(*domain.Section), args.Error(1)
}

func (s *Service) Delete(id int) error {
	args := s.Called(id)
	return args.Error(0)
}
