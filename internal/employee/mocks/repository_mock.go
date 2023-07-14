package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) GetAll() []domain.Employee {
	args := r.Called()
	return args.Get(0).([]domain.Employee)
}

func (r *Repository) Get(id int) *domain.Employee {
	args := r.Called(id)
	return args.Get(0).(*domain.Employee)
}

func (r *Repository) Exists(CardNumberID string) bool {
	args := r.Called(CardNumberID)
	return args.Get(0).(bool)
}

func (r *Repository) Save(employee domain.Employee) int {
	args := r.Called(employee)
	return args.Get(0).(int)
}

func (r *Repository) Update(employee domain.Employee) {
	r.Called(employee)
}

func (r *Repository) Delete(id int) {
	r.Called(id)
}

func (r *Repository) CountInboundOrdersByAllEmployees() []domain.InboundOrdersByEmployee {
	args := r.Called()
	return args.Get(0).([]domain.InboundOrdersByEmployee)
}

func (r *Repository) CountInboundOrdersByEmployee(id int) *domain.InboundOrdersByEmployee {
	args := r.Called(id)
	return args.Get(0).(*domain.InboundOrdersByEmployee)
}