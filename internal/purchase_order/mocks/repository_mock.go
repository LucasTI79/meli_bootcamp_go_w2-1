package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) Get(id int) *domain.PurchaseOrder {
	args := r.Called(id)
	return args.Get(0).(*domain.PurchaseOrder)
}
func (r *Repository) Exists(orderNumber string) bool {
	args := r.Called(orderNumber)
	return args.Get(0).(bool)
}
func (r *Repository) Save(purchaseOrder domain.PurchaseOrder) int {
	args := r.Called(purchaseOrder)
	return args.Get(0).(int)
}
