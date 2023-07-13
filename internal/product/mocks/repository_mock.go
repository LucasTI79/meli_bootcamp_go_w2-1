package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) GetAll() []domain.Product {
	args := r.Called()
	return args.Get(0).([]domain.Product)
}

func (r *Repository) Get(id int) *domain.Product {
	args := r.Called(id)
	return args.Get(0).(*domain.Product)
}

func (r *Repository) Exists(productCode string) bool {
	args := r.Called(productCode)
	return args.Get(0).(bool)
}

func (r *Repository) Save(product domain.Product) int {
	args := r.Called(product)
	return args.Get(0).(int)
}

func (r *Repository) Update(product domain.Product) {
	r.Called(product)
}

func (r *Repository) Delete(id int) {
	r.Called(id)
}

func (r *Repository) CountRecordsByAllProducts() []domain.RecordsByProductReport {
	args := r.Called()
	return args.Get(0).([]domain.RecordsByProductReport)
}
func (r *Repository) CountRecordsByProduct(id int) *domain.RecordsByProductReport {
	args := r.Called(id)
	return args.Get(0).(*domain.RecordsByProductReport)
}
