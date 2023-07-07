package mocks

import (

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) GetAll() []domain.Warehouse {
	args := r.Called()
	return args.Get(0).([]domain.Warehouse)
}

func (r *Repository) Get(id int) *domain.Warehouse {
	args := r.Called(id)
	return args.Get(0).(*domain.Warehouse)
}

func (r *Repository) Exists(warehouseCode string) bool {
	args := r.Called(warehouseCode)
	return args.Get(0).(bool)
}

func (r *Repository) Save(warehouse domain.Warehouse) int {
	args := r.Called(warehouse)
	return args.Get(0).(int)
}

func (r *Repository) Update(warehouse domain.Warehouse) {
	r.Called(warehouse)
}

func (r *Repository) Delete(id int) {
	r.Called(id)
}
