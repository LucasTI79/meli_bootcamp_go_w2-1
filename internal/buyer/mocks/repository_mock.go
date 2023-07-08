package mocks

import (

	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) GetAll() []domain.Buyer {
	args := r.Called()
	return args.Get(0).([]domain.Buyer)
}

func (r *Repository) Get(id int) *domain.Buyer {
	args := r.Called(id)
	return args.Get(0).(*domain.Buyer)
}

func (r *Repository) Exists(cardNumberID string) bool {
	args := r.Called(cardNumberID)
	return args.Get(0).(bool)
}

func (r *Repository) Save(buyer domain.Buyer) int {
	args := r.Called(buyer)
	return args.Get(0).(int)
}

func (r *Repository) Update(buyer domain.Buyer) {
	r.Called(buyer).Get(0)

}

func (r *Repository) Delete(id int) {
	r.Called(id)
}

func (r *Repository) CountPuchasesbyAllBuyers() []domain.PuchasesByBuyerReport {
	args := r.Called()
	return args.Get(0).([]domain.PuchasesByBuyerReport)
}

func (r *Repository) CountPuchasesbyBuyer(id int) *domain.PuchasesByBuyerReport {
	args := r.Called()
	return args.Get(0).(*domain.PuchasesByBuyerReport)
}