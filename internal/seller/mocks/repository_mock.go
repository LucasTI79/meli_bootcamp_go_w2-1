package mocks

import (
	"github.com/extmatperez/meli_bootcamp_go_w2-1/internal/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) GetAll() []domain.Seller {
	args := r.Called()
	return args.Get(0).([]domain.Seller)
}

func (r *Repository) Get(id int) *domain.Seller {
	args := r.Called(id)
	return args.Get(0).(*domain.Seller)
}

func (r *Repository) Exists(cid int) bool {
	args := r.Called(cid)
	return args.Get(0).(bool)
}

func (r *Repository) Save(seller domain.Seller) int {
	args := r.Called(seller)
	return args.Get(0).(int)
}

func (r *Repository) Update(seller domain.Seller) {
	r.Called(seller)
}

func (r *Repository) Delete(id int) {
	r.Called(id)
}
