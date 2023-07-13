package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) Get(id int) *domain.InboundOrder {
	args := r.Called(id)
	return args.Get(0).(*domain.InboundOrder)
}

func (r *Repository) Save(i domain.InboundOrder) int {
	args := r.Called(i)
	return args.Get(0).(int)
}

func (r *Repository) Exists(orderNumber string) bool {
	args := r.Called(orderNumber)
	return args.Get(0).(bool)
}