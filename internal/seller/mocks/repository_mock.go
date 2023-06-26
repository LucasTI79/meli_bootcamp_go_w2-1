package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) GetAll(ctx context.Context) []domain.Seller {
	args := r.Called()
	return args.Get(0).([]domain.Seller)
}

func (r *Repository) Get(ctx context.Context, id int) *domain.Seller {
	args := r.Called(id)
	return args.Get(0).(*domain.Seller)
}

func (r *Repository) Exists(ctx context.Context, cid int) bool {
	args := r.Called(cid)
	return args.Get(0).(bool)
}

func (r *Repository) Save(ctx context.Context, seller domain.Seller) int {
	args := r.Called(seller)
	return args.Get(0).(int)
}

func (r *Repository) Update(ctx context.Context, seller domain.Seller) {
	r.Called(seller)
}

func (r *Repository) Delete(ctx context.Context, id int) {
	r.Called(id)
}
