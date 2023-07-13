package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Service struct {
	mock.Mock
}

func (s *Service) GetAll() []domain.Employee {
	args := s.Called()
	return args.Get(0).([]domain.Employee)
}

func (s *Service) Get(id int) (*domain.Employee, error) {
	args := s.Called(id)
	return args.Get(0).(*domain.Employee), args.Error(1)
}

func (s *Service) Create(p domain.Employee) (*domain.Employee, error) {
	args := s.Called(p)
	return args.Get(0).(*domain.Employee), args.Error(1)
}

func (s *Service) Update(id int, p domain.UpdateEmployee) (*domain.Employee, error) {
	args := s.Called(id, p)
	return args.Get(0).(*domain.Employee), args.Error(1)
}

func (s *Service) Delete(id int) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *Service) CountInboundOrdersByAllEmployees() []domain.InboundOrdersByEmployee {
	args := s.Called()
	return args.Get(0).([]domain.InboundOrdersByEmployee)
}

func (s *Service) CountInboundOrdersByEmployee(id int) (*domain.InboundOrdersByEmployee, error) {
	args := s.Called(id)
	return args.Get(0).(*domain.InboundOrdersByEmployee), args.Error(1)
}