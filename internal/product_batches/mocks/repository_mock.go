package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) Get(id int) *domain.ProductBatches {
	args := r.Called(id)
	return args.Get(0).(*domain.ProductBatches)
}
func (r *Repository) Exists(BatchNumber string) bool {
	args := r.Called(BatchNumber)
	return args.Get(0).(bool)
}
func (r *Repository) Save(pb domain.ProductBatches) int {
	args := r.Called(pb)
	return args.Get(0).(int)
}
func (r *Repository) CheckSectionExists(id int) bool {
	args := r.Called(id)
	return args.Get(0).(bool)
}
func (r *Repository) CheckProductExists(id int) bool {
	args := r.Called(id)
	return args.Get(0).(bool)
}
func (r *Repository) CountProductBatchesBySection() []domain.CountProductBatchesBySection {
	args := r.Called()
	return args.Get(0).([]domain.CountProductBatchesBySection)
}
