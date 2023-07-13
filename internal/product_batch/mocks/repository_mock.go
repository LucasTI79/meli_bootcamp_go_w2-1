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
func (r *Repository) Save(pb domain.ProductBatch) int {
	args := r.Called(pb)
	return args.Int(0)
}
func (r *Repository) Get(id int) *domain.ProductBatch {
	args := r.Called(id)
	return args.Get(0).(*domain.ProductBatch)
}
