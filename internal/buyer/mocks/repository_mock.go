package mocks

import (
	"context"

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) GetAll(ctx context.Context) []domain.Buyer {
	args := r.Called()
	return args.Get(0).([]domain.Buyer)
}

func (r *Repository) Get(ctx context.Context, id int) *domain.Buyer {
	args := r.Called(id)
	return args.Get(0).(*domain.Buyer)
}

func (r *Repository) Exists(ctx context.Context, cardNumberID string) bool {
	args := r.Called(cardNumberID)
	return args.Get(0).(bool)
}

func (r *Repository) Save(ctx context.Context, buyer domain.Buyer) int {
	args := r.Called(buyer)
	return args.Get(0).(int)
}

func (r *Repository) Update(ctx context.Context, buyer domain.Buyer) {
	r.Called(buyer).Get(0)

}

func (r *Repository) Delete(ctx context.Context, id int) {
	r.Called(id)
}
