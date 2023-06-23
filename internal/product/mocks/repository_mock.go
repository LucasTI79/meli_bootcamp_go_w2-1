package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) GetAll(ctx context.Context) []domain.Product {
	args := r.Called()
	return args.Get(0).([]domain.Product)
}

func (r *Repository) Get(ctx context.Context, id int) *domain.Product {
	args := r.Called(id)
	return args.Get(0).(*domain.Product)
}

func (r *Repository) Exists(ctx context.Context, productCode string) bool {
	args := r.Called(productCode)
	return args.Get(0).(bool)
}

func (r *Repository) Save(ctx context.Context, product domain.Product) int {
	args := r.Called(product)
	return args.Get(0).(int)
}

func (r *Repository) Update(ctx context.Context, product domain.Product) {
	// args := r.Called(product)
	// return args.Get(0).(int)
}

func (r *Repository) Delete(ctx context.Context, id int) {
}
