package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) Exists(batchNumber int) bool {
	args := r.Called(batchNumber)
	return args.Bool(0)
}
func (r *Repository) Save(productBatches domain.ProductBatches) int {
	args := r.Called(productBatches)
	return args.Int(0)
}
func (r *Repository) Get(id int) *domain.ProductBatches {
	args := r.Called(id)
	return args.Get(0).(*domain.ProductBatches)
}

func (r *Repository) CountProductsByAllSections() ([]domain.ProductsBySectionReport, error) {
	args := r.Called()
	return args.Get(0).([]domain.ProductsBySectionReport), args.Error(1)
}
func (r *Repository) CountProductsBySection(id int) ([]domain.ProductsBySectionReport, error) {
	args := r.Called(id)
	return args.Get(0).([]domain.ProductsBySectionReport), args.Error(1)
}
