package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) Create(productBatches domain.ProductBatches) error {
	args := r.Called(productBatches)
	return args.Error(0)
}
func (r *Repository) Exists(batchNumber int) bool {
	args := r.Called(batchNumber)
	return args.Bool(0)
}
func (r *Repository) Save(productBatches domain.ProductBatches) (int, error) {
	args := r.Called(productBatches)
	return args.Int(0), args.Error(1)
}
func (r *Repository) Get() ([]domain.ProductBatches, error) {
	args := r.Called()
	return args.Get(0).([]domain.ProductBatches), args.Error(1)
}
