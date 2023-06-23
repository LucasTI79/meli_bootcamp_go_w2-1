package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) GetAll(ctx context.Context) []domain.Employee {
	return nil
}

func (r *Repository) Get(ctx context.Context, id int) *domain.Employee {
	args := r.Called(id)
	return args.Get(0).(*domain.Employee)
}

func (r *Repository) Exists(ctx context.Context, CardNumberID string) bool {
	args := r.Called(CardNumberID)
	return args.Get(0).(bool)
}

func (r *Repository) Save(ctx context.Context, employee domain.Employee) int {
	args := r.Called(employee)
	return args.Get(0).(int)
}

func (r *Repository) Update(ctx context.Context, employee domain.Employee) {
}

func (r *Repository) Delete(ctx context.Context, id int) {
}