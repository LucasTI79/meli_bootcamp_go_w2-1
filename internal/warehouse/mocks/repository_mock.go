package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) GetAll(ctx context.Context) []domain.Warehouse {
	args := r.Called()
	return args.Get(0).([]domain.Warehouse)
}

func (r *Repository) Get(ctx context.Context, id int) *domain.Warehouse {
	args := r.Called(id)
	return args.Get(0).(*domain.Warehouse)
}

func (r *Repository) Exists(ctx context.Context, productCode string) bool {
	args := r.Called(productCode)
	return args.Get(0).(bool)
}

func (r *Repository) Save(ctx context.Context, warehouse domain.Warehouse) int {
	args := r.Called(warehouse)
	return args.Get(0).(int)
}

func (r *Repository) Update(ctx context.Context, warehouse domain.Warehouse) {
	r.Called(warehouse)
}

func (r *Repository) Delete(ctx context.Context, id int) {
	r.Called(id)
}
